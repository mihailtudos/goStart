package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

type digit [5]string

func main() {
	empty := false
	for {
		screen.Clear()
		screen.MoveTopLeft()
		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()
		clock := [8]digit{getDigit(h / 10), getDigit(h % 10), getSeparator(empty), getDigit(m / 10), getDigit(m % 10), getSeparator(empty), getDigit(s / 10), getDigit(s % 10)}

		for i, _ := range clock[0] {
			for _, c := range clock {
				fmt.Printf("%-5s", c[i])
			}
			fmt.Println()
		}
		empty = !empty
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
		" █ ",
		"   ",
		" █ ",
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
