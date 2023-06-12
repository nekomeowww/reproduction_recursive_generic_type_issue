package with_generics

// It doesn't matter if the innerT1 struct unexported or exported, the result is the same.
// It also doesn't matter if the R type parameter is infer to a pointer or not, the result is the same.
type innerT1[T any, R *T1[T]] struct {
	reference *R
}

type T1[T any] struct {
	e *innerT1[T, *T1[T]]
}
