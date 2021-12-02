package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	horizontal, depth := 0, 0
	scanner := bufio.NewScanner(data)
	re := regexp.MustCompile("[0-9]+")
	
	for scanner.Scan() {
		action := strings.TrimSpace(re.Split(scanner.Text(), -1)[0])
		amount, _ := strconv.Atoi(re.FindString(scanner.Text()))
		
		switch action {
			case "forward": {
				horizontal += amount
			}
			case "up": {
				depth -= amount
			}
			case "down": {
				depth += amount
			}
		}
	}

	println(horizontal * depth)
}
