package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	"rebrain/03_Module/task_01/hw"
)

func main() {
	fmt.Println(xstrings.Shuffle(hw.City()))
	fmt.Println(hw.Digit())
}
