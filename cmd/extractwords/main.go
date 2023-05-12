package main

import (
	"bufio"
	"fmt"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/file"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	curDir, _   = os.Getwd()
	specailStrs = []string{",", "[", "]", "/", "–", "_", "\"", "<", ">", "=", "(", ")", ":", "\n", "\r", ";", ".", "*", "`", "{", "}", "$", "'", "#", "%", "?", "\\", "-", "@", "+", "|", "!", "~"}
)

func main() {
	dir := "C:\\Users\\ooooo\\Development\\Code\\Demo\\istio.io\\content\\en"
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	targetFile := filepath.Join(curDir, "output.txt")
	words := TransformWords(dir)
	err := os.WriteFile(targetFile, []byte(words), 0644)
	if err != nil {
		panic(fmt.Sprintf("WriteFile to %s, error: %v", targetFile, err))
	}
}

func TransformWords(dir string) string {
	var words []string
	file.TravelDir(dir, func(path string) {
		fmt.Println(path)
		if !strings.Contains(path, ".md") {
			return
		}

		file, err := os.Open(path)
		if err != nil {
			return
		}

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			words = append(words, strings.Split(filterLine(line), " ")...)
		}
	})
	return filterWords(words)
}

func filterWords(words []string) string {
	var sb strings.Builder
	m := make(map[string]bool)
	for _, word := range words {
		if len(word) <= 3 {
			continue
		}
		if _, exist := m[word]; !exist {
			sb.WriteString(word)
			sb.WriteString(" ")
			m[word] = true
		}
	}
	return sb.String()
}

func filterLine(line string) string {
	for _, str := range specailStrs {
		line = strings.ReplaceAll(line, str, " ")
	}
	return line
}
