package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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
	infiniteKill()
}

func infiniteKill() {
	chars := []string{`\`, `/`, `|`, `-`}
	for {
		idx := rand.Intn(len(chars))
		fmt.Printf("%d", idx)
		c := chars[idx]
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(250 * time.Millisecond)
	}
}

func sumOfnEvenNumbers() {
	reader := bufio.NewReader(os.Stdin)
	var n1, n2, sum int
	var msg string
	var errConv error
	fmt.Println("Please enter 1st number:")
	input, err := reader.ReadString('\n')
	n1, errConv = strconv.Atoi(strings.TrimSpace(input))
	if err != nil || errConv != nil {
		fmt.Println("Please enter a number")
		return
	}

	fmt.Println("Please enter 2st number:")
	input, err = reader.ReadString('\n')
	if n2, errConv = strconv.Atoi(strings.TrimSpace(input)); err != nil || errConv != nil {
		fmt.Println("Please enter a number")
		return
	}

	if n1 > n2 {
		fmt.Println("1st number must be greater then 2nd")
		return
	}

	for {
		if n1%2 == 0 {
			sum += n1
			if n1+2 <= n2 {
				msg += strconv.Itoa(n1) + " + "
			} else {
				msg += strconv.Itoa(n1)
			}
		}
		if n2 < n1 {
			break
		}
		n1++
	}
	fmt.Println(msg+" =", sum)
}

func sumOfNumbersVerbose() {
	var (
		i   int
		res string
		sum int
	)

	for {
		i++
		if i < 10 {
			sum += i
			res += strconv.Itoa(i) + " + "
		}
		if i == 10 {
			sum += i
			res += strconv.Itoa(i) + " = " + strconv.Itoa(sum)
			break
		}
	}

	fmt.Println(res)
}

func sumOfNumbers() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid input")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please enter number")
		return
	}
	var sum int

	for {
		sum += n
		n--
		if n == 0 {
			break
		}
	}
	fmt.Printf("Total is %d\n", sum)
}

func operationsTable() {
	ops := []string{"*", "/", "-", "+", "%"}

	switch n := len(os.Args); {
	case n == 1:
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	case n == 2:
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	case n == 3:
		n, err := strconv.Atoi(os.Args[2])
		if !slices.Contains(ops, os.Args[1]) || err != nil || n < 0 {
			fmt.Println("Invalid operator.")
			return
		}
	default:
		fmt.Println("Invalid input")
		return
	}

	s, _ := strconv.Atoi(os.Args[2])

	fmt.Printf("%5s", os.Args[1])
	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= s; j++ {
			var res int
			switch os.Args[1] {
			case "*":
				res = i * j
			case "/":
				if j != 0 {
					res = i / j
				}
			case "-":
				res = i - j
			case "+":
				res = i + j
			case "%":
				if j != 0 {
					res = i * j
				}
			}
			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}

func multiplicationTable() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	s, err := strconv.Atoi(os.Args[1])
	if err != nil || s < 1 {
		fmt.Println("Wrong size")
		return
	}

	header := "X "
	for i := 0; i <= s; i++ {
		header += "\t" + strconv.Itoa(i)
	}
	fmt.Println(header)
	for i := 0; i <= s; i++ {
		row := strconv.Itoa(i) + "\t"
		for j := 0; j <= s; j++ {
			row += strconv.Itoa(j*i) + "\t"
		}
		fmt.Println(row)
	}
}

func convertToSwitch2() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch strings.ToLower(month) {
	case "april", "june", "september", "november":
		days = 30
	case "march", "january", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return

	}

	fmt.Printf("%q has %d days.\n", month, days)
}

func stringManipulator() {
	const def = ` [command] [string]

Available commands: lower, upper and title`
	if len(os.Args) != 3 {
		fmt.Println(def)
		return
	}
	cmd, str := os.Args[1], os.Args[2]

	switch cmd {
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "title":
		fmt.Println(strings.ToTitle(str))
	default:
		fmt.Printf("Unknown command: %q", cmd)

	}
}

func convertToSwitch() {
	const (
		usage       = "Usage: [username] [password]"
		errUser     = "Access denied for %q.\n"
		errPwd      = "Invalid password for %q.\n"
		accessOK    = "Access granted to %q.\n"
		user, user2 = "jack", "inanc"
		pass, pass2 = "1888", "1879"
	)

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fmt.Printf(accessOK, u)
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}

func richterScale() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a number.")
		return
	}

	if n, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		var desc string
		switch true {
		case n < 2.0:
			desc = "very minor"
		case n < 3.0:
			desc = "very minor"
		case n < 4.0:
			desc = "minor"
		case n < 5.0:
			desc = "light"
		default:
			desc = "massive"
		}

		r := "less than 2.0"
		if n > 2 && n < 10 {
			r = strconv.FormatFloat(math.Floor(n), 'f', 0, 64) + " - " + strconv.FormatFloat(math.Ceil(n)-0.1, 'f', 1, 64)
		} else if n >= 10 {
			r = "10+"
		}

		fmt.Printf("%v's richter scale is %v.\n", desc, r)
	} else {
		fmt.Println("Give me a number.")
	}
}

func getPeriodOfDay() {
	fmt.Printf("Current time is: %v:%v\n", time.Now().Local().Hour(), time.Now().Local().Minute())

	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("evening")
	case hour >= 12:
		fmt.Println("afternoon")
	case hour >= 6:
		fmt.Println("morning")
	default:
		fmt.Println("night")
	}
}

func getMonthDays() {
	year := time.Now().Year()
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	month := strings.ToLower(os.Args[1])

	monthsOf31Days := []string{"january", "march", "may", "july", "august", "october", "december"}
	monthsOf30Days := []string{"april", "june", "september", "november"}

	idx := slices.IndexFunc(monthsOf31Days, func(e string) bool {
		return month == e
	})

	if idx != -1 {
		fmt.Printf("%s has %d dats.\n", month, 31)
		return
	}

	idx = slices.IndexFunc(monthsOf30Days, func(e string) bool {
		return month == e
	})

	if idx != -1 {
		fmt.Printf("%s has %d dats.\n", month, 30)
		return
	}

	if month == "february" {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			fmt.Printf("%s has %d dats.\n", month, 29)
		} else {
			fmt.Printf("%s has %d dats.\n", month, 28)
		}
		return
	}

	fmt.Println("Give me a month name")
}

func lapYear() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	if n, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
		if n%4 == 0 && (n%100 != 0 || n%400 == 0) {
			fmt.Printf("%d is a lap year\n", n)
		} else {
			fmt.Printf("%d is not a lap year\n", n)
		}

		return
	}

	println("Give me a year number")
}

func oddEven() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	var msg string
	if n, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
		if n%2 == 0 {
			msg = "%d is even number"
		} else {
			msg = "%d is odd number"
		}

		if n%8 == 0 {
			msg += " and it's divisible by 8."
		}

		fmt.Printf(msg+"\n", n)
	} else {
		fmt.Printf("%d is not a number\n", n)
	}
}

func movieRating() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age.")
		return
	}

	if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
		if n > 17 {
			fmt.Println("R-Rated")
		} else if n > 12 && n < 18 {
			fmt.Println("PG-13")
		} else {
			fmt.Println("PG-Rated")
		}
		return
	} else {
		fmt.Printf("Wrong age: %d.\n", n)
	}

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
