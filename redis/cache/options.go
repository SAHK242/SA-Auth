package cache

import (
	"context"
	"log"
	"time"
)

var (
	defaultTTL = 10 * time.Minute

	logger = log.Default()
)

type (
	options[T any] struct {
		loadFunc      func(ctx context.Context, key string) (T, error)
		keyEncodeFunc func(key string) string
		keyDecodeFunc func(key string) string
		errFunc       func(err error)
		ttl           time.Duration
		statsEnabled  bool
	}

	Option[T any] func(*options[T])
)

func WithLoader[T any](f func(ctx context.Context, key string) (T, error)) Option[T] {
	return func(o *options[T]) {
		if f == nil {
			panic("loader function cannot be nil")
		} else {
			o.loadFunc = f
		}
	}
}

func WithKeyGenerator[T any](f func(key string) string) Option[T] {
	return func(o *options[T]) {
		if f == nil {
			o.keyEncodeFunc = defaultKeyGenerator
		} else {
			o.keyEncodeFunc = f
		}
	}
}

func WithKeyDecoder[T any](f func(key string) string) Option[T] {
	return func(o *options[T]) {
		if f == nil {
			o.keyDecodeFunc = defaultKeyGenerator
		} else {
			o.keyDecodeFunc = f
		}
	}
}

func WithErrorHandler[T any](f func(err error)) Option[T] {
	return func(o *options[T]) {
		if f == nil {
			o.errFunc = defaultErrorHandler
		} else {
			o.errFunc = f
		}
	}
}

func WithTTL[T any](ttl time.Duration) Option[T] {
	return func(o *options[T]) {
		if ttl == 0 {
			o.ttl = defaultTTL
		} else {
			o.ttl = ttl
		}
	}
}

func WithStats[T any]() Option[T] {
	return func(o *options[T]) {
		o.statsEnabled = true
	}
}

func eval[T any](opts []Option[T]) *options[T] {
	o := &options[T]{
		errFunc:       defaultErrorHandler,
		keyEncodeFunc: defaultKeyGenerator,
		keyDecodeFunc: defaultKeyGenerator,
		ttl:           defaultTTL,
		statsEnabled:  false,
	}

	for _, opt := range opts {
		opt(o)
	}

	if o.keyEncodeFunc == nil {
		panic("key generator function cannot be nil")
	}

	if o.loadFunc == nil {
		logger.Println("loader function is not provided")
	}

	return o
}

func defaultErrorHandler(err error) {
	logger.Println(err)
}

func defaultKeyGenerator(key string) string {
	return key
}
