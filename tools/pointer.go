package tools

import "fmt"

// Pass By Value
// Pass By Reference
//   1 - Make a Pointer to the value
//   2 - Pass that Pointer around

// & => Make Pointers
// * => Dereference and type

func pointer() {
	number := 2

	fmt.Println("First: ", number)
	makeTwice(&number)
	fmt.Println("Third: ", number)
}

func makeTwice(number *int) {
	*number *= 2
}
