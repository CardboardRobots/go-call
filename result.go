package call

type Result[T any, E any] struct {
	Value T
	Error E
}

func Ok[T any](value T) Result[T, any] {
	return Result[T, any]{
		Value: value,
	}
}

func Err[E any](error E) Result[any, E] {
	return Result[any, E]{
		Error: error,
	}
}
