package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)
	depths := []int{}
	sum := 0

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, i)
	}

	for i, v := range depths {
		if i != 0 && v > depths[i-1] {
			sum++
		}
	}

	println(sum)
}