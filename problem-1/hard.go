/*********************************************************
Bioinformatics Contest 2018 = Problem 1-2 (Hard)
*********************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMaxATP(glucose, oxygen, dollars int) string {
	gluCount := 0
	oxyCount := 0
	for {
		if dollars-glucose >= 0 {
			dollars -= glucose
			gluCount++
		} else {
			break
		}
		for i := 0; i < 6; i++ {
			if dollars-oxygen >= 0 {
				dollars -= oxygen
				oxyCount++
			} else {
				break
			}
		}
		if dollars-oxygen <= 0 {
			break
		}
	}
	oxyFloat := float64(oxyCount)
	// fmt.Println(oxyFloat)
	producedATP := 6 * oxyFloat
	if 6*gluCount > oxyCount {
		producedATP += 2
	}
	return strconv.FormatFloat(producedATP, 'f', -1, 64)
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
		outputFile, err := os.Create(strings.TrimSuffix(file, ".txt") + ".hardout.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file")
			os.Exit(1)
		}
		var bufToPrint string = ""
		for scanner.Scan() {
			str := strings.Split(scanner.Text(), " ")
			// glucose, err := strconv.ParseFloat(str[0], 64)
			// oxygen, err := strconv.ParseFloat(str[1], 64)
			// dollars, err := strconv.ParseFloat(str[2], 64)
			glucose, err := strconv.Atoi(str[0])
			oxygen, err := strconv.Atoi(str[1])
			dollars, err := strconv.Atoi(str[2])
			if err != nil {

			}
			// outputFile.WriteString(findMaxATP(glucose, oxygen, dollars) + "\n")
			bufToPrint += findMaxATP(glucose, oxygen, dollars) + "\n"
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		outputFile.WriteString(bufToPrint)
	}
}
