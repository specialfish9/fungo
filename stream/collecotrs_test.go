package stream_test

import (
	"fmt"
	"fungo/stream"
	"fungo/test"
	"strconv"
	"testing"
)

func testMapCollector(t *testing.T) bool {
	s := stream.Of([]int{1, 2, 3})

	m := stream.Collect[int, map[int]string](s, stream.ToMap(stream.Id[int], strconv.Itoa))

	expected := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	for k1, v1 := range expected {
		v2, ok := m[k1]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}

func TestCollector(t *testing.T) {
	fmt.Println("===STREAM COLLECTOR TEST===")
	test.ShouldNotFail(t, "MapCollector", testMapCollector, true)
}
