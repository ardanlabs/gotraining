package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		fmt.Printf("word: %s -- position %d\n", scanner.Text(), count)
		count++
	}
}
