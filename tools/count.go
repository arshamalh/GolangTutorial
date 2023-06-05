package tools

func Count(str string) map[string]int {
	characters := map[string]int{}
	for _, char := range str {
		if _, ok := characters[string(char)]; !ok {
			characters[string(char)] = 0
		}
		characters[string(char)]++
	}
	return characters
}
