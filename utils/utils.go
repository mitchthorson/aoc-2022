package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Abs(num int) int {
	if num < 0 {
		return - num
	}
	return num
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	Check(err)
	return bufio.NewScanner(file)
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	Check(err)
	return strings.TrimSuffix(string(content), "\n")
}
func ReadInputLines(day int) []string {
	filename := fmt.Sprintf("day-%02d/input.txt", day)
	return ReadLines(filename)
}


func ReadInput(day int) string {
	filename := fmt.Sprintf("day-%02d/input.txt", day)
	return ReadFile(filename)
}

func ReadTestInput(day int) string {
	filename := fmt.Sprintf("day-%02d/test_input.txt", day)
	return ReadFile(filename)
}
