package main

import (
	"bytes"
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const size = 1e7

func main() {
	advOpsPractice()
}

func advOpsPractice() {
	// #1: Create a string slice: `names` with a length and
	//     capacity of 5, and print it.
	names := make([]string, 5)
	s.Show("1st step", names)

	// #2: Append the following names to the names slice:
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("2nd step", names)

	//#3: Overwrite the name slice by creating a new slice
	//     using make().

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("3rd step", names)

	// #4: Copy only the first two elements of the following
	//     array to the last two elements of the `names` slice.

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]
	s.Show("4th step", names)

	// #5:  Only copy the last 3 elements of the `names` slice
	//      to a new slice: `clone`.

	clone := make([]string, len(names[2:]), cap(names))
	copy(clone[:], names[2:])

	s.Show("5th step (before append)", clone)
	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	// #6: Slice the `clone` slice between 2nd and 4th (inclusive)
	//     elements into a new slice: `sliced`.
	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")
	s.Show("6th step", clone, sliced)
}

func correctLyrics() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	lyric = append([]string{"yesterday"}, lyric...)
	lyric = append(lyric, lyric[8:13]...)
	lyric = append(lyric[:8], lyric[13:]...)
	fmt.Printf("%s\n", lyric)
}

func sortBackingArray() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	midd := items[len(items)/2-1 : len(items)/2+2]
	sort.Strings(midd)
	fmt.Println()
	fmt.Println("Sorted  :", items)
}

func fixBackingArrayIssue() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// ONLY ADD YOUR CODE HERE
	//
	// Ensure that nums slice never changes even though
	// the mine slice changes.
	mine := append([]int(nil), nums[:3]...)

	// ----------------------------------------

	// DON'T TOUCH THE FOLLOWING CODE
	//
	// This code changes the elements of the nums
	// slice.
	//
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}

func observeCapGrowth() {
	arr := []int(nil)
	prevSize := 1
	for g, i := 0, 0; i < 1e7; i++ {
		if len(arr) == cap(arr) {
			fmt.Printf("len: %d cap: %d grow: %d%%\n", len(arr), cap(arr), g)
			if cap(arr) != 0 {
				prevSize = cap(arr)
			}
		}
		arr = append(arr, i)
		g = (cap(arr) * 100 / prevSize) - 100
	}
	var (
		nums   []int
		oldCap float64
	)
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
}

func observeLenCap() {
	// --- #1 ---
	// 1. create a new slice named: games
	//var games []string
	// 2. print the length and capacity of the games slice
	//fmt.Printf("Len: %d	Cap: %d\n", len(games), cap(games))
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//games := []string{}
	// 4. print the length and capacity of the games slice
	//fmt.Printf("Len: %d	Cap: %d\n", len(games), cap(games))
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//games = append(games, "pacman", "mario", "tetris", "doom")
	// 6. print the length and capacity of the games slice
	//fmt.Printf("Len: %d	Cap: %d\n", len(games), cap(games))
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
	games := []string{"pacman", "mario", "tetris", "doom"}
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	// 2. print its length and capacity along the way (in the loop).
	fmt.Println()
	for i := 0; i < 4; i++ {
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(games[:i]), cap(games[:i]))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	// 2. print the games and the new slice's len and cap

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	// 3. append a new element to the new slice
	zero = append(zero, "super")

	// 4. print the new slice's lens and caps
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	// 5. repeat the last two steps 5 times (use a loop)
	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	fmt.Println()
	games[0] = "command & conquer games"
	zero[0] = "command & conquer"

	fmt.Printf("games : %q\n", games)
	fmt.Printf("zero  : %q\n", zero)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	fmt.Printf("games : c: %q\n", games[:5])
}

func observeMemUsage() {
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var arr [size]int

	// 2. print the memory usage (use the report func).
	report("First array memory usage")

	// 3. copy the array to a new array.
	newArr := arr

	// 4. print the memory usage
	report("Second array memory usage")

	// 5. pass the array to the passArray function
	passArray(newArr)
	report("After the call array memory usage")
	// 6. convert one of the arrays to a slice
	slice1 := arr[:]

	// 7. slice only the first 1000 elements of the array
	slice2 := arr[:1e3]

	// 8. slice only the elements of the array between 1000 and 10000
	slice3 := arr[1e3:1e4]

	// 9. print the memory usage (report func)
	fmt.Println(newArr[0], slice2[0], slice1[0], slice3[0])
	report("After slicing the arrays memory usage")
	// 10. pass the one of the slices to the passSlice function
	passSlice(slice3)
	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("slice1 size: %d \n", unsafe.Sizeof(slice1)/8)
	fmt.Printf("slice2 size: %d \n", unsafe.Sizeof(slice2)/8)
	fmt.Printf("slice3 size: %d \n", unsafe.Sizeof(slice3)/8)
	fmt.Printf("Arr size: %d \n", unsafe.Sizeof(arr)/8)
	fmt.Printf("NewArr size: %d \n", unsafe.Sizeof(newArr)/8)
}

// passes [size]int array about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}

func sortMiddle() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	m := len(items) / 2
	middle := items[m-1 : m+2]
	sort.Strings(middle)
	fmt.Println("Sorted  :", items)
}

func internalsBackingArr() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// ONLY ADD YOUR CODE HERE
	//
	// Ensure that nums slice never changes even though
	// the mine slice changes.
	mine := append([]int(nil), nums...)
	// ----------------------------------------

	// DON'T TOUCH THE FOLLOWING CODE
	//
	// This code changes the elements of the nums
	// slice.
	//
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}

func slicingHousingPrices() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		to, from, idx int
		rows          []string
	)

	rows = strings.Split(data, "\n")
	header := strings.Split(rows[0], separator)

	positions := os.Args[1:]
	if len(positions) >= 1 {
		idx = indexOf(header, positions[0])
		if idx != -1 {
			from = idx
		}
	}

	if len(positions) >= 2 {
		idx2 := indexOf(header, positions[1])
		if idx2 != -1 {
			if idx > idx2 {
				from = 0
			}

			to = idx2 + 1
		} else {
			to = len(header)
		}
	}

	if to == 0 {
		to = len(header)
	}

	for _, v := range header[from:to] {
		fmt.Printf("%-15s", v)
	}

	fmt.Println()
	for _, row := range rows[1:] {
		fmt.Println()
		col := strings.Split(row, separator)

		for _, v := range col[from:to] {
			fmt.Printf("%-15s", v)
		}
	}
}

func indexOf(header []string, s string) int {
	return slices.IndexFunc(header, func(item string) bool {
		return strings.ToLower(item) == strings.ToLower(s)
	})
}

func slicingByArgs() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("\t%q\n\n", ships)
	positions := os.Args[1:]
	if len(positions) < 1 || len(positions) > 2 {
		fmt.Println("\tProvide only the [starting] and [stopping] positions")
		return
	}

	start, stop := 0, len(ships)

	if len(positions) >= 1 {
		if v, err := strconv.Atoi(positions[0]); err == nil && v > -1 && v <= len(ships) {
			start = v
		} else {
			fmt.Println("\tWrong positions")
			return
		}
	}

	if len(positions) == 2 {
		if v, err := strconv.Atoi(positions[1]); err == nil && v > -1 && start <= v {
			stop = v
		} else {
			fmt.Println("\tWrong positions")
			return
		}
	}

	fmt.Printf("\t%v\n", ships[start:stop])
}

func slicingBasics() {
	data := "2 4 6 1 3 5"
	stringNumbers := strings.Split(data, " ")
	var numbers []int
	for i := range stringNumbers {
		if number, err := strconv.Atoi(stringNumbers[i]); err == nil {
			numbers = append(numbers, number)
		}
	}
	l := len(numbers)
	fmt.Printf("numbers    : %v\n", numbers)
	evans := numbers[:3]
	fmt.Printf("evans      : %v\n", evans)
	odds := numbers[3:]
	fmt.Printf("odds       : %v\n", odds)
	midd := numbers[l/2 : l/2+2]
	fmt.Printf("midd       : %v\n", midd)
	two := numbers[:2]
	fmt.Printf("two        : %v\n", two)
	lastTwo := numbers[l-2:]
	fmt.Printf("lastTwo    : %v\n", lastTwo)
	eventLast := numbers[2:3]
	fmt.Printf("last even  : %v\n", eventLast)
	oddLastTwo := numbers[4:]
	fmt.Printf("odd last 2 : %v\n", oddLastTwo)
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
