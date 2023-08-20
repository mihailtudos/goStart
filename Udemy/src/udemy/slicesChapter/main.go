package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	nilSlices()
}

func nilSlices() {
	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var alive []bool

	fmt.Printf("names:     %T %d %t\n", names, len(names), nil == names)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), nil == distances)
	fmt.Printf("data:      %T %d %t\n", data, len(data), nil == data)
	fmt.Printf("ratios:    %T %d %t\n", ratios, len(ratios), nil == ratios)
	fmt.Printf("alive:    %T %d %t\n", alive, len(alive), nil == alive)
}

func sliceOfUniques() {
	rand.NewSource(time.Now().UnixMicro())
	size, _ := strconv.Atoi(os.Args[1])
	var uniques []int
	for len(uniques) < size {
		n := rand.Intn(size + 1)
		if slices.Contains(uniques, n) {
			continue
		} else {
			uniques = append(uniques, n)
		}
	}
	sort.Ints(uniques)
	fmt.Printf("%v \n", uniques)
}
