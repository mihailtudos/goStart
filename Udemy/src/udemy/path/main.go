package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"udemy.com/path/utils"
	"unicode"
	"unicode/utf8"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/utf8string"
)

type Creds struct {
	username string
	password string
}

func main() {
	convertFeetToMeters()
}

func convertFeetToMeters() {
	if len(os.Args) < 2 {
		fmt.Println(utils.Usage)
		return
	}

	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}

	meters := feet * 0.3048
	fmt.Printf("%g feet is %0.2g meters.\n", feet, meters)
}

func challengeTwo() {
	l := len(os.Args)
	if l != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	c := []Creds{{username: "jack", password: "1888"}, {username: "inanc", password: "1879"}}
	un, pw := os.Args[1], os.Args[2]
	var cUser Creds

	for _, user := range c {
		if user.username == un {
			cUser = user
			break
		}
	}

	if cUser == (Creds{}) {
		fmt.Printf("Access denied for %q\n", un)
	} else if cUser.username == un && cUser.password == pw {
		fmt.Printf("Access granted to %q\n", un)
	} else {
		fmt.Printf("Invalid pw for %q\n", un)
	}
}

func challengeOne() {
	l := len(os.Args)
	if l != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	c := Creds{username: "jack", password: "1888"}

	un, pw := os.Args[1], os.Args[2]

	if c.username == un && c.password == pw {
		fmt.Printf("Access granted to %q\n", un)
	} else if c.username == un {
		fmt.Printf("Invalid pw for %q\n", un)
	} else {
		fmt.Printf("Access denied for %q\n", un)
	}
}

func vowelOrCons() {
	l := len(os.Args)
	var input = ""
	if l == 2 {
		input = os.Args[1]
	} else {
		fmt.Println("Give me a letter")
		return
	}

	if len(input) > 1 {
		fmt.Println("Give me a letter")
		return
	}

	first := unicode.ToLower(utf8string.NewString(input).At(0))

	if isVowel(first) {
		fmt.Printf("%q it's a vowel.\n", first)
		return
	}

	if isSemiVowel(first) {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", first)
		return
	}

	fmt.Printf("%q is a consonant.\n", first)
}

func isVowel(f rune) bool {
	vowels := []rune{'a', 'e', 'o', 'i'}

	idx := slices.IndexFunc(vowels, func(r rune) bool {
		return r == f
	})

	return idx != -1
}

func isSemiVowel(f rune) bool {
	semi := []rune{'y', 'w'}

	idx := slices.IndexFunc(semi, func(r rune) bool {
		return r == f
	})

	return idx != -1
}

func argCount() {
	s := len(os.Args)

	if s == 1 {
		fmt.Println("Give me args")
	} else if s == 2 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if s == 3 {
		str := os.Args[1] + " " + os.Args[2]

		fmt.Printf("There are two: %q\n", str)
	} else {
		fmt.Printf("There are %d arguments.\n", s)
	}
}

func simplifyCondition() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

func printAge(age int) {
	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
}

func printFullName() {
	var (
		fn, ln string
	)
	fmt.Println("Enter your first name:")
	fmt.Scanln(&fn)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&ln)

	fmt.Printf("%s %s\n", fn, ln)
}

func printTypeBool(v bool) {
	fmt.Printf("Type of %v is %[1]T.\n", v)
}

func printTypeString(v string) {
	fmt.Printf("Type of %v is %[1]T.\n", v)
}

func printTypeFloat(v float64) {
	fmt.Printf("Type of %v is %[1]T.\n", v)
}

func printTypeInt(v int) {
	fmt.Printf("Type of %v is %[1]T.\n", v)
}

func printQuote() {
	msg := "Hello World"
	fmt.Printf("%q\n", msg)
}

func printTheTemp() {
	t := 24.2612
	fmt.Printf("Temperature is %.2f degrees.\n", t)
}

func falseClaims() {
	tf := false
	fmt.Printf("There are %t claims.", tf)
}

func printYourName() {
	fn := os.Args[1]
	ln := os.Args[2]
	fmt.Printf("%s %s\n", fn, ln)
}

func printYourAge() {
	fmt.Println("Enter your age:")
	var age string
	fmt.Scanln(&age)
	fmt.Printf("My age is %s\n", age)

}

func printFTest() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// swiss army knife %v verb
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

	// why use other verbs than? because: type-safety
	// uncomment to see the warnings:
	//
	// fmt.Printf("Planet: %d\n", planet)
	// fmt.Printf("Distance: %s millions kms\n", distance)
	// fmt.Printf("Orbital Period: %t days\n", orbital)
	// fmt.Printf("Does %v has life? %f\n", planet, hasLife)

	// correct verbs:
	// fmt.Printf("Planet: %s\n", planet)
	// fmt.Printf("Distance: %d millions kms\n", distance)
	// fmt.Printf("Orbital Period: %f days\n", orbital)
	// fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// precision
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}

func iotaSeasons() {
	const (
		Winter = 12 - (iota * 3)
		Fall
		Summer
		Spring
	)
	fmt.Println(Winter, Spring, Summer, Fall)
}

func enumConst2() {
	const _ = iota
	const (
		_ = iota
		Jan
		Feb
		Mar
	)
	fmt.Println(Jan, Feb, Mar)
}

func enumConst() {
	const max = 11
	const (
		Nov = max - iota
		Oct
		Sep
	)
	fmt.Println(Nov, Oct, Sep)
}

func trimRight() {
	name := "inanç           "
	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}

func trimIt() {
	msg := `
	
	The weather looks good.
I should go and play.



	`

	fmt.Println(strings.TrimSpace(msg))
}

func toLower() {
	input := os.Args[1]
	fmt.Println(strings.ToLower(input))
	fmt.Println()
}

func improveTheBanger() {
	var msg string
	fmt.Println("Please enter a message:")
	fmt.Scanln(&msg)
	l := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", l)

	fmt.Println(s)
}

func countTheChars() {
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	fmt.Printf(`Your name is %d long`, utf8.RuneCountInString(name))
	fmt.Println()
}

func rawConcat() {
	var name string
	fmt.Scanln(&name)
	msg := `hi %s!
how are you?
`

	fmt.Printf(msg, name)
}

func printJson() {
	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"
	json = `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}
`
	fmt.Println(json)
}

func windowPath() {
	path := `
	c:\program files\duper super\fun.txt
	c:\program files\really\funny.png
	`
	fmt.Println(path)
}

func repeatChar() {
	w := os.Args[1]
	w1 := strings.Repeat("!", len(w)) + w + strings.Repeat("!", len(w))
	fmt.Println(w1)
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
