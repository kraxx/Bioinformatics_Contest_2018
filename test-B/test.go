package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findSubstrings(str string, sub string) (ret []int) {
	l := len(sub)
	for i := range str[:len(str)-l] {
		if str[i:i+l] == sub {
			ret = append(ret, i+1)
		}
	}
	return
}

func main() {
	if len(os.Args) == 2 {

		file := os.Args[1]
		openFile, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file")
			os.Exit(1)
		}
		scanner := bufio.NewScanner(openFile)
		scanner.Scan() // Scan over first line, irrelevant to me; returns true if scanned, false if done
		outputFile, err := os.Create("output.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file")
			os.Exit(1)
		}
		for scanner.Scan() {
			str := scanner.Text()
			scanner.Scan()
			sub := scanner.Text()
			intArr := findSubstrings(str, sub)
			for i, v := range intArr {
				if i > 0 {
					outputFile.WriteString(" ")
				}
				outputFile.WriteString(strconv.Itoa(v))
			}
			outputFile.WriteString("\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
