package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide directory path")
		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	emptyFiles := []string(nil)

	for _, file := range files {
		//fmt.Printf("%d. %s\n", i+1, file.Name())
		if file, err := os.Stat(filepath.Join(args[0], file.Name())); err == nil {
			if file.Size() == 0 {
				emptyFiles = append(emptyFiles, file.Name())
			}
		} else {
			fmt.Println("Something went wrong.")
		}
	}

	for i, file := range emptyFiles {
		fmt.Printf("%d. %s\n", i+1, file)
	}
}
