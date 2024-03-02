package main

import "fmt"

func reduce2L(sli []int, i int, f func(int, int) int, v int) int {
	if i >= len(sli) {
		return v
	}

	var z = f(v, sli[i])

	return reduce2L(sli, i+1, f, z)
}

func main() {
	s := []int{1, 2, 3}

	fmt.Printf("%d\n", reduce2L(s, 0, func(i1, i2 int) int { fmt.Printf("sum %d + %d\n", i1, i2); return i1 + i2 }, 0))

}
