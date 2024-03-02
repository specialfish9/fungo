package stream_test

import (
	"fmt"
	"fungo/stream"
	"fungo/test"
	"testing"
)

// ToSlice

func testToSlice(t *testing.T) []int {
	return stream.Of([]int{1, 2, 3, 4}).ToSlice()
}

func testToSliceEmpty(t *testing.T) []int {
	return stream.Of([]int{}).ToSlice()
}

func testToSliceNil(t *testing.T) []int {
	return stream.Of[int](nil).ToSlice()
}

// Map

func testMap(t *testing.T) []int {
	return stream.Of([]int{1, 2, 3}).Map(func(i int) int { return i * 10 }).ToSlice()
}

func testMapEmpty(t *testing.T) []int {
	return stream.Of([]int{}).Map(func(i int) int { return i * 10 }).ToSlice()
}

func testMapNil(t *testing.T) []int {
	return stream.Of[int](nil).Map(func(i int) int { return i * 10 }).ToSlice()
}

// Filter

func testFilter(t *testing.T) []int {
	return stream.Of([]int{1, 2, 3}).Filter(func(i int) bool { return i%2 == 0 }).ToSlice()
}

func testFilterEmpty(t *testing.T) []int {
	return stream.Of([]int{}).Filter(func(i int) bool { return i%2 == 0 }).ToSlice()
}

func testFilterNil(t *testing.T) []int {
	return stream.Of[int](nil).Filter(func(i int) bool { return i%2 == 0 }).ToSlice()
}

// Take

func testTake(t *testing.T) []int {
	return stream.Of([]int{4, 5, 6}).Take(2)
}

func testTakeEmpty(t *testing.T) []int {
	return stream.Of([]int{}).Take(0)
}

func testTakeTooMany(t *testing.T) {
	stream.Of([]int{1, 2, 3, 4}).Take(10)
}

func testTakeNil(t *testing.T) []int {
	return stream.Of[int](nil).Take(0)
}

// Count

func testCount(t *testing.T) int {
	return stream.Of([]int{1, 2, 3}).Count()
}

func testCountEmpty(t *testing.T) int {
	return stream.Of([]int{}).Count()
}

func testCountNil(t *testing.T) int {
	return stream.Of[int](nil).Count()
}

// FoldL

func testFoldL(t *testing.T) int {
	return stream.Of([]int{4, 2, 2}).FoldL(func(acc, curr int) int { return acc / curr }, 400)
}

// FoldR

func testFoldR(t *testing.T) int {
	return stream.Of([]int{2, 2, 4}).FoldR(func(curr, acc int) int { return acc / curr }, 400)
}

func TestStream(t *testing.T) {
	fmt.Println("===STREAM TESTS===")

	test.ShouldNotFailSlice(t, "ToSlice", testToSlice, []int{1, 2, 3, 4})
	test.ShouldNotFailSlice(t, "ToSlice empty", testToSliceEmpty, []int{})

	test.ShouldNotFailSlice(t, "Map", testMap, []int{10, 20, 30})
	test.ShouldNotFailSlice(t, "Map empty", testMapEmpty, []int{})
	test.ShouldNotFailSlice(t, "Map nil", testMapNil, nil)

	test.ShouldNotFailSlice(t, "Filter", testFilter, []int{2})
	test.ShouldNotFailSlice(t, "Filter empty", testFilterEmpty, []int{})
	test.ShouldNotFailSlice(t, "Filter nil", testFilterNil, nil)

	test.ShouldNotFailSlice(t, "Take", testTake, []int{4, 5})
	test.ShouldNotFailSlice(t, "Take empty", testTakeEmpty, []int{})
	test.ShouldNotFailSlice(t, "Take nil", testTakeNil, nil)
	test.ShouldFail(t, "Take too many", testTakeTooMany)

	test.ShouldNotFail(t, "Count", testCount, 3)
	test.ShouldNotFail(t, "Count empty", testCountEmpty, 0)
	test.ShouldNotFail(t, "Count nil", testCountNil, 0)
	test.ShouldNotFail(t, "FoldL", testFoldL, 25)
	test.ShouldNotFail(t, "FoldR", testFoldR, 25)
}
