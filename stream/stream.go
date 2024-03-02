package stream

type thunk[T any] func() *Stream[T]

type Stream[T any] struct {
	value T
	next  thunk[T]
}

func new[T any](value T, next thunk[T]) *Stream[T] {
	return &Stream[T]{
		value: value,
		next:  next,
	}
}

func (s *Stream[T]) Hd() T {
	return s.value
}

func (s *Stream[T]) Tl() *Stream[T] {
	return s.next()
}

func (s *Stream[T]) ToSlice() []T {
	if s == nil {
		return []T{}
	}

	return append([]T{s.value}, s.Tl().ToSlice()...)
}

func (s *Stream[T]) Map(f func(T) T) *Stream[T] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[T] { return s.Tl().Map(f) })
}

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	if s == nil {
		return nil
	}

	if f(s.Hd()) {
		return new(s.Hd(), func() *Stream[T] { return s.Tl().Filter(f) })
	}

	return s.Tl().Filter(f)
}

func (s *Stream[T]) FoldL(f func(acc T, curr T) T, z T) T {
	if s == nil {
		return z
	}

	v := f(z, s.Hd())

	return s.next().FoldL(f, v)
}

func (s *Stream[T]) FoldR(f func(acc T, curr T) T, z T) T {
	if s == nil {
		return z
	}

	return f(s.Hd(), s.next().FoldR(f, z))
}

func (s *Stream[T]) Take(n int) []T {
	if n == 0 {
		return []T{}
	}

	return append([]T{s.Hd()}, s.Tl().Take(n-1)...)
}

func (s *Stream[T]) Count() int {
	if s == nil {
		return 0
	}
	return 1 + s.Tl().Count()
}
