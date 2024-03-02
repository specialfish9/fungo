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

func reduce(sli []int, f func(int, int) int, v int) int {
	var z int = v

	for _, x := range sli {
		z = f(z, x)
	}

	return z
}

func reduce2(sli []int, i int, f func(int, int) int, v int) int {
	if sli == nil {
		return v
	}

	return f(sli[i], reduce2(sli[i:], i+1, f, v))
}
