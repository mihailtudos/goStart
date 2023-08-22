package main

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	housingPrices()
}

func housingPrices() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
	type row [5]string
	head := row(strings.Split(header, separator))
	var (
		locations                  []string
		sizes, beds, baths, prices []int
	)
	rows := strings.Split(data, "\n")

	for _, v := range rows {
		for i, r := range strings.Split(v, separator) {
			var (
				v   int
				err error
			)
			if v, err = strconv.Atoi(r); err != nil {
				v = 0
			}
			switch i {
			case 0:
				locations = append(locations, r)
			case 1:
				sizes = append(sizes, v)
			case 2:
				beds = append(beds, v)
			case 3:
				baths = append(baths, v)
			case 4:
				prices = append(prices, v)
			}
		}
	}

	for i := range head {
		fmt.Printf("%-15s", head[i])
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d\n", prices[i])
	}
	fmt.Printf("%s\n%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", strings.Repeat("=", 75), "", avgInts(sizes), avgInts(beds), avgInts(baths), avgInts(prices))
}

func avgInts(data []int) float64 {
	var sum float64
	for i := range data {
		sum += float64(data[i])
	}
	return sum / float64(len(data))
}

func appendAndSort() {
	inputs := os.Args[1:]
	if len(inputs) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var numbers []int
	for _, v := range inputs {
		if number, err := strconv.Atoi(v); err == nil {
			numbers = append(numbers, number)
		}
	}
	sort.Ints(numbers)
	fmt.Printf("%v", numbers)
}

func fixSlice() {
	toppings := []string{"black olives", "green peppers"}

	toppings = append(toppings, "extra cheese", "onions")

	fmt.Printf("pizza       : %s\n", toppings)
}

func appendNil() {
	var toppings []string
	var departures []time.Time
	var graduationYears []int
	var states []bool

	toppings = append(toppings, "Pineapple", "Ham", "Pepperoni", "Cheese", "Mushrooms")
	departures = append(departures, time.Now(), time.Now().Add(time.Hour*2), time.Now().AddDate(0, 1, 0))
	graduationYears = append(graduationYears, 2020, 2021, 2022)
	states = append(states, true, true, true, false)

	fmt.Printf("pizza toppings:   %v\n", toppings)
	fmt.Printf("graduation years: %v\n", graduationYears)
	fmt.Printf("states: 	  %v\n", states)
	fmt.Printf("states: 	  %v\n", departures)
}

func compareSlices2() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, png...)

	if bytes.Compare(png, header) != 0 {
		fmt.Printf("They are equal")
	} else {
		fmt.Printf("They are not equal")
	}
}

func compareSlices() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesB)
	sort.Strings(namesA)
	fmt.Printf("%#v\n", namesA)
	fmt.Printf("%#v\n", namesB)

	if len(namesA) == len(namesB) {
		equal := true
		for i, v := range namesA {
			if namesB[i] != v {
				equal = false
				break
			}
		}

		if equal {
			fmt.Printf("They are equal.\n")
			return
		}
	}

	fmt.Printf("They not are equal.\n")
}

func fixSlicesIssue() {
	names := []string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	books := []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	sort.Ints(nums[:])

	fmt.Printf("%s\n", strings.Join(names, " and "))

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}

func updateToSlices() {
	names := []string{"Einstein", "Tesla", "Shepard"}
	distances := []int{50, 40, 75, 30, 125}
	data := []byte{'H', 'E', 'L', 'L', 'O'}
	ratios := []float64{3.14145}
	alives := []bool{true, false, true, false}
	zero := []byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)
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
