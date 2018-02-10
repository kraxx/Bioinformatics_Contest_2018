/*********************************************************
Bioinformatics Contest 2018 : Problem 1-2 (Hard)
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

	// If pure fermentation costs more than aerobic + fermentation;
	// cost of 38ATP through fermentation > cost of 38ATP through aerobic + fermentation
	if glucose * 19 > glucose + (oxygen * 6) {
		for {
			// Always buy 1 glucose first. Oxygen won't work without glucose
			if dollars-glucose >= 0 {
				dollars -= glucose
				gluCount++
			// Even if we can't afford any more oxygen, we'll continue to buyup glucose with excess dollars until we're out
			} else {
				break
			}
			// Buy up to 6 oxygen. 6 mol oxygen to 1 mol glucose
			oxyTmp := dollars / oxygen
			if oxyTmp >= 6 {
				dollars -= 6 * oxygen
				oxyCount += 6
			} else {
				dollars -= oxyTmp * oxygen
				oxyCount += oxyTmp
			}
		} 
		// Just fermentation, it's cheaper
	} else {

		gluCount = dollars / glucose
	}
	producedATP := (6 * float64(oxyCount)) + (2 * float64(gluCount)) 
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
		outputFile, err := os.Create(strings.TrimSuffix(file, ".txt") + ".hardout2.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file")
			os.Exit(1)
		}
		// Buffer to hold our answer string
		var bufToPrint string = ""
		for scanner.Scan() {
			str := strings.Split(scanner.Text(), " ")
			glucose, err := strconv.Atoi(str[0])
			oxygen, err := strconv.Atoi(str[1])
			dollars, err := strconv.Atoi(str[2])
			if err != nil {

			}
			bufToPrint += findMaxATP(glucose, oxygen, dollars) + "\n"
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		outputFile.WriteString(bufToPrint)
	}
}
