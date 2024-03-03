package stream

func Of[T any](slice []T) *Stream[T] {
	if len(slice) == 0 {
		return nil
	}
	return new(slice[0], func() *Stream[T] { return Of(slice[1:]) })
}

func CountFrom(n int) *Stream[int] {
	return new(n, func() *Stream[int] { return CountFrom(n + 1) })
}
