package stream

type thunk[T any] func() *Stream[T]

type Predicate[T any] func(T) bool
type Transformation[T, U any] func(T) U
type TransformationI[T, U any] func(int, T) U

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

func (s *Stream[T]) Map(f Transformation[T, T]) *Stream[T] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[T] { return s.Tl().Map(f) })
}

func (s *Stream[T]) MapI(f TransformationI[T, T]) *Stream[T] {
	return s.mapI(f, 0)
}

func (s *Stream[T]) mapI(f TransformationI[T, T], index int) *Stream[T] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[T] { return s.Tl().mapI(f, index+1) })
}

func (s *Stream[T]) Map2(f Transformation[T, any]) *Stream[any] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[any] { return s.Tl().Map2(f) })
}

func (s *Stream[T]) Map2I(f TransformationI[T, any]) *Stream[any] {
	return s.map2I(f, 0)
}

func (s *Stream[T]) map2I(f TransformationI[T, any], index int) *Stream[any] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[any] { return s.Tl().map2I(f, index+1) })
}

func (s *Stream[T]) Filter(f Predicate[T]) *Stream[T] {
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

func (s *Stream[T]) ForEach(f func(T)) {
	if s == nil {
		return
	}
	f(s.Hd())
	s.Tl().ForEach(f)
}

func (s *Stream[T]) AnyMatch(f Predicate[T]) bool {
	if s == nil {
		return false
	}

	if f(s.Hd()) {
		return true
	}

	return s.Tl().AnyMatch(f)
}

func (s *Stream[T]) AllMatch(f Predicate[T]) bool {
	if s == nil {
		return true
	}

	if !f(s.Hd()) {
		return false
	}

	return s.Tl().AllMatch(f)
}
