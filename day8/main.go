package main

import (
	"bufio"
	"log"
	"os"
)

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

type Trees [][]*Tree

func (t *Tree) VisibleFromRow(trees []*Tree, reverse bool) (bool, int) {
	viewingDistance := 0
	for i := range trees {
		if reverse {
			i = len(trees) - 1 - i
		}
		viewingDistance++
		if trees[i].Height >= t.Height {
			return false, viewingDistance
		}
	}
	return true, viewingDistance
}

func (t *Tree) VisibleFromColumn(trees Trees, currentTreeIndex int, reverse bool) (bool, int) {
	viewingDistance := 0
	for i := range trees {
		if reverse {
			i = len(trees) - 1 - i
		}
		viewingDistance++
		if trees[i][currentTreeIndex].Height >= t.Height {
			return false, viewingDistance
		}
	}
	return true, viewingDistance
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
			topTrees := t[:treeRow]
			bottomTrees := t[treeRow+1:]

			visibleFromLeft, leftViewingDistance := currentTree.VisibleFromRow(leftTrees, true)
			visibleFromRight, rightViewingDistance := currentTree.VisibleFromRow(rightTrees, false)
			visibleFromTop, topViewingDistance := currentTree.VisibleFromColumn(topTrees, tree, true)
			visibleFromBottom, bottomViewingDistance := currentTree.VisibleFromColumn(bottomTrees, tree, false)

			currentTree.ScenicScore = leftViewingDistance * rightViewingDistance * topViewingDistance * bottomViewingDistance
			currentTree.Visible = visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBottom

			log.Printf("Row: %d\nHeight: %d\nScenic Score: %d\n\n", treeRow, currentTree.Height, currentTree.ScenicScore)
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
	higestScenicScore := 0
	for treeRow := range trees {
		for tree := range trees[treeRow] {
			if trees[treeRow][tree].Visible {
				nVisible++
			}
			if trees[treeRow][tree].ScenicScore > higestScenicScore {
				higestScenicScore = trees[treeRow][tree].ScenicScore
			}
		}
	}
	log.Printf("Number of visible trees: %d\nHighest Scenic Score: %d", nVisible, higestScenicScore)
}
