package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	greetFivePpl()
}

func greetFivePpl() {
	ppl := os.Args[1:]
	fmt.Printf("There are %d people!\n", len(ppl))
	for i := 0; i < len(ppl); i++ {
		fmt.Printf("Hello great %s!\n", ppl[i])
	}
	fmt.Println("Nice to meet you all.")
}

func greetMorePpl() {
	ppl := os.Args[1:]
	fmt.Printf("There are %d people!\n", len(ppl))
	for _, name := range ppl {
		fmt.Printf("Hello great %s!\n", name)
	}
	fmt.Println("Nice to meet you all.")
}

func printThePath() {
	fmt.Println(os.Args[0][2:])
}

func countTheArgs() {
	count := len(os.Args[1:])
	fmt.Printf("There are %d names.\n", count)
}

func convertAndFix5() {
	min := int8(127)
	max := int16(1000)
	fmt.Println((max) + int16(min))
}

func convertAndFix4() {
	age := 2
	fmt.Println(7.5 + float64(age))
}

func convertAndFix3() {
	fmt.Println((5.5))
}

func convertAndFix2() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

func convertAndFix() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

func discardFile() {
	dir, _ := path.Split("/src/udemy/path/main.go")
	fmt.Println(dir)
}

func swapper2() {
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)
}

func swapper() {
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)
}

func multiShort() {
	_, b := multi()
	fmt.Println(b)
}

func multi() (int, int) {
	return 5, 4
}

func multiAssign2() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on " + planet)
	fmt.Println("It's " + strconv.FormatBool(isTrue))
	fmt.Printf("It is %.2f degrees.", temp)
}

func multiAssign() {
	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}

func findTheRectangular() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = (width + height) * 2
	fmt.Printf(strconv.Itoa(perimeter))
}

func assignWithExpression() {
	n := 0.
	n = 3.14 * 2

	fmt.Println(n)
}

func makeItBlue() {
	color := "green"
	color = "blue"
	fmt.Println(color)
}

func varsToVar() {
	color := "green"
	color = "dark " + color

	fmt.Println(color)
}
