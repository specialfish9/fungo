package stream_test

import (
	"fmt"
	"fungo/stream"
	"fungo/test"
	"strconv"
	"testing"
)

// Map Func
func testMapFunc(t *testing.T) []string {
	s := stream.Of([]int{1, 2, 3})
	return stream.Map(s, strconv.Itoa).ToSlice()
}

func testMapFuncEmpty(t *testing.T) []string {
	s := stream.Of([]int{})
	return stream.Map(s, strconv.Itoa).ToSlice()
}

func testMapFuncNil(t *testing.T) []string {
	s := stream.Of[int](nil)
	return stream.Map(s, strconv.Itoa).ToSlice()
}

// MapI Func
func testMapIFunc(t *testing.T) []string {
	s := stream.Of([]int{4, 5, 6})
	return stream.MapI(s, func(i int, _ int) string { return strconv.Itoa(i) }).ToSlice()
}

func testMapIFuncEmpty(t *testing.T) []string {
	s := stream.Of([]int{})
	return stream.MapI(s, func(i int, _ int) string { return strconv.Itoa(i) }).ToSlice()
}

func testMapIFuncNil(t *testing.T) []string {
	s := stream.Of[int](nil)
	return stream.MapI(s, func(i int, _ int) string { return strconv.Itoa(i) }).ToSlice()
}

func TestFunctions(t *testing.T) {
	fmt.Println("===STREAM FUNCTIONS TESTS===")

	test.ShouldNotFailSlice(t, "Map function", testMapFunc, []string{"1", "2", "3"})
	test.ShouldNotFailSlice(t, "Map function empty", testMapFuncEmpty, []string{})
	test.ShouldNotFailSlice(t, "Map function nil", testMapFuncNil, nil)

	test.ShouldNotFailSlice(t, "MapI function", testMapIFunc, []string{"0", "1", "2"})
	test.ShouldNotFailSlice(t, "MapI function empty", testMapIFuncEmpty, []string{})
	test.ShouldNotFailSlice(t, "MapI function nil", testMapIFuncNil, nil)
}
