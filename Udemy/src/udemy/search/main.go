package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Something went wrong when loading the env file. STOPPING...")
	}
	words := filepath.SplitList(os.Getenv("PATH"))
	query := os.Args[1:]

	for _, q := range query {
		q = strings.ToLower(q)
		for i, w := range words {
			w = strings.ToLower(w)

			if strings.ContainsAny(w, q) {
				fmt.Printf("#%-2d: %v\n", i+1, w)
			}
		}
	}

}
