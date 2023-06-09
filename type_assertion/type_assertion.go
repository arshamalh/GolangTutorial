package type_assertion

import (
	"fmt"
	"tutorials/tools"
)

// A type assertion provides access to an interface value's underlying concrete value.

func Assert() {
	var out_1, out_2 = tools.RandomOutputType()

	// Type Switch
	switch value := out_1.(type) {
	case int:
		result := Add(value, 2)
		fmt.Println(result)
	case string:
		fmt.Println("Your string is \"" + value + "\"")
	default:
		fmt.Printf("Out 1: %v - %T\n", value, value)
	}

	// Type Assertion
	number_1, ok_1 := out_1.(int)
	number_2, ok_2 := out_2.(int)
	if ok_1 && ok_2 {
		result := Add(number_1, number_2)
		fmt.Println(number_1, number_2, result)
	} else {
		fmt.Printf("Out 1: %v - %T\n", out_1, out_1)
		fmt.Printf("Out 2: %v - %T\n", out_2, out_2)
	}
}

func Add(number_1, number_2 int) int {
	return number_1 + number_2
}
