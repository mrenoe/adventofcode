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
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	numbers := strings.SplitAfter(text, "")

	sum := 0
	step := len(numbers) / 2
	for i := range numbers {
		var first, second int
		first = i
		second = (i + step) % len(numbers)
		if numbers[first] == numbers[second] {
			num, _ := strconv.Atoi(numbers[first])
			sum += num
			fmt.Println("The same: first: ", numbers[first], " second: ", numbers[second], " sum: ", sum)
		}
	}
	fmt.Println("sum:", sum)
}
