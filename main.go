package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to AOC 2022!")
	day := os.Args[1]
	fmt.Printf("Running day %s\n", day)
}
