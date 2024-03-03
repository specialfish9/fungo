package stream

func Map[T, U any](s *Stream[T], f Transformation[T, U]) *Stream[U] {
	if s == nil {
		return nil
	}
	return new(f(s.Hd()), func() *Stream[U] { return Map(s.Tl(), f) })
}

func MapI[T, U any](s *Stream[T], f TransformationI[T, U]) *Stream[U] {
	return mapI(s, f, 0)
}

func mapI[T, U any](s *Stream[T], f TransformationI[T, U], index int) *Stream[U] {
	if s == nil {
		return nil
	}
	return new(f(index, s.Hd()), func() *Stream[U] { return mapI(s.Tl(), f, index+1) })
}

func Id[T any](t T) T { return t }
