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

func mustDoubleSwap(codon string) bool {

	return codon == "CUU" || codon == "CUC" || // Leucine
			codon == "CGU" || codon == "CGC" || // Arginine
			AA[codon] == serine
}

func canDoubleSwap(codon string) bool {

	return mustDoubleSwap(codon) || codon == "UUA" || codon == "UUG" || // Leucine
			codon == "AGA" || codon == "AGG" || // Arginine
			codon == "UAG" || codon == "UGA" // Stop codons
}

func canSingleSwap(codon string) bool {
	return !mustDoubleSwap(codon) || AA[codon] != methionine || AA[codon] != tryptophan
}

func allocateSitesArray(scanner *bufio.Scanner, numSites int) ([][]int, bool) {

	sites := make([][]int, numSites)
	for i := range sites {
		scanner.Scan()
		siteRange := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(siteRange[0])
		r, _ := strconv.Atoi(siteRange[1])

		// fmt.Println(l, r)

		// Can't change start codon
		if r <= 3 {
			return sites, false
		}
		sites[i] = make([]int, 2)
		sites[i][0], sites[i][1] = l, r
	}
	return sites, true
}

func changeCodon(sequence string, l, r int) (int, bool) {

	// If we can and must, we double swap
	ds := false
	dsIdx := 0

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


		if canDoubleSwap(codon) {
			if AA[codon] == serine {

				if r % 3 > 0 {
					ds = true
					dsIdx = (r % 3) + 1
				}

			} else if AA[codon] == arginine || AA[codon] == leucine || AA[codon] == stopCodon {

				if r % 3 > 1 {
					ds = true
					dsIdx = (r % 3) + 1					
				}

			}
		}

		// Almost all codons will change amino acids if their first nucleotide changes
		if blockIdx == 0 {

			// Certain arginine codons are the only ones allowed to swap their first nucleotide
			if codon == "CGA" || codon == "CGG" || codon == "AGA" || codon == "AGG" {
				return i + 1, false
			}
			continue
			// Same applies to 2nd nucleotide
		} else if blockIdx == 1 {
			// Stop codons UAA and UGA are the only exemptions
			if codon == "UAA" || codon == "UGA" {
				return i + 1, false
			}
			continue
		} else {

			// All codons are fine with swapping their 3rd nucleotide except these two
			if codon == "UGA" || codon == "UGG" {
				continue
			}
			return i + 1, false
		}

	}
	if ds == true {
		return dsIdx, true
	}
	// No swaps
	return -1, false
}

func findMinChanges(scanner *bufio.Scanner) int {

	scanner.Scan()
	sequence := strings.Trim(scanner.Text(), " ")
	fmt.Println(len(sequence))
	scanner.Scan()
	numSites, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))

	// Allocate 2D int array to hold all of our site ranges
	sites, passed := allocateSitesArray(scanner, numSites)
	if passed == false {
		return -1
	}

	lastChangedIdx := 0
	doubleSwapped := false
	changes := 0

	for i := 0; i < numSites; i++ {

		l, r := sites[i][0], sites[i][1]

/*
		// If a nucleotide has been swapped, move on
		if l <= lastChangedIdx && lastChangedIdx <= r {
			continue
		}
*/
		// If rightmost nucleotide can overlap onto another site
		if i < numSites - 1 {
			rBlockIdx := r % 3
			nextL := sites[i+1][0]

			// If next siteRange's Leftmost index can overlap with our current rightmost index
			if (rBlockIdx == 1 && nextL - r <= 2) || (rBlockIdx == 2 && nextL - r == 1) {
				// Rightmost codon
				rBlock := ((r - 1) / 3) * 3
				rCodon := sequence[rBlock:rBlock+3]

				nextR := sites[i+1][1]
				// If rightmost siteRange of nextR is restricted to range of the last codon
				if (rBlockIdx == 1 && nextR - r <= 2) || (rBlockIdx == 2 && nextR - r == 1) {

					if mustDoubleSwap(rCodon) == true {
						if AA[rCodon] == serine {
							// We need specific nucleotides on left and right sites
							if rBlockIdx != 1 || nextL % 3 != 2 {
								return -1
							} else {
								changes += 2
								i++
								continue
							}
						} else if AA[rCodon] == leucine || AA[rCodon] == arginine {
							// If we are restricted to using only 2nd nucleotide of left siteRange
							if rBlockIdx == 2 && r - l == 0 {
								return -1
							} else {
								changes += 2
								i++
								continue
							}
						}
					}
				}

				// Perform a double swap if possible
				if canDoubleSwap(rCodon) == true {

					if AA[rCodon] == leucine || AA[rCodon] == arginine {
						// If we are restricted to using only 2nd nucleotide of left siteRange
						if !(rBlockIdx == 2 && r - l == 0) {
							changes += 2
							i++
							continue
						}
					} else if rCodon == "UAG" || rCodon == "UGA" {
						changes += 2
						i++
						continue
					}
				}

			}			
		}

		lastChangedIdx, doubleSwapped = changeCodon(sequence, l, r)
		// If a nucleotide can be swapped, we're good
		if doubleSwapped == true {
			changes += 2
		} else if lastChangedIdx != -1 {
			changes++
		} else {
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
		// outputFile, err := os.Create(strings.TrimSuffix(file, ".txt") + ".out.txt")
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, "Error creating output file")
		// 	os.Exit(1)
		// }

		scanner := bufio.NewScanner(openFile)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)
		scanner.Scan()
		numTests, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))

		// Buffer to hold our answer string
		// var bufToPrint string = ""

		for i := 0; i < numTests; i++ {
			wew := findMinChanges(scanner)
			fmt.Println(wew)
			// bufToPrint += strconv.Itoa(wew) + "\n"
		}
		// outputFile.WriteString(bufToPrint)
	}
}
