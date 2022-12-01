package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	Check(err)
	return strings.TrimSuffix(string(content), "\n")
}

func ReadInput(day int) string {
	filename := fmt.Sprintf("day-%02d/input.txt", day)
	return readFile(filename)
}

func ReadTestInput(day int) string {
	filename := fmt.Sprintf("day-%02d/test_input.txt", day)
	return readFile(filename)
}
