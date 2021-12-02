package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	horizontal, depth := 0, 0
	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ") 
		option := input[0]
		value, _ := strconv.Atoi(input[1])
		
		switch option {
			case "forward": {
				horizontal += value
			}
			case "up": {
				depth -= value
			}
			case "down": {
				depth += value
			}
		}
	}

	println(horizontal * depth)
}