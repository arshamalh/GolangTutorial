package generics

import "fmt"

func Print() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(2.5, 3.0))
	fmt.Println(Add("arsham", " is coming"))
}

type Summable interface {
	int | float64 | string | uint
}

func Add[T Summable](num_1, num_2 T) T {
	return num_1 + num_2
}
