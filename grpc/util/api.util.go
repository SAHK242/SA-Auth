package grpcutil

import (
	"fmt"
	"log"
)

type (
	// Function is a function that takes one argument and returns one value
	Function[T, R any] func(T) R

	// ErrFunction is a function that takes one argument and returns one value and an error
	ErrFunction[T, R any] func(T) (R, error)

	// Supplier is a function that takes no argument and returns one value
	Supplier[T any] func() T

	// ErrSupplier is a function that takes no argument and returns one value and an error
	ErrSupplier[T any] func() (T, error)

	// Consumer is a function that takes one argument and returns nothing
	Consumer[T any] func(T)

	// ErrConsumer is a function that takes one argument and returns nothing and an error
	ErrConsumer[T any] func(T) error

	// Runnable is a function that takes no argument and returns nothing
	Runnable func()

	// ErrRunnable is a function that takes no argument and returns nothing and an error
	ErrRunnable func() error

	// Predicate is a function that takes one argument and returns a boolean
	Predicate[T any] func(T) bool

	// ErrPredicate is a function that takes one argument and returns a boolean and an error
	ErrPredicate[T any] func(T) (bool, error)

	// BiFunction is a function that takes two arguments and returns one value
	BiFunction[T, U, R any] func(T, U) R

	// BiErrFunction is a function that takes two arguments and returns one value and an error
	BiErrFunction[T, U, R any] func(T, U) (R, error)

	// BiSupplier is a function that takes no argument and returns two values
	BiSupplier[T, U any] func() (T, U)

	// BiErrSupplier is a function that takes no argument and returns two values and an error
	BiErrSupplier[T, U any] func() (T, U, error)

	// BiConsumer is a function that takes two arguments and returns nothing
	BiConsumer[T, U any] func(T, U)

	// BiErrConsumer is a function that takes two arguments and returns nothing and an error
	BiErrConsumer[T, U any] func(T, U) error

	// BiPredicate is a function that takes two arguments and returns a boolean
	BiPredicate[T, U any] func(T, U) bool

	// BiErrPredicate is a function that takes two arguments and returns a boolean and an error
	BiErrPredicate[T, U any] func(T, U) (bool, error)

	// TriFunction is a function that takes three arguments and returns one value
	TriFunction[T, U, V, R any] func(T, U, V) R

	// TriErrFunction is a function that takes three arguments and returns one value and an error
	TriErrFunction[T, U, V, R any] func(T, U, V) (R, error)

	// TriSupplier is a function that takes no argument and returns three values
	TriSupplier[T, U, V any] func() (T, U, V)

	// TriErrSupplier is a function that takes no argument and returns three values and an error
	TriErrSupplier[T, U, V any] func() (T, U, V, error)

	// TriConsumer is a function that takes three arguments and returns nothing
	TriConsumer[T, U, V any] func(T, U, V)

	// TriErrConsumer is a function that takes three arguments and returns nothing and an error
	TriErrConsumer[T, U, V any] func(T, U, V) error

	// TriPredicate is a function that takes three arguments and returns a boolean
	TriPredicate[T, U, V any] func(T, U, V) bool

	// TriErrPredicate is a function that takes three arguments and returns a boolean and an error
	TriErrPredicate[T, U, V any] func(T, U, V) (bool, error)

	// DoublyFunction is a function that takes two arguments and returns two values
	DoublyFunction[T1, T2, R1, R2 any] func(T1, T2) (R1, R2)

	// DoublyErrFunction is a function that takes two arguments and returns two values and an error
	DoublyErrFunction[T1, T2, R1, R2 any] func(T1, T2) (R1, R2, error)

	// TriplyFunction is a function that takes three arguments and returns three values
	TriplyFunction[T1, T2, T3, R1, R2, R3 any] func(T1, T2, T3) (R1, R2, R3)

	// TriplyErrFunction is a function that takes three arguments and returns three values and an error
	TriplyErrFunction[T1, T2, T3, R1, R2, R3 any] func(T1, T2, T3) (R1, R2, R3, error)
)

var (
	recoverFunc = func() {
		if r := recover(); r != nil {
			log.Printf("recovered in WithSafe. %v\n", r)
		}
	}

	recoverErrFunc = func(err *error) {
		if r := recover(); r != nil {
			log.Printf("recovered in WithSafe. %v\n", r)

			if e, ok := r.(error); ok {
				*err = e
			} else {
				*err = fmt.Errorf("%v", r)
			}
		}
	}
)

func WithSafeFunction[T, R any](in T, f Function[T, R]) R {
	defer recoverFunc()
	return f(in)
}

func WithSafeErrFunction[T, R any](in T, f ErrFunction[T, R]) (out R, err error) {
	defer recoverErrFunc(&err)
	return f(in)
}

func WithSafeSupplier[T any](f Supplier[T]) T {
	defer recoverFunc()
	return f()
}

func WithSafeErrSupplier[T any](f ErrSupplier[T]) (out T, err error) {
	defer recoverErrFunc(&err)
	return f()
}

func WithSafeConsumer[T any](in T, f Consumer[T]) {
	defer recoverFunc()
	f(in)
}

func WithSafeErrConsumer[T any](in T, f ErrConsumer[T]) (err error) {
	defer recoverErrFunc(&err)
	return f(in)
}

func WithSafeRunnable(f Runnable) {
	defer recoverFunc()
	f()
}

func WithSafeErrRunnable(f ErrRunnable) (err error) {
	defer recoverErrFunc(&err)
	return f()
}

func WithSafePredicate[T any](in T, f Predicate[T]) bool {
	defer recoverFunc()
	return f(in)
}

func WithSafeErrPredicate[T any](in T, f ErrPredicate[T]) (out bool, err error) {
	defer recoverErrFunc(&err)
	return f(in)
}

func WithSafeBiFunction[T, U, R any](in1 T, in2 U, f BiFunction[T, U, R]) R {
	defer recoverFunc()
	return f(in1, in2)
}

func WithSafeErrBiFunction[T, U, R any](in1 T, in2 U, f BiErrFunction[T, U, R]) (out R, err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2)
}

func WithSafeBiSupplier[T, U any](f BiSupplier[T, U]) (out1 T, out2 U) {
	defer recoverFunc()
	return f()
}

func WithSafeErrBiSupplier[T, U any](f BiErrSupplier[T, U]) (out1 T, out2 U, err error) {
	defer recoverErrFunc(&err)
	return f()
}

func WithSafeBiConsumer[T, U any](in1 T, in2 U, f BiConsumer[T, U]) {
	defer recoverFunc()
	f(in1, in2)
}

func WithSafeErrBiConsumer[T, U any](in1 T, in2 U, f BiErrConsumer[T, U]) (err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2)
}

func WithSafeBiPredicate[T, U any](in1 T, in2 U, f BiPredicate[T, U]) bool {
	defer recoverFunc()
	return f(in1, in2)
}

func WithSafeErrBiPredicate[T, U any](in1 T, in2 U, f BiErrPredicate[T, U]) (out bool, err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2)
}

func WithSafeTriFunction[T, U, V, R any](in1 T, in2 U, in3 V, f TriFunction[T, U, V, R]) R {
	defer recoverFunc()
	return f(in1, in2, in3)
}

func WithSafeErrTriFunction[T, U, V, R any](in1 T, in2 U, in3 V, f TriErrFunction[T, U, V, R]) (out R, err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2, in3)
}

func WithSafeTriSupplier[T, U, V any](f TriSupplier[T, U, V]) (out1 T, out2 U, out3 V) {
	defer recoverFunc()
	return f()
}

func WithSafeErrTriSupplier[T, U, V any](f TriErrSupplier[T, U, V]) (out1 T, out2 U, out3 V, err error) {
	defer recoverErrFunc(&err)
	return f()
}

func WithSafeTriConsumer[T, U, V any](in1 T, in2 U, in3 V, f TriConsumer[T, U, V]) {
	defer recoverFunc()
	f(in1, in2, in3)
}

func WithSafeErrTriConsumer[T, U, V any](in1 T, in2 U, in3 V, f TriErrConsumer[T, U, V]) (err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2, in3)
}

func WithSafeTriPredicate[T, U, V any](in1 T, in2 U, in3 V, f TriPredicate[T, U, V]) bool {
	defer recoverFunc()
	return f(in1, in2, in3)
}

func WithSafeErrTriPredicate[T, U, V any](in1 T, in2 U, in3 V, f TriErrPredicate[T, U, V]) (out bool, err error) {
	defer recoverErrFunc(&err)
	return f(in1, in2, in3)
}

func WithSafeDoublyFunction[T1, T2, R1, R2 any](in1 T1, int2 T2, f DoublyFunction[T1, T2, R1, R2]) (out1 R1, out2 R2) {
	defer recoverFunc()
	return f(in1, int2)
}

func WithSafeErrDoublyFunction[T1, T2, R1, R2 any](in1 T1, int2 T2, f DoublyErrFunction[T1, T2, R1, R2]) (out1 R1, out2 R2, err error) {
	defer recoverErrFunc(&err)
	return f(in1, int2)
}

func WithSafeTriplyFunction[T1, T2, T3, R1, R2, R3 any](in1 T1, int2 T2, in3 T3, f TriplyFunction[T1, T2, T3, R1, R2, R3]) (out1 R1, out2 R2, out3 R3) {
	defer recoverFunc()
	return f(in1, int2, in3)
}

func WithSafeErrTriplyFunction[T1, T2, T3, R1, R2, R3 any](in1 T1, int2 T2, in3 T3, f TriplyErrFunction[T1, T2, T3, R1, R2, R3]) (out1 R1, out2 R2, out3 R3, err error) {
	defer recoverErrFunc(&err)
	return f(in1, int2, in3)
}
