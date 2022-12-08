package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type Rucksack struct {
	CompartmentOne []Item
	CompartmentTwo []Item
}

type Item struct {
	Type     rune
	Priority int
}

func (i *Item) Add(itemType rune) {
	i.Type = itemType

	if unicode.IsLower(i.Type) {
		i.Priority = strings.IndexRune("abcdefghijklmnopqrstuvwxyz", i.Type) + 1
	} else {
		i.Priority = strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", i.Type) + 27
	}
}

func (r *Rucksack) DuplicateItems() []Item {
	var duplicateItems []Item
	for _, c1Item := range r.CompartmentOne {
		for _, c2Item := range r.CompartmentTwo {
			if c1Item.Type == c2Item.Type {
				duplicateItems = append(duplicateItems, c1Item)
			}
		}
	}
	return duplicateItems
}

func main() {
	rucksack := Rucksack{}

	data, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		for i, rune := range scanner.Text() {
			item := Item{}
			item.Add(rune)

			if i <= (len(scanner.Text())-1)/2 {
				rucksack.CompartmentOne = append(rucksack.CompartmentOne, item)
			} else {
				rucksack.CompartmentTwo = append(rucksack.CompartmentTwo, item)
			}
		}
	}
}
