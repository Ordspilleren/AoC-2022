package main

import (
	"bufio"
	"log"
	"os"
)

type Tree struct {
	Height  int
	Visible bool
}

type Trees [][]*Tree

func (t *Tree) VisibleFromRow(trees []*Tree) bool {
	visible := true
	for i := range trees {
		if trees[i].Height >= t.Height {
			visible = false
			break
		}
	}
	return visible
}

func (t *Tree) VisibleFromColumn(trees Trees, currentTreeIndex int) bool {
	visible := true
	for i := range trees {
		if trees[i][currentTreeIndex].Height >= t.Height {
			visible = false
			break
		}
	}
	return visible
}

func (t Trees) CalculateVisibility() {
	for treeRow := range t {
		for tree := range t[treeRow] {
			currentTree := t[treeRow][tree]

			if treeRow == 0 || treeRow == len(t)-1 {
				currentTree.Visible = true
				continue
			}

			if tree == 0 || tree == len(t[treeRow])-1 {
				currentTree.Visible = true
				continue
			}

			leftTrees := t[treeRow][:tree]
			rightTrees := t[treeRow][tree+1:]

			if len(rightTrees) >= 2 {
				log.Print(rightTrees[0])
			}

			topTrees := t[:treeRow]
			bottomTrees := t[treeRow+1:]

			currentTree.Visible = currentTree.VisibleFromRow(leftTrees) || currentTree.VisibleFromRow(rightTrees) ||
				currentTree.VisibleFromColumn(topTrees, tree) || currentTree.VisibleFromColumn(bottomTrees, tree)
		}
	}
}

func main() {
	trees := Trees{}

	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	currentRow := 0
	for scanner.Scan() {
		row := []*Tree{}
		for _, rune := range scanner.Text() {
			treeHeight := int(rune - '0')
			row = append(row, &Tree{Height: treeHeight})
		}
		trees = append(trees, row)
		currentRow++
	}

	trees.CalculateVisibility()

	nVisible := 0
	for treeRow := range trees {
		for tree := range trees[treeRow] {
			if trees[treeRow][tree].Visible {
				nVisible++
			}
		}
	}
	log.Printf("Number of visible trees: %d", nVisible)
}
