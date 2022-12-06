package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	workDays := days[:5]
	fmt.Println(workDays)
	weekends := make([]string, 2)
	weekends = days[len(workDays):]
	fmt.Println(weekends)
	allDays := append(workDays, weekends...)
	fmt.Println(allDays)
}
