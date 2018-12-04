package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("day1.txt")
	defer f.Close()

	freq := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		freq += num
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading from file: ", err)
	}

	fmt.Printf("%d\n", freq)
}