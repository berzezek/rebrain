package main

import "fmt"

func main() {
	personReader := map[string]map[string]int{
		"john": {
			"books":      2,
			"periodical": 3,
		},
		"alex": {
			"books":      0,
			"periodical": 2,
		},
		"mike": {
			"books":      3,
			"periodical": 2,
		},
	}

	personReader["maria"] = map[string]int{
		"books":      2,
		"periodical": 1,
	}

	personReader["jim"] = map[string]int{
		"books":      0,
		"periodical": 1,
	}

	countReader := 0
	for _, b := range personReader {
		if b["books"] > 0 {
			countReader++
		}
	}
	fmt.Println(countReader)
}
