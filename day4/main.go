package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sections []int

type AssignmentPair struct {
	ElfOne Sections
	ElfTwo Sections
}

func makeSections(min, max int) Sections {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (a *AssignmentPair) IsFullyContained() bool {
	if a.ElfOne[0] <= a.ElfTwo[0] && a.ElfOne[len(a.ElfOne)-1] >= a.ElfTwo[len(a.ElfTwo)-1] {
		return true
	} else if a.ElfOne[0] >= a.ElfTwo[0] && a.ElfOne[len(a.ElfOne)-1] <= a.ElfTwo[len(a.ElfTwo)-1] {
		return true
	}

	return false
}

func (a *AssignmentPair) Overlaps() bool {
	for _, elfOneSection := range a.ElfOne {
		for _, elfTwoSection := range a.ElfTwo {
			if elfOneSection == elfTwoSection {
				return true
			}
		}
	}
	return false
}

func main() {
	assignmentPairs := []AssignmentPair{}

	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		assignmentpair := AssignmentPair{}
		elves := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(elves[0], "-")
		elf2 := strings.Split(elves[1], "-")
		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(elf1[1])
		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(elf2[1])
		assignmentpair.ElfOne = makeSections(elf1Start, elf1End)
		assignmentpair.ElfTwo = makeSections(elf2Start, elf2End)
		assignmentPairs = append(assignmentPairs, assignmentpair)
	}

	var totalFullyContainedPairs int
	var totalOverlappingPairs int
	for _, pair := range assignmentPairs {
		if pair.IsFullyContained() {
			totalFullyContainedPairs++
		}
		if pair.Overlaps() {
			totalOverlappingPairs++
		}
	}

	log.Printf("Total fully contained pairs: %d", totalFullyContainedPairs)
	log.Printf("Total overlapping pairs: %d", totalOverlappingPairs)
}
