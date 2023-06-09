package tools

import "testing"

func TestFizzBuzz(t *testing.T) {
	test_cases := []struct {
		// name string
		given int
		want  string
	}{
		{
			given: 5,
			want:  "buzz",
		},
		{
			given: 7,
			want:  "7",
		},
		{
			given: 15,
			want:  "fizz-buzz",
		},
	}
	for _, test := range test_cases {
		got := FizzBuzz(test.given)
		if test.want != got {
			t.Errorf("got (%s) is not equal to want (%s)", got, test.want)
		}

	}
}
