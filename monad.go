package monads

type Monad[T any] interface {
	Bind(func(T) Monad[T]) Monad[T]
	Return(T) Monad[T]
}
