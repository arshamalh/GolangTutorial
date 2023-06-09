package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	test_cases := []struct {
		// name string
		given string
		want  map[string]int
	}{
		{"aa", map[string]int{"a": 2}},
		{"abc bca", map[string]int{" ": 1, "a": 2, "b": 2, "c": 2}},
	}

	for _, test := range test_cases {
		got := Count(test.given)
		assert.Equalf(t, test.want, got, "not equal %v %v", got, test.want)
	}
}
