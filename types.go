package gonads


type Monad[T any] interface {
	Bind(func(T) Monad[T]) Monad[T]
	Return(T) Monad[T]
}

type Functor[T any] interface {
	Map(func(T) T) Functor[T]
}
