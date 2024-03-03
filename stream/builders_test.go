package stream_test

import (
	"fmt"
	"fungo/stream"
	"fungo/test"
	"testing"
)

// Of
func testOf(t *testing.T) bool {
	s := stream.Of([]int{1, 2, 3})

	return s.Hd() == 1 && s.Tl().Hd() == 2 && s.Tl().Tl().Hd() == 3
}

func testOfEmpty(t *testing.T) bool {
	s := stream.Of([]int{})

	return s == nil
}

func testOfNil(t *testing.T) bool {
	s := stream.Of[int](nil)

	return s == nil
}

func TestBuilder(t *testing.T) {
	fmt.Println("===STREAM BUILDER TESTS===")

	test.ShouldNotFail(t, "Of", testOf, true)
	test.ShouldNotFail(t, "Of empty", testOfEmpty, true)
	test.ShouldNotFail(t, "Of nil", testOfNil, true)
}
