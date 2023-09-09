package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	sortToFile2()
}

func sortToFile2() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(items)

	var data []byte
	for i, s := range items {
		s = strconv.Itoa(i) + ". " + s
		data = append(data, s...)
		data = append(data, '\n')
	}

	err := ioutil.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sortAndWriteToFile() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)
	items := []byte(nil)
	for _, v := range args {
		items = append(items, v...)
		items = append(items, '\n')
	}

	err := os.WriteFile("sorted.txt", items, 0644)
	if err != nil {
		fmt.Println("Something went wrong when writing to file.", err)
	}
}

func writeToFile() {
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
