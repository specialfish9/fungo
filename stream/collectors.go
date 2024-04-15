package stream

import "golang.org/x/exp/maps"

// A Collector takes consumes all the element of a Stream in a collection of type
// U.
type Collector[T, U any] interface {
	Collect(*Stream[T]) U
}

func Collect[T, U any](s *Stream[T], collector Collector[T, U]) U {
	return collector.Collect(s)
}

// MapCollector is a collector that populate a map[K,V] with the elements of
// the Stream. To do so, it applies to each element of the Stream both its
// keyMapper and valueMapper.
type MapCollector[T any, K comparable, V any] struct {
	// keyMapper map the values of the Stream to the keys of the resulting map.
	keyMapper func(T) K
	// valueMapper map the values of the Stream to the values of the resulting map.
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
