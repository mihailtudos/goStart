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
	sliceLiterals()
}

func sliceLiterals() {
	names := []string{"Petru", "Marat", "Gabi"}
	distances := []int{10, 20, 22, 30, 55}
	data := []uint8{'b', 's', '2', ',', '.'}
	ratios := []float64{1.23, 0.8}
	alive := []bool{true, true, false, true}

	fmt.Printf("names:     %T %d %t\n", names, len(names), nil == names)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), nil == distances)
	fmt.Printf("data:      %T %d %t\n", data, len(data), nil == data)
	fmt.Printf("ratios:    %T %d %t\n", ratios, len(ratios), nil == ratios)
	fmt.Printf("alive:     %T %d %t\n", alive, len(alive), nil == alive)
}

func zeroSlices() {
	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alive     []bool
	)

	names = []string{}
	distances = []int{}
	data = []uint8{}
	ratios = []float64{}
	alive = []bool{}

	fmt.Printf("names:     %T %d %t\n", names, len(names), nil == names)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), nil == distances)
	fmt.Printf("data:      %T %d %t\n", data, len(data), nil == data)
	fmt.Printf("ratios:    %T %d %t\n", ratios, len(ratios), nil == ratios)
	fmt.Printf("alive:     %T %d %t\n", alive, len(alive), nil == alive)
}

func nilSlices() {
	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alive     []bool
	)

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
