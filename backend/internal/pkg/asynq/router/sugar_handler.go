package router

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"reflect"
	"runtime"
	"waterfall-backend/internal/pkg/asynq/message"
)

func sugarHandler(handler interface{}) asynq.HandlerFunc {
	fVal := reflect.ValueOf(handler)
	fType := fVal.Type()
	fName := runtime.FuncForPC(fVal.Pointer()).Name()

	contextType := reflect.TypeOf((*context.Context)(nil)).Elem()

	// 1-2 параметра входящие
	//  1 - ctx - контекст
	//  2 - req - структура запроса (необязательный)

	// 1 параметра результата
	//  1 - error - ошибка

	if fType.NumIn() == 0 || fType.NumIn() > 2 ||
		fType.NumOut() != 1 {
		panic(fmt.Errorf("ошибка в обработчике %s %s: некорректное число параметров", fName, fVal.String()))
	}
	if fType.In(0) != contextType {
		panic(fmt.Errorf("ошибка в обработчике %s %s: первый аргумент должен быть context.Context", fName, fVal.String()))
	}

	hasReq := false
	var reqType reflect.Type

	if fType.NumIn() == 2 {
		hasReq = true
		reqType = fType.In(1)
	}

	var binder func() interface{}
	var isReqPtr bool

	if hasReq {
		t := reqType
		if reqType.Kind() == reflect.Ptr {
			t = reqType.Elem()
			isReqPtr = true
		}

		binder = func() interface{} {
			return reflect.New(t).Interface()
		}
	}

	invokeHandler := func(ctx context.Context, task *asynq.Task) error {
		in := []reflect.Value{reflect.ValueOf(ctx)}
		taskPayload := task.Payload()

		if hasReq {
			req := binder()

			err := message.Unmarshal(taskPayload, req)
			if err != nil {
				return err
			}

			inReq := reflect.ValueOf(req)
			if !isReqPtr {
				inReq = inReq.Elem()
			}

			in = append(in, inReq)
		}

		out := fVal.Call(in)

		err, _ := out[0].Interface().(error)
		return err
	}

	return invokeHandler
}
