package lazy

import (
	"sync"
)

func Cond[T any](cond bool, v1 *Lazy[T], v2 *Lazy[T]) (T, error) {
	if cond {
		return v1.Eval()
	}
	return v2.Eval()
}

func Of[T any](f func() T) *Lazy[T] {
	return &Lazy[T]{
		New: func() (T, error) {
			return f(), nil
		},
	}
}

func OfE[T any](f func() (T, error)) *Lazy[T] {
	return &Lazy[T]{New: f}
}

func Noop[T any]() *Lazy[T] {
	return &Lazy[T]{
		New: func() (T, error) {
			var noop T
			return noop, nil
		},
	}
}

type Lazy[T any] struct {
	New   func() (T, error)
	once  sync.Once
	value T
	err   error
}

func (this *Lazy[T]) Eval() (T, error) {
	this.once.Do(func() {
		if this.New != nil {
			v, err := this.New()
			this.value = v
			this.err = err
			this.New = nil // so that f can now be GC'ed
		}
	})

	return this.value, this.err
}
