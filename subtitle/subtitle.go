package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const subs_dir = "./" // "./"

func main() {
	files, _ := os.ReadDir(subs_dir)
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".srt") || strings.HasSuffix(filename, ".vtt") {
			Process(filename)
		}
	}
}

func Process(filename string) {
	path := filepath.Join(subs_dir, filename)
	bytes, _ := os.ReadFile(path)
	content := string(bytes)

	lines := strings.Split(content, "\n")
	words := ExtractWords(lines)
	words_map := MakeWordsMap(words)
	write_bytes, _ := json.MarshalIndent(words_map, "", "	")
	os.WriteFile(path+".json", write_bytes, 0666)
}

func ExtractWords(lines []string) []string {
	words := []string{}
	indexer := 2
	for _, line := range lines {
		if line == "\r" {
			indexer += 2
			continue
		}
		if indexer > 0 {
			indexer--
			continue
		}
		words = append(words, strings.Split(line, " ")...)
	}
	return words
}

func MakeWordsMap(words []string) map[string]int {
	words_map := make(map[string]int, 0)
	for _, word := range words {
		word = strings.NewReplacer(",", "", ".", "", "\r", "", "\"", "").Replace(word)
		// word = strings.ToLower(word)
		// make a function to remove common words
		if _, ok := words_map[word]; !ok {
			words_map[word] = 0
		}
		words_map[word]++
	}
	return words_map
}
