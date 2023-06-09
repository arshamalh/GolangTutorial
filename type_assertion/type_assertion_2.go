package type_assertion

import (
	"fmt"
	"tutorials/cmap"
)

// A type assertion provides access to an interface value's underlying concrete value.

func AssertGrades() {
	math_grades := make(cmap.CMap[float64])
	math_grades.Set("Arsham", 20.0)
	math_grades.Set("Abtin", 18.0)
	math_grades.Set("Kaveh", 16.0)

	grades := math_grades.Values()
	// new_grades := Convertor(grades)
	average := Average(grades)
	fmt.Println(average)
}

func Convertor(grades []any) []float64 {
	new_grades := make([]float64, 0)
	for _, grade := range grades {
		new_grade, ok := grade.(float64)
		if ok {
			new_grades = append(new_grades, new_grade)
		}
	}
	return new_grades
}

func Average(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}
