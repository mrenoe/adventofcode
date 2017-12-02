package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	numbers := strings.SplitAfter(text, "")

	sum := 0
	for i := range(numbers) {
		var first, second int
		second = i
		if i < 1 {
			first = len(numbers) - 1
		}else{
			first = i-1
		}
		if (numbers[first] == numbers[second]){
			num, _ := strconv.Atoi(numbers[first])
			sum += num
			fmt.Println("The same: first: ", numbers[first]," second: ", numbers[second], " sum: ", sum)
		}
	}
	fmt.Println("sum:", sum)

}