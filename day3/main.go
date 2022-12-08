package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type Group struct {
	Rucksacks []Rucksack
	Badge     Item
}

type Rucksack struct {
	CompartmentOne []Item
	CompartmentTwo []Item
}

type Item struct {
	Type     rune
	Priority int
}

func (g *Group) AssignBadge() {
	itemList := make(map[rune]int)
	for _, rucksack := range g.Rucksacks {
		var addedItems []rune
		for _, item := range rucksack.Items() {
			var alreadyAdded bool
			for _, ai := range addedItems {
				if ai == item.Type {
					alreadyAdded = true
				}
			}
			if !alreadyAdded {
				addedItems = append(addedItems, item.Type)
				itemList[item.Type] += 1
			}
		}
	}

	for k, v := range itemList {
		if v == 3 {
			item := Item{}
			item.Add(k)
			g.Badge = item
		}
	}
}

func (r *Rucksack) Items() []Item {
	return append(r.CompartmentOne, r.CompartmentTwo...)
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
			if c2Item.Type == c1Item.Type {
				duplicateItems = append(duplicateItems, c1Item)
				break
			}
		}
	}
	return duplicateItems
}

func main() {
	groups := []Group{}

	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	currentGroup := Group{}
	counter := 0
	for scanner.Scan() {
		if counter != 3 {
			counter++
		}

		rucksack := Rucksack{}
		for i, rune := range scanner.Text() {
			item := Item{}
			item.Add(rune)

			if i <= (len(scanner.Text())-1)/2 {
				rucksack.CompartmentOne = append(rucksack.CompartmentOne, item)
			} else {
				rucksack.CompartmentTwo = append(rucksack.CompartmentTwo, item)
			}
		}
		currentGroup.Rucksacks = append(currentGroup.Rucksacks, rucksack)

		if counter == 3 {
			groups = append(groups, currentGroup)
			currentGroup = Group{}
			counter = 0
		}
	}

	var groupTotal int
	for _, group := range groups {
		group.AssignBadge()
		groupTotal += group.Badge.Priority
		log.Printf("Badge: %c", group.Badge)

		var total int
		for _, rucksack := range group.Rucksacks {
			duplicates := rucksack.DuplicateItems()
			log.Printf("Type: %c\nPriority: %d\n", duplicates[0].Type, duplicates[0].Priority)
			total += duplicates[0].Priority
		}
		log.Printf("Total: %d", total)
	}

	log.Printf("Group Total: %d", groupTotal)
}
