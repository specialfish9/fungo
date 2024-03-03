package stream

import "golang.org/x/exp/maps"

type Collector[T, U any] interface {
	Collect(*Stream[T]) U
}

func Collect[T, U any](s *Stream[T], collector Collector[T, U]) U {
	return collector.Collect(s)
}

type MapCollector[T any, K comparable, V any] struct {
	keyMapper   func(T) K
	valueMapper func(T) V
}

var _ Collector[int, map[int]string] = (*MapCollector[int, int, string])(nil)

func ToMap[T any, K comparable, V any](keyMapper func(T) K, valueMapper func(T) V) *MapCollector[T, K, V] {
	return &MapCollector[T, K, V]{
		keyMapper:   keyMapper,
		valueMapper: valueMapper,
	}
}

func (mc *MapCollector[T, K, V]) Collect(s *Stream[T]) map[K]V {
	if s == nil {
		return nil
	}

	res := map[K]V{}

	res[mc.keyMapper(s.Hd())] = mc.valueMapper(s.Hd())

	maps.Copy(res, mc.Collect(s.Tl()))

	return res

}
