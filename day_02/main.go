package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	body, err := os.ReadFile("input/2.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	input := bytes.Split(body, []byte("\n"))
	// code starts here
	println(Part1(input))
	println(Part2(input))
}
