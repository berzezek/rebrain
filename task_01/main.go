package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := "104"
	b := 35
	i, err := strconv.Atoi(a)
	if err == nil {
		fmt.Println(i)
	}

	s := strconv.Itoa(b)
	fmt.Println(s)

	// new task
	var EuropeanVelocity, AmericanVelocity float64
	EuropeanVelocity = 120.4 * 1000 / 3600
	AmericanVelocity = 130 * 1.609 * 1000 / 3600
	EuropeanVelocity = math.Floor(EuropeanVelocity*100) / 100
	AmericanVelocity = math.Floor(AmericanVelocity*100) / 100
	fmt.Println("EuropeanVelocity:", EuropeanVelocity, "\nAmericanVelocity:", AmericanVelocity)
}
