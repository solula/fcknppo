package enqueuer

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"sync"
	"waterfall-backend/internal/pkg/asynq/message"
)

type Priority string

const (
	PriorityCritical Priority = "critical"
	PriorityDefault  Priority = "default"
	PriorityLow      Priority = "low"
)

type (
	HandlerFunc    func(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error)
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

type Enqueuer interface {
	Enqueue(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error)
	Use(mws ...MiddlewareFunc)
}

type enqueuer struct {
	client *asynq.Client

	mu  sync.RWMutex
	mws []MiddlewareFunc
}

func NewEnqueuer(client *asynq.Client) Enqueuer {
	return &enqueuer{
		client: client,
		mws:    make([]MiddlewareFunc, 0),
	}
}

func (r *enqueuer) Enqueue(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// r.enqueue - конечный обработчик, ставящий задачу в очередь
	h := r.enqueue

	// Вызываем миддлвари в обратном порядке
	for i := len(r.mws) - 1; i >= 0; i-- {
		h = r.mws[i](h)
	}

	return h(ctx, typename, req, opts...)
}

func (r *enqueuer) Use(mws ...MiddlewareFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, fn := range mws {
		r.mws = append(r.mws, fn)
	}
}

func (r *enqueuer) enqueue(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error) {
	headers, ok := message.OutgoingHeadersFromContext(ctx)
	if !ok {
		return "", fmt.Errorf("хэдеры не найдены")
	}
	payload, err := message.Marshal(req, headers)
	if err != nil {
		return "", err
	}

	task := asynq.NewTask(typename, payload)
	taskInfo, err := r.client.Enqueue(task, opts...)
	if err != nil {
		return "", err
	}

	return taskInfo.ID, nil
}
