package monads

import (

	"fmt"
	"reflect"

	"strings"
)

type Maybe[T any] interface {
    Monad[T]
    Unwrap() T
}

type SomeType[T any] struct {
	value T
}

type NoneType[T any] struct {}

func Some[T any](value T) Maybe[T] {
	return &SomeType[T]{value: value}
}

func None[T any]() Maybe[T] {
	return &NoneType[T]{}
}

func (m *SomeType[T]) Bind(f func(T) Monad[T]) Monad[T] {
	return f(m.value)
}

func (m *SomeType[T]) Unwrap() T {
	return m.value
}
func (m *SomeType[T]) String() string {
	typeName := reflect.TypeOf(m).Elem().Name()
	return fmt.Sprintf("Some[%s{%v}", strings.SplitAfter(typeName, "[")[1], m.value)
}

func (m *SomeType[T]) Return(value T) Monad[T] {
	return Some(value)
}

func (m *NoneType[T]) Bind(f func(T) Monad[T]) Monad[T] {
	return None[T]()
}


func (m *NoneType[T]) Unwrap() T {
	panic("Unwrap called on None")
}

func (m *NoneType[T]) String() string {
	typeName := reflect.TypeOf(m).Elem().Name()
	return fmt.Sprintf("None[%s", strings.SplitAfter(typeName, "[")[1])
}

func (m *NoneType[T]) Return(value T) Monad[T] {
	return None[T]()
}
