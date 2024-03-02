package test

import (
	"fmt"
	"testing"
)

func It[T comparable](t *testing.T, name string, testFun func(*testing.T) T, expected T) {
	fmt.Print("It " + name + ": ")
	result := testFun(t)
	if expected != result {
		fmt.Printf("❌(Expected %v, got %v)\n", expected, result)
		t.Fail()
		return
	}
	fmt.Println("✅")
}

func ShouldFail(t *testing.T, name string, testFun func(*testing.T)) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("✅")
		}
	}()
	fmt.Print("It should fail - " + name + ": ")
	testFun(t)
	fmt.Println("❌(Expected to fail)")
	t.Fail()
}

func ShouldNotFail[T comparable](t *testing.T, name string, testFun func(*testing.T) T, expected T) {
	It(t, "should not fail - "+name, testFun, expected)
}

func ItSlice[T comparable](t *testing.T, name string, testFun func(*testing.T) []T, expected []T) {
	fmt.Print("It " + name + ": ")
	result := testFun(t)
	if !areTheSameSlice(result, expected) {
		fmt.Printf("❌(Expected %v, got %v)\n", expected, result)
		t.Fail()
		return
	}
	fmt.Println("✅")
}

func areTheSameSlice[T comparable](result, expected []T) bool {
	if len(result) != len(expected) {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			return false
		}
	}

	return true
}

func ShouldNotFailSlice[T comparable](t *testing.T, name string, testFun func(*testing.T) []T, expected []T) {
	ItSlice(t, "should not fail - "+name, testFun, expected)
}
