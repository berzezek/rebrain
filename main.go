package main

import "fmt"

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(increment()())
	fmt.Println(increment()())
	fmt.Println(increment()())
	fmt.Println(increment()())
}
