package stream

// Map applies the given transformation to the elements of the given stream,
// returning a new Stream with the transformed elements.
func Map[T, U any](s *Stream[T], f Transformation[T, U]) *Stream[U] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[U] { return Map(s.Tl(), f) })
}

// MapI is like Map but with a transformation with index.
func MapI[T, U any](s *Stream[T], f TransformationI[T, U]) *Stream[U] {
	return mapI(s, f, 0)
}

// mapI is the internal implementation of MapI.
func mapI[T, U any](s *Stream[T], f TransformationI[T, U], index int) *Stream[U] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[U] { return mapI(s.Tl(), f, index+1) })
}

// Id is the identity function. It simply returns what it takes as input.
func Id[T any](t T) T { return t }
