/*********************************************************
Bioinformatics Contest 2018 = Problem 1-1 (Easy)
*********************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMaxATP(glucose, oxygen, dollars float64) string {
	val := (dollars / (glucose + (6 * oxygen))) * 38
	return strconv.FormatFloat(val, 'f', -1, 64)
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
		outputFile, err := os.Create(strings.TrimSuffix(file, ".txt") + "out.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file")
			os.Exit(1)
		}
		for scanner.Scan() {
			str := strings.Split(scanner.Text(), " ")
			glucose, err := strconv.ParseFloat(str[0], 64)
			oxygen, err := strconv.ParseFloat(str[1], 64)
			dollars, err := strconv.ParseFloat(str[2], 64)
			if err != nil {

			}
			outputFile.WriteString(findMaxATP(glucose, oxygen, dollars) + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
