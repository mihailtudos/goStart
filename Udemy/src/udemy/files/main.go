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

	total := 0
	for _, file := range files {
		//fmt.Printf("%d. %s\n", i+1, file.Name())
		if file, err := os.Stat(filepath.Join(args[0], file.Name())); err == nil {
			if file.Size() == 0 {
				total += len(file.Name()) + 1
			}
		} else {
			fmt.Println("Something went wrong.")
		}
	}

	fmt.Printf("%d bytes used.\n", total)
	names := make([]byte, 0, total)

	for _, file := range files {
		//fmt.Printf("%d. %s\n", i+1, file.Name())
		if file, err := os.Stat(filepath.Join(args[0], file.Name())); err == nil {
			if file.Size() == 0 {
				names = append(names, file.Name()...)
				names = append(names, '\n')
			}
		} else {
			fmt.Println("Something went wrong.")
		}
	}
	fmt.Printf("%s", names)

	err = os.WriteFile("empty-file.txt", names, 0644)
	if err != nil {
		fmt.Println("Something went wrong when writing to file.", err)
	}
}
