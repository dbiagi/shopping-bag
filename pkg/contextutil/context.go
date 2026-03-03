package contextutil

import (
	"context"
)

type ContextKey string

const (
	ServiceConfigKey ContextKey = "_serviceConfig"
)

func WithServiceConfig[T any](ctx context.Context, config T) context.Context {
	return context.WithValue(ctx, ServiceConfigKey, config)
}

func GetServiceConfig[T any](ctx context.Context) T {
	if config, ok := ctx.Value(ServiceConfigKey).(T); ok {
		return config
	}

	var zeroValue T
	return zeroValue
}
