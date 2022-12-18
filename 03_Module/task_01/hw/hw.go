package hw

import "rebrain/03_Module/task_01/wordz"

var Cities = []string{
	"Moscow",
	"Tashkent",
	"Kiev",
	"Minsk",
	"Astana",
	"Bishkek",
}

func City() string {
	wordz.Words = Cities
	s := wordz.Random()
	return s
}

var Digits = []string{
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
}

func Digit() string {
	wordz.Words = Digits
	s := wordz.Random()
	return s
}
