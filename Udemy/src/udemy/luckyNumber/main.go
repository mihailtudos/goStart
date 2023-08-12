package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		firstTimeWinner = `Woooow you got it from the first try you are so lucky 🍀.
🥇 FIRST TIME WINNER!!!`
		usage = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
		tries        = 3
		lostMessages = []string{"☠️  YOU LOST... Try again?", "LOOSER ☠️☠️☠️", "😱 YOU LOST. TRY AGAIN? 😱😱😱"}
		winMessages  = []string{"🎉 YOU WON", "🎉 YOU'RE AWESOME", "🎉 AMAZING, YOU WON 🎉"}
	)

	reader := bufio.NewReader(os.Stdin)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(11)
	var guesses int
	fmt.Printf(usage, tries)
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Printf("You are cheeting :D the lucky number is %d\n\n", n)
	}

	for guesses < tries {
		fmt.Println("Enter a number between 0 and 10")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("please enter a single number between 0 and 10")
			continue
		}

		userNumber, convErr := strconv.Atoi(strings.TrimSpace(input))
		if convErr != nil || userNumber < 0 || userNumber > 10 {
			fmt.Println("Entered value must be a number between 0 and 10")
			continue
		}

		if userNumber == n {
			fmt.Println(winMessages[r.Intn(len(winMessages))])
			if guesses == 0 {
				fmt.Println(firstTimeWinner)
			}
			fmt.Printf("Congrats you won using %d chances out of available 3", guesses+1)
			break
		}

		guesses++
		if guesses < 3 {
			fmt.Printf("Ohh no, you were so close, try again you have %d chances left.\n", 3-guesses)
		} else {
			fmt.Println(lostMessages[r.Intn(len(lostMessages))])
		}
	}
}
