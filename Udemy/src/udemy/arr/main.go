package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	moods2()
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
	} else {
		i = 0
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
