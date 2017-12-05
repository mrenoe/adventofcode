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
	fmt.Print("Enter text:")
	var instructions []int
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if strings.Compare(text, "") == 0{
			break
		}
		num, _ := strconv.Atoi(text)
		instructions = append(instructions, num)
	}
	//fmt.Println(instructions)
	steps:=0
	i:=0
	for i<len(instructions){
		curStep := instructions[i]
		//fmt.Println("CurInstruction is: ", curStep)
		if curStep >= 3{
			instructions[i]--
		}else{
			instructions[i]++	
		}
		//fmt.Println("Now incremented to  ", instructions[i])
		i = i+curStep
		//fmt.Println("Next instruction will be at:  ", i, " in: ", instructions)
		steps++
	}
	fmt.Println("steps to break!: ", steps)
}
