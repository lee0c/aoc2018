package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	freq := 0
	visited := make(map[int]bool)
	done := false

	for true {
		f, _ := os.Open("day1.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
	
		for scanner.Scan() {
			visited[freq] = true
			num, _ := strconv.Atoi(scanner.Text())
			freq += num

			if visited[freq] {
				fmt.Printf("Repeated: %d\n", freq)
				done = true
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(os.Stderr, "reading from file: ", err)
		}	

		if done {
			break
		}
	}
}