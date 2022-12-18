package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	defer logTime()
}

func getIn(pathIn, pathOut string) {
	fileIn, _ := os.Open(pathIn)
	defer func(fileIn *os.File) {
		err := fileIn.Close()
		if err != nil {

		}
	}(fileIn)
	fileOut, _ := os.OpenFile(pathOut, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	lineNum := 0
	buffNum := 0
	defer func(fileOut *os.File) {
		fmt.Println("Количество строк:", lineNum)
		fmt.Println("Количество байт:", buffNum)
		err := fileOut.Close()
		if err != nil {
		}
	}(fileOut)

	s := bufio.NewScanner(fileIn)
	w := bufio.NewWriter(fileOut)
	for s.Scan() {
		lineNum++
		lineStr := strconv.Itoa(lineNum)
		_, err := w.WriteString(lineStr + ". " + s.Text() + "\n")
		buffNum += w.Buffered()
		if err != nil {
			return
		}
	}
	err := w.Flush()
	if err != nil {
		return
	}
}

func logTime() {
	start := time.Now()
	pathIn := "./task_06/in.txt"
	pathOut := "./task_06/out.txt"
	getIn(pathIn, pathOut)
	t := time.Now()
	fmt.Println(t.Sub(start))

}
