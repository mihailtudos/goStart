package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Something went wrong when loading the env file. STOPPING...")
	}
	words := strings.Fields(os.Getenv("CORPUS"))
	filters := []string{"the", "and", "or"}
	query := os.Args[1:]
query:
	for _, q := range query {
		q = strings.ToLower(q)
	filter:
		for i, w := range words {
			w = strings.ToLower(w)

			if slices.Contains(filters, w) {
				break filter
			}

			if w == q {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue query
			}
		}
	}

}
