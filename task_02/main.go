package main

import (
	"fmt"
	"math"
)

func main() {
	A := new(int)
	B := 5
	A = &B
	fmt.Println(*A)

	R := new(float64)
	*R = 35 / math.Pi
	S := math.Pi * math.Pow(*R, 2)
	S = math.Floor(S*100) / 100
	fmt.Println(S)
}
