package stream

type thunk[T any] func() *Stream[T]

// Predicate is a type of function which represents a generic statement of truth.
// It takes one input and performs a boolean evaluation.
type Predicate[T any] func(T) bool

// Transformation is a type of function that take as input one element of a type T,
// applies transoformations into an element of type U and then returns it.
type Transformation[T, U any] func(T) U

// TransformationI is like Transformation but the transofrmation takes also
// an index as parameter.
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

// Hd returns the head of the Stream. The head of a Stream is defined as the
// its first element.
func (s *Stream[T]) Hd() T {
	return s.value
}

// Tl returns the tail of the Stream. The tail of a Stream is defined as a pointer
// to the next element of the Stream.
func (s *Stream[T]) Tl() *Stream[T] {
	return s.next()
}

// ToSlice consumes the Stream saving values into a slice. Returns an empty
// slice if the Stream is nil.
func (s *Stream[T]) ToSlice() []T {
	if s == nil {
		return []T{}
	}

	return append([]T{s.value}, s.Tl().ToSlice()...)
}

// Map applies a transformation to each element of a stream, returning a new
// Stream with the transformed elements of the same type.
func (s *Stream[T]) Map(f Transformation[T, T]) *Stream[T] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[T] { return s.Tl().Map(f) })
}

// MapI is like Map but applies a transformation with index.
func (s *Stream[T]) MapI(f TransformationI[T, T]) *Stream[T] {
	return s.mapI(f, 0)
}

// mapI is the internal implementation of MapI.
func (s *Stream[T]) mapI(f TransformationI[T, T], index int) *Stream[T] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[T] { return s.Tl().mapI(f, index+1) })
}

// Map2 applies a transformation from a type to another to each element of the
// Stream. It returns a new Stream with containing the transformed elements.
func (s *Stream[T]) Map2(f Transformation[T, any]) *Stream[any] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[any] { return s.Tl().Map2(f) })
}

// Map2I is like Map2 but applies a transformation with index.
func (s *Stream[T]) Map2I(f TransformationI[T, any]) *Stream[any] {
	return s.map2I(f, 0)
}

// map2I is the internal implementation of Map2I.
func (s *Stream[T]) map2I(f TransformationI[T, any], index int) *Stream[any] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[any] { return s.Tl().map2I(f, index+1) })
}

// Filter returns a new Stream with all the elements that fullfill the given
// predicate.
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

// Take consumes the first n element of the Stream, saving them into a slice.
func (s *Stream[T]) Take(n int) []T {
	if n == 0 {
		return []T{}
	}

	return append([]T{s.Hd()}, s.Tl().Take(n-1)...)
}

// Count returns the number of element componing the Stream. In case of a Stream
// of unlimited size, it won't terminate.
func (s *Stream[T]) Count() int {
	if s == nil {
		return 0
	}
	return 1 + s.Tl().Count()
}

// ForEach applies the given function f to each element of the Stream.
func (s *Stream[T]) ForEach(f func(T)) {
	if s == nil {
		return
	}
	f(s.Hd())
	s.Tl().ForEach(f)
}

// AnyMatch returns true iff there is at least one element of the stream that
// fullfill the given predicate.
func (s *Stream[T]) AnyMatch(f Predicate[T]) bool {
	if s == nil {
		return false
	}

	if f(s.Hd()) {
		return true
	}

	return s.Tl().AnyMatch(f)
}

// AllMatch returns true iff all elements of the Stream fullfill the given predicate.
func (s *Stream[T]) AllMatch(f Predicate[T]) bool {
	if s == nil {
		return true
	}

	if !f(s.Hd()) {
		return false
	}

	return s.Tl().AllMatch(f)
}
