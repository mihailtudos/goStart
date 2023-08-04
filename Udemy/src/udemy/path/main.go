package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "Unknown"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println("Hello " + name)
}
