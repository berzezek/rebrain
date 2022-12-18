package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		s, f := openMyFile("./task_07/data/data_out.txt")
		for s.Scan() {
			fmt.Println(s.Text())
		}
		defer func() {
			err := f.Close()
			if err != nil {
				return
			}
		}()
	}()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			fmt.Println(panicErr)
		}
	}()
	scanFile("./task_07/in.txt")
}

func scanFile(path string) {
	f, _ := os.Open(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	s := bufio.NewScanner(f)

	row := 0
	err := os.Mkdir("./task_07/data", 0755)
	if err != nil {
	}

	fileOut, _ := os.OpenFile("./task_07/data/data_out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer func() {
		err := fileOut.Close()
		if err != nil {

		}
	}()
	w := bufio.NewWriter(fileOut)
	for s.Scan() {
		row += 1
		splitS := strings.Split(s.Text(), "|")
		for _, i := range splitS {
			if len(i) == 0 {
				panic(fmt.Sprintf("Empty field on string %d", row))
			}
		}
		_, err := w.WriteString(fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", row, splitS[0], splitS[1], splitS[2]))
		if err != nil {
			return
		}
		errFl := w.Flush()
		if errFl != nil {
			return
		}
	}
}

func openMyFile(path string) (*bufio.Scanner, *os.File) {
	f, _ := os.Open(path)
	s := bufio.NewScanner(f)
	return s, f
}
