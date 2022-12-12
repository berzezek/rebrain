package main

import "fmt"

func contains(a []string, x string) bool {
	for _, i := range a {
		if i == x {
			return true
		}
	}
	return false
}

func getMax(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	fmt.Println(contains([]string{"a", "b", "z"}, "z"))
	fmt.Println(getMax([]int{1, 2, 3, 4, -8, 2}))
}
