package stream_test

import (
	"fmt"
	"fungo/stream"
	"fungo/test"
	"strconv"
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

// MapI
func testMapI(t *testing.T) []int {
	return stream.Of([]int{4, 5, 6}).MapI(func(i int, _ int) int { return i }).ToSlice()
}

func testMapIEmpty(t *testing.T) []int {
	return stream.Of([]int{}).MapI(func(i int, _ int) int { return i }).ToSlice()
}

func testMapINil(t *testing.T) []int {
	return stream.Of[int](nil).MapI(func(i int, _ int) int { return i }).ToSlice()
}

// Map2
func testMap2(t *testing.T) []string {
	values := stream.Of([]int{1, 2, 3}).Map2(func(i int) any { return strconv.Itoa(i) }).ToSlice()
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

func testMap2Empty(t *testing.T) []string {
	values := stream.Of([]int{}).Map2(func(i int) any { return strconv.Itoa(i) }).ToSlice()
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

func testMap2Nil(t *testing.T) []string {
	values := stream.Of[int](nil).Map2(func(i int) any { return strconv.Itoa(i) }).ToSlice()
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
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

// ForEach
func testForEach(t *testing.T) []bool {
	n := 3
	flags := make([]bool, n)
	oneToTen := stream.CountFrom(0).Take(n)

	stream.Of(oneToTen).ForEach(func(i int) {
		flags[i] = true
	})

	return flags
}

func testForEachEmpty(t *testing.T) bool {
	// test that it does not throw errors
	stream.Of([]int{}).ForEach(func(i int) {})
	return true
}

func testForEachNil(t *testing.T) bool {
	// test that it does not throw errors
	stream.Of[int](nil).ForEach(func(i int) {})
	return true
}

// AllMatch
func testAllMatch(t *testing.T) bool {
	return stream.Of([]int{2, 4, 6, 8}).AllMatch(func(i int) bool { return i%2 == 0 })
}

func testAllMatchNoMatch(t *testing.T) bool {
	return stream.Of([]int{1, 2, 3, 4, 5}).AllMatch(func(i int) bool { return i%2 == 0 })
}

func testAllMatchEmpty(t *testing.T) bool {
	return stream.Of([]int{}).AllMatch(func(i int) bool { return i == 3 })
}

func testAllMatchNil(t *testing.T) bool {
	return stream.Of[int](nil).AllMatch(func(i int) bool { return i == 3 })
}

// AnyMatch
func testAnyMatch(t *testing.T) bool {
	return stream.Of([]int{1, 2, 3, 4, 5}).AnyMatch(func(i int) bool { return i == 3 })
}

func testAnyMatchNoMatch(t *testing.T) bool {
	return stream.Of([]int{1, 2, 3, 4, 5}).AnyMatch(func(i int) bool { return i == 0 })
}

func testAnyMatchEmpty(t *testing.T) bool {
	return stream.Of([]int{}).AnyMatch(func(i int) bool { return i == 3 })
}

func testAnyMatchNil(t *testing.T) bool {
	return stream.Of[int](nil).AnyMatch(func(i int) bool { return i == 3 })
}

func TestStream(t *testing.T) {
	fmt.Println("===STREAM TESTS===")

	test.ShouldNotFailSlice(t, "ToSlice", testToSlice, []int{1, 2, 3, 4})
	test.ShouldNotFailSlice(t, "ToSlice empty", testToSliceEmpty, []int{})

	test.ShouldNotFailSlice(t, "Map", testMap, []int{10, 20, 30})
	test.ShouldNotFailSlice(t, "Map empty", testMapEmpty, []int{})
	test.ShouldNotFailSlice(t, "Map nil", testMapNil, nil)

	test.ShouldNotFailSlice(t, "MapI", testMapI, []int{0, 1, 2})
	test.ShouldNotFailSlice(t, "MapI empty", testMapIEmpty, []int{})
	test.ShouldNotFailSlice(t, "MapI nil", testMapINil, nil)

	test.ShouldNotFailSlice(t, "Map2", testMap2, []string{"1", "2", "3"})
	test.ShouldNotFailSlice(t, "Map2 empty", testMap2Empty, []string{})
	test.ShouldNotFailSlice(t, "Map2 nil", testMap2Nil, nil)

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

	test.ShouldNotFailSlice(t, "ForEach", testForEach, []bool{true, true, true})
	test.ShouldNotFail(t, "ForEach empty", testForEachEmpty, true)
	test.ShouldNotFail(t, "ForEach nil", testForEachNil, true)

	test.ShouldNotFail(t, "AnyMatch match", testAnyMatch, true)
	test.ShouldNotFail(t, "AnyMatch no match", testAnyMatchNoMatch, false)
	test.ShouldNotFail(t, "AnyMatch empty", testAnyMatchNil, false)
	test.ShouldNotFail(t, "AnyMatch nil", testAnyMatchNil, false)

	test.ShouldNotFail(t, "AllMatch match", testAllMatch, true)
	test.ShouldNotFail(t, "AllMatch no match", testAllMatchNoMatch, false)
	test.ShouldNotFail(t, "AllMatch empty", testAllMatchNil, true)
	test.ShouldNotFail(t, "AllMatch nil", testAllMatchNil, true)
}
