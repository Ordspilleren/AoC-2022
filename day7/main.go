package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
type Directory struct {
	Name        string
	Directories []Directory
	Files       []File
}
*/

type Directory map[string][]File

type File struct {
	Name string
	Size string
}

func main() {
	data, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	filesystem := make(Directory)
	var currentDir strings.Builder
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "$ cd") {
			dir := strings.Split(scanner.Text(), " ")[2]
			if dir == ".." {
				currentDir.Reset()
			}
			currentDir.WriteString(dir)
			if _, ok := filesystem[currentDir.String()]; !ok {
				currentDir.WriteString("/")
				filesystem[currentDir.String()] = []File{}
			}
		}
	}

	log.Print(filesystem)
}
