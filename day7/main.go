package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	Name        string
	Size        int
	Directories []*Directory
	Files       []File
	Parent      *Directory
}

type File struct {
	Name string
	Size int
}

type Filesystem []*Directory

var totalSizes int
var candidatesForDeletion []int

func (d *Directory) PropagateSize() {
	if d.Parent != nil {
		d.Parent.Size += d.Size
	}
}

func (d *Directory) CalculateSize() {
	d.PropagateSize()
	for _, file := range d.Files {
		d.Size += file.Size
	}
	for _, directory := range d.Directories {
		directory.CalculateSize()
		d.Size += directory.Size
	}
}

func (d *Directory) FindCulprits() {
	if d.Size <= 100000 {
		totalSizes += d.Size
		log.Printf("Directory %s has size %d", d.Name, d.Size)
	}
	for _, directory := range d.Directories {
		directory.FindCulprits()
	}
}

func (d *Directory) Print(depth int) {
	for _, directory := range d.Directories {
		var depthString strings.Builder
		for i := 0; i < depth; i++ {
			depthString.WriteString("    ")
		}
		log.Printf("%s%s (dir, size=%d)", depthString.String(), directory.Name, directory.Size)
		for _, file := range directory.Files {
			log.Printf("%s%s%s (file, size=%d)", &depthString, &depthString, file.Name, file.Size)
		}
		directory.Print(depth + 1)
	}
}

func (f Filesystem) Iterate() {
	for _, directory := range f {
		directory.CalculateSize()
		log.Printf("%s (dir, size=%d)", directory.Name, directory.Size)
		for _, file := range directory.Files {
			log.Printf("    %s (file, size=%d)", file.Name, file.Size)
		}
		directory.Print(1)
		directory.FindCulprits()
	}
}

func (d *Directory) FindSmallestForDeletion(availableSpace int, usedSpace int, unusedSpaceNeeded int) {
	for _, directory := range d.Directories {
		if (availableSpace-usedSpace)+directory.Size >= unusedSpaceNeeded {
			candidatesForDeletion = append(candidatesForDeletion, directory.Size)
		}
		directory.FindSmallestForDeletion(availableSpace, usedSpace, unusedSpaceNeeded)
	}
}

func (f Filesystem) Cleanup(availableSpace int, unusedSpaceNeeded int) {
	for _, directory := range f {
		directory.FindSmallestForDeletion(availableSpace, f[0].Size, unusedSpaceNeeded)
	}
	sort.Ints(candidatesForDeletion)
	log.Printf("Directory to delete: %d", candidatesForDeletion)
}

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	filesystem := Filesystem{}
	var currentDir *Directory
	for scanner.Scan() {
		splitOutput := strings.Split(scanner.Text(), " ")
		if strings.Contains(scanner.Text(), "$ cd") {
			dir := splitOutput[2]
			if dir == ".." {
				currentDir = currentDir.Parent
				continue
			}
			var duplicateDir *Directory
			if currentDir != nil {
				for i := range currentDir.Directories {
					if currentDir.Directories[i].Name == dir {
						duplicateDir = currentDir.Directories[i]
					}
				}
			}
			if duplicateDir == nil {
				directory := Directory{Name: dir, Parent: currentDir}
				if currentDir != nil {
					currentDir.Directories = append(currentDir.Directories, &directory)
				} else {
					filesystem = append(filesystem, &directory)
				}
				currentDir = &directory
			} else {
				currentDir = duplicateDir
			}
		} else if strings.ContainsAny(splitOutput[0], "1234567890") {
			fileName := splitOutput[1]
			fileSize, _ := strconv.Atoi(splitOutput[0])
			file := File{Name: fileName, Size: fileSize}
			currentDir.Files = append(currentDir.Files, file)
		}
	}
	filesystem.Iterate()
	log.Printf("Total sizes: %d", totalSizes)

	filesystem.Cleanup(70000000, 30000000)
}
