package main

import (
	"fmt"
	"path"
	"strconv"
)

func main() {
	discardFile()
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
