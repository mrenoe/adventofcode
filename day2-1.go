package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		numbers := strings.Fields(text)
		var largest, smallest int
		for i := range numbers {
			val, _ := strconv.Atoi(numbers[i])
			if i == 0 {
				smallest, largest = val, val
				continue
			} else if val > largest {
				largest = val
			} else if val < smallest {
				smallest = val
			}
		}
		fmt.Println("largest: ", largest, " smallest: ", smallest)
		sum += (largest - smallest)
		fmt.Println("Sum:", sum)
	}
}
