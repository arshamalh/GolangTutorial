package tools

import (
	"fmt"
	"strings"
	"tutorials/tools"
)

// Module is a project or service or workspace we are working on.

// Packages are a set of go codes with a certain goal,
// There are three types of packages:
//  - Std-lib
//  - Third-party
//  - Project structure

func main() {
	// Count
	word := "Hello world"
	count := tools.Count(word)
	fmt.Println(word, count)
	fmt.Println(strings.Repeat("-", 10))

	// FizzBuzz
	for i := 0; i < 20; i++ {
		fmt.Print(i, " ", tools.FizzBuzz(i), "  -  ")
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 10))

	// Greeting
	fmt.Println(tools.Greeting("Arsham", 23, 1e5))
	fmt.Println(tools.Greeting("Abtin", 32, 4e4))
	fmt.Println(tools.Greeting("Saam", 41, 7e4))
}
