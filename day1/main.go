package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Food          []int
	TotalCalories int
}

func main() {
	elves := []Elf{}
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	currentElf := Elf{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			elves = append(elves, currentElf)
			currentElf = Elf{}
			continue
		}

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		currentElf.Food = append(currentElf.Food, value)
	}

	for i := range elves {
		sum := 0

		for _, value := range elves[i].Food {
			sum += value
		}

		elves[i].TotalCalories = sum
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})

	for _, elf := range elves {
		log.Printf("Total Calories: %d\n", elf.TotalCalories)
	}

	log.Printf("The top three elves are carrying %d calories", elves[0].TotalCalories+elves[1].TotalCalories+elves[2].TotalCalories)
}
