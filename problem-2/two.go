/*********************************************************
Bioinformatics Contest 2018 : Problem 2-2
*********************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	phenylAlanine = iota
	leucine       = iota
	serine        = iota
	tyrosine      = iota
	cysteine      = iota
	tryptophan    = iota
	proline       = iota
	histidine     = iota
	glutamine     = iota
	arginine      = iota
	isoleucine    = iota
	methionine    = iota
	threonine     = iota
	asparagine    = iota
	lysine        = iota
	valine        = iota
	alanine       = iota
	asparticAcid  = iota
	glutamicAcid  = iota
	glycine       = iota
	stopCodon     = iota
)

var AA = map[string]int{
	"UUU": phenylAlanine,
	"UUC": phenylAlanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UAA": stopCodon,
	"UAG": stopCodon,
	"UGU": cysteine,
	"UGC": cysteine,
	"UGA": stopCodon,
	"UGG": tryptophan,
	"CUU": leucine,
	"CUC": leucine,
	"CUA": leucine,
	"CUG": leucine,
	"CCU": proline,
	"CCC": proline,
	"CCA": proline,
	"CCG": proline,
	"CAU": histidine,
	"CAC": histidine,
	"CAA": glutamine,
	"CAG": glutamine,
	"CGU": arginine,
	"CGC": arginine,
	"CGA": arginine,
	"CGG": arginine,
	"AUU": isoleucine,
	"AUC": isoleucine,
	"AUA": isoleucine,
	"AUG": methionine,
	"ACU": threonine,
	"ACC": threonine,
	"ACA": threonine,
	"ACG": threonine,
	"AAU": asparagine,
	"AAC": asparagine,
	"AAA": lysine,
	"AAG": lysine,
	"AGU": serine,
	"AGC": serine,
	"AGA": arginine,
	"AGG": arginine,
	"GUU": valine,
	"GUC": valine,
	"GUA": valine,
	"GUG": valine,
	"GCU": alanine,
	"GCC": alanine,
	"GCA": alanine,
	"GCG": alanine,
	"GAU": asparticAcid,
	"GAC": asparticAcid,
	"GAA": glutamicAcid,
	"GAG": glutamicAcid,
	"GGU": glycine,
	"GGC": glycine,
	"GGA": glycine,
	"GGG": glycine,
}

func changeCodon(sequence string, l, r int) int {

	for i := l - 1; i < r; i++ {

		// Skip start codon
		if i <= 2 {
			continue
		}
		// Block represents index of beginning of a codon
		block := (i / 3) * 3
		// index of said block
		blockIdx := i % 3
		codon := sequence[block : block+3]

		// Almost all codons will change amino acids if their first nucleotide changes
		if blockIdx == 0 {

			// Certain arginine codons are the only ones allowed to swap their first nucleotide
			if codon == "CGA" || codon == "CGG" || codon == "AGA" || codon == "AGG" {
				return i + 1
			}
			continue
			// Same applies to 2nd nucleotide
		} else if blockIdx == 1 {
			// Stop codons UAA and UGA are the only exemptions
			if codon == "UAA" || codon == "UGA" {
				return i + 1
			}
			continue
		} else {

			// All codons are fine with swapping their 3rd nucleotide except these two
			if codon == "UGA" || codon == "UGG" {
				continue
			}
			return i + 1
		}
	}
	// No swaps
	return -1
}

func findMinChanges(scanner *bufio.Scanner) int {

	scanner.Scan()
	sequence := strings.Trim(scanner.Text(), " ")
	scanner.Scan()
	numSites, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))

	lastChangedIdx := 0
	changes := 0

	for j := 0; j < numSites; j++ {
		scanner.Scan()
		siteRange := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(siteRange[0])
		r, _ := strconv.Atoi(siteRange[1])

		// Can't change start codon
		if r <= 3 {
			// Finish reading, return error
			for k := j; k < numSites - 1; k++ {
				scanner.Scan()
			}
			return -1
		}
		// If a nucleotide has been swapped, move on
		if l <= lastChangedIdx && lastChangedIdx <= r {
			continue
		}
		lastChangedIdx = changeCodon(sequence, l, r)
		if lastChangedIdx != -1 {
			changes++
		} else {
			// Finish reading, return error
			for k := j; k < numSites; k++ {
				scanner.Scan()
			}
			return -1
		}
	}
	if changes == 0 {
		return -1
	}
	return changes
}

func main() {
	if len(os.Args) == 2 {

		file := os.Args[1]
		openFile, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file")
			os.Exit(1)
		}
		// outputFile, err := os.Create(strings.TrimSuffix(file, ".txt") + "out.txt")
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, "Error creating output file")
		// 	os.Exit(1)
		// }

		scanner := bufio.NewScanner(openFile)
		scanner.Scan()
		numTests, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))

		for i := 0; i < numTests; i++ {
			wew := findMinChanges(scanner)
			fmt.Println(wew)
		}
	}
}
