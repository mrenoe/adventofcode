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
		if strings.Compare(text, "end") == 0{
			break
		}
		numbers := strings.Fields(text)
		for i := range numbers {
			val1, _ := strconv.Atoi(numbers[i])
			for j := range numbers{
				val2, _ := strconv.Atoi(numbers[j])
				if i == j {
					continue
				} else if val1 % val2 == 0{
					var diff int
					if val1 > val2{
						diff = val1/val2
					}else {
						diff = val2/val1
					}	
					sum += diff
					break
				}
			}
		}
		fmt.Println("Sum:", sum)
	}
}