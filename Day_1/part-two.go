package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	data,  _ := os.Open("input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	depths := [3]int{}
	sum := 0

	for i := 0; i < 3; i++ {
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())
		depths[i] = input
	}

	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())
		if currentDepth > depths[0] {
			sum++
		}

		depths[0] = depths[1]
		depths[1] = depths[2]
		depths[2] = currentDepth
	}

	println(sum)
}