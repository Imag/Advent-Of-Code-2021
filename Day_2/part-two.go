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

	x, y, aim := 0, 0, 0
	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ") 
		option := input[0]
		value, _ := strconv.Atoi(input[1])
		
		switch option {
			case "forward": {
				x += value
				y += aim * value
			}
			case "up": {
				aim -= value
			}
			case "down": {
				aim += value
			}
		}
	}

	println(x * y)
}