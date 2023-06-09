package tools

import "fmt"

// 20000 > x => bad salary
// 20k < x < 100k => fair salary
// 100k < x => good salary

func Greeting(name string, age int, salary int) string {
	var result string
	if age < 12 {
		result = fmt.Sprintf("Hello %s, you are a child", name)
	} else if age > 12 && age < 20 {
		result = fmt.Sprintf("Hello %s, you are a teenager", name)
	} else {
		result = fmt.Sprintf("Hello %s, you are %d years old", name, age)
	}

	return result
}
