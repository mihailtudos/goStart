package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

type digit [5]string

func main() {
	myClock()
}

func newClock() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder{
		"███",
		"█",
		"███",
		"  █",
		"███",
	}
	six := placeholder{
		"███",
		"█",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, minute, second := now.Hour(), now.Minute(), now.Second()
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
		}

		for line := range clock[0] {
			// Print a line for each placeholder in digits
			for idx, d := range clock {
				next := clock[idx][line]
				if d == colon && second%2 != 0 {
					next = "   "
				}
				fmt.Printf("%-5v", next)
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
		fmt.Println()
	}
}

func myClock() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()
		clock := [8]digit{getDigit(h / 10), getDigit(h % 10), getSeparator(s%2 != 0), getDigit(m / 10), getDigit(m % 10), getSeparator(s%2 != 0), getDigit(s / 10), getDigit(s % 10)}

		for i, _ := range clock[0] {
			for _, c := range clock {
				fmt.Printf("%-5s", c[i])
			}
			fmt.Println()
		}

		fmt.Println()
		time.Sleep(time.Duration(time.Millisecond * 1000))
	}
}

func getSeparator(empty bool) digit {
	if empty {
		return digit{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		}
	}
	return digit{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
}

func getDigit(d int) digit {
	zero := digit{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := digit{
		"██",
		" █",
		" █",
		" █",
		"███",
	}
	two := digit{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := digit{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := digit{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := digit{
		"███",
		"█",
		"███",
		"  █",
		"███",
	}
	six := digit{
		"███",
		"█",
		"███",
		"█ █",
		"███",
	}
	seven := digit{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := digit{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := digit{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [10]digit{zero, one, two, three, four, five, six, seven, eight, nine}
	return digits[d]
}
