package original_scenario

import (
	"fmt"
)

// A struct that contains a generic type and a pointer to the container type itself
// to help to return the original container type after calling a method in order to
// achieve chained method calls just like WithXXX(...).WithXXXB(...), etc.
//
// NOTICE: TypeB[T] isn't necessary to reproduce such behavior, but it's added here
// to give a overview of the whole picture of what I initially wanted to achieve.
type CommonOption[T any, C TypeA[T] | TypeB[T]] struct {
	value     T
	container *C
}

// Let's say this method assigns a value to the field `value` and returns the
// original container type.
func (o *CommonOption[T, C]) WithValue(v T) *C {
	o.value = v

	return o.container
}

// A general generic struct that embedded the CommonOption[T, TypeA[T]] struct
// I described above. And the second of the type parameter to CommonOption[T, TypeA[T]]
// is TypeA[T] itself.
//
// In this scenario, TypeA is not behaved the same way as TypeB.
type TypeA[T any] struct {
	*CommonOption[T, TypeA[T]]

	subFieldA string
}

func NewTypeA[T any]() *TypeA[T] {
	t := &TypeA[T]{
		subFieldA: "random name",
	}
	t.CommonOption = &CommonOption[T, TypeA[T]]{container: t}

	return t
}

// Let's say this method reads the subFieldA and returned a formatted string
// that contains the value of the field `value` in the CommonOption[T, TypeA[T]].
func (t TypeA[T]) GetValue() string {
	return fmt.Sprintf("%s: %v", t.subFieldA, t.value)
}

// ----------------------------------

// A general generic struct that embedded the CommonOption[T, TypeB[T]] struct
// I described above. And the second of the type parameter to CommonOption[T, TypeB[T]]
// is TypeB[T] itself.
//
// In this scenario, TypeB is not behaved the same way as TypeA.
type TypeB[T any] struct {
	*CommonOption[T, TypeB[T]]
}

func NewTypeB[T any]() *TypeB[T] {
	t := &TypeB[T]{}
	t.CommonOption = &CommonOption[T, TypeB[T]]{container: t}

	return t
}

// Let's say this method appended a "TypeB: " prefix to the value of the field `value`
// in the CommonOption[T, TypeB[T]] and returned the formatted string.
func (t TypeB[T]) GetValue() string {
	return fmt.Sprintf("TypeB: %v", t.value)
}
