package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	wordFinder()
}

func wordFinder() {
	if len(os.Args) < 2 {
		fmt.Println("Please give me a word to search.")
		return
	}

	f := []string{"and", "or", "was", "the", "since", "very"}
	words := strings.Fields("lazy cat jumps again and again and again since the beginning this was very important")
	search := os.Args[1:]

	for i, v := range words {
		words[i] = strings.ToLower(strings.TrimSpace(v))
	}

	for i, v := range f {
		f[i] = strings.ToLower(strings.TrimSpace(v))
	}

	for i, v := range search {
		search[i] = strings.ToLower(strings.TrimSpace(v))
	}

	for i, v := range search {
		if !slices.Contains(f, v) && slices.Contains(words[:], v) {
			fmt.Printf("#%d - %v\n", i, v)
		}
	}
}

func sortArr() {
	if len(os.Args) != 6 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	var numbers [5]int

	for i, val := range os.Args[1:] {
		if v, err := strconv.Atoi(val); err == nil {
			numbers[i] = v
			continue
		}
	}
	for range numbers {
		for i, v := range numbers {
			if i < len(numbers)-1 && v > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	fmt.Printf("%v\n", numbers)

}

func avgSum() {
	if len(os.Args) != 6 {
		println("Please tell me numbers (maximum 5 numbers).")
		return
	}
	var (
		numbers [5]float64
		sum     float64
	)

	for i, val := range os.Args[1:] {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			numbers[i] = v
			continue
		}
		numbers[i] = 0
	}

	for _, v := range numbers {
		sum += v
	}

	fmt.Printf("Your number: %v\n", numbers)
	fmt.Printf("Average: %.2f\n", sum/float64(len(numbers)))
}

func searchBook() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a book title")
		return
	}
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	lq := strings.ToLower(os.Args[1])
	var res bool
	println("Search Results:")
	for i, t := range books {
		if strings.Contains(strings.ToLower(t), lq) {
			res = true
			fmt.Printf("#%d - %s\n", i, t)
		}
	}

	if !res {
		fmt.Printf("The book doesn't exist\n")
	}
}

func currencyConvertor() {
	m := 1
	var err error
	if len(os.Args) >= 2 {
		if m, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Printf("First parameter must be a number %T given", os.Args[1])
			return
		}
	}
	const (
		GBP = iota
		RON
		EUR
	)
	rates := []float64{
		GBP: 1.2,
		RON: 4.5,
		EUR: 1.1}

	fmt.Printf("%v GBP is %.2f\n", m, rates[GBP]*float64(m))
	fmt.Printf("%v RON is %.2f\n", m, rates[RON]*float64(m))
	fmt.Printf("%v EUR is %.2f\n", m, rates[EUR]*float64(m))
}

func wizardPrint() {
	scientrists := [...][5]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"}}
	fmt.Println("==================================================")
	for _, scientist := range scientrists {
		for _, detail := range scientist {
			fmt.Printf("%-15v", detail)
		}
		fmt.Println()
	}
	fmt.Println("==================================================")
}

func moods2() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter your name. \n[name] [positive|negative]")
		return
	}

	moods := [...][3]string{
		{"bad", "upset", "sad"},
		{"happy", "awesome", "good"},
	}

	name := os.Args[1]
	mood := os.Args[2]
	if types := [...]string{"positive", "negative"}; !slices.Contains(types[:], mood) {
		fmt.Println("Please enter your name. \n[name] [positive|negative]")
		return
	}

	rand.NewSource(time.Now().UnixMicro())
	var i int
	if mood == "positive" {
		i = 1
	}
	var j = rand.Intn(len(moods[i]))
	fmt.Printf("%s is feeling %s\n", name, moods[i][j])
}

func arrayAssign() {
	books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}
	upper := books
	lower := books

	for i, v := range books {
		lower[i] = strings.ToLower(v)
		upper[i] = strings.ToUpper(v)
	}

	fmt.Printf("%q\n", books)
	fmt.Printf("%q\n", lower)
	fmt.Printf("%q\n", upper)
}

func compareArrays() {
	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	//Use     : uint8 for one of the arrays instead of byte here.
	//Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]uint8

	fmt.Println(data == image)
}

func fixCode() {
	names := [...]string{
		"Einstein",
		"Shepard",
		"Tesla"}

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden"}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

func getAndSet() {
	name := [...]string{"Constantin", "Petru"}
	distances := [...]int{30, 33, 55, 341, 9}
	data := [...]byte{'H', 'e', 'm', '2', '3'}
	currency := [...]float64{3.1428}
	states := [...]bool{true, true, false, true}
	zero := [...]byte{}

	for i := 0; i < len(name); i++ {
		fmt.Printf("%s ", name[i])
	}
	fmt.Println()
	for i := 0; i < len(distances); i++ {
		fmt.Printf("%d ", distances[i])
	}
	fmt.Println()
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()
	for i := 0; i < len(currency); i++ {
		fmt.Printf("%0.2f", currency[i])
	}
	fmt.Println()
	for i := 0; i < len(states); i++ {
		fmt.Printf("%t ", states[i])

	}

	fmt.Println()
	for i := 0; i < len(zero); i++ {
		fmt.Printf("%v ", zero[i])

	}

	fmt.Println("\n")

	for _, v := range name {
		fmt.Printf("%s ", v)
	}

	fmt.Println()
	for _, v := range distances {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
	for _, v := range data {
		fmt.Printf("%q ", v)
	}

	fmt.Println()
	for _, v := range currency {
		fmt.Printf("%.2f ", v)
	}

	fmt.Println()
	for _, v := range currency {
		fmt.Printf("%.2f ", v)
	}

	fmt.Println()
	for _, v := range states {
		fmt.Printf("%t ", v)
	}

	fmt.Println()
	for _, v := range zero {
		fmt.Printf("%q ", v)
	}

	fmt.Println()
}

func declareEmpty() {
	name := [3]string{}
	distances := [5]int{}
	data := [5]uint{}
	currency := [1]float64{}
	states := [4]bool{}
	zero := [0]uint8{}

	fmt.Printf(`
name: %#v
distances: %#v
data: %#v
currency: %#v
states: %#v
zero: %#v
`, name, distances, data, currency, states, zero)

	fmt.Printf(`
name: %T
distances: %T
data: %T
currency: %T
states: %T
zero: %T
`, name, distances, data, currency, states, zero)

	fmt.Printf(`
name: %q
distances: %v
data: %v
currency: %0.2f
states: %v
zero: %v
`, name, distances, data, currency, states, zero)

}

func moods() {
	moods := [...]string{"sad", "terrible", "happy", "awesome", "good", "fine"}
	if len(os.Args) != 2 {
		fmt.Println("Provide name")
		return
	}

	name := os.Args[1]
	rand.NewSource(time.Now().UnixMicro())

	fmt.Printf("%s feels %s\n", name, moods[rand.Intn(len(moods))])
}
