package message

import "context"

type outgoingHeadersKey struct{}

func OutgoingHeadersToContext(ctx context.Context, headers map[string]interface{}) context.Context {
	return context.WithValue(ctx, outgoingHeadersKey{}, headers)
}

func OutgoingHeadersFromContext(ctx context.Context) (map[string]interface{}, bool) {
	headers, ok := ctx.Value(outgoingHeadersKey{}).(map[string]interface{})
	if !ok {
		return nil, false
	}
	return headers, true
}

type incomingHeadersKey struct{}

func IncomingHeadersToContext(ctx context.Context, headers map[string]interface{}) context.Context {
	return context.WithValue(ctx, incomingHeadersKey{}, headers)
}

func IncomingHeadersFromContext(ctx context.Context) (map[string]interface{}, bool) {
	headers, ok := ctx.Value(incomingHeadersKey{}).(map[string]interface{})
	if !ok {
		return nil, false
	}
	return headers, true
}
