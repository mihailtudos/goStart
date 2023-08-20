package main

import (
	"time"
)

type digit [5]string

func main() {
	t := time.Now()
	h, m, s := t.Hour(), t.Minute(), t.Second()
	clock := [8]digit{getDigit(h / 10), getDigit(h % 10), getSeparator(), getDigit(m / 10), getDigit(m % 10), getSeparator(), getDigit(s / 10), getDigit(s % 10)}
	for range clock {

	}
}

func getSeparator() digit {
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
