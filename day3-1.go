package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Much of this based on work from /u/glassmountain on reddit
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter start val: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	input, _ := strconv.Atoi(text)
	start64 := float64(input)
	start64 = math.Ceil(math.Sqrt(start64))
	a := int(start64)
	mid := int(math.Floor(start64 / 2))

	s := make([][]int, a)
	for i := range s {
		s[i] = make([]int, a)
	}
	s[mid][mid] = 1
	max := a * a
	idxVal := 1
	for i := 0; i < max; i++ {
		for j := 0; j < 2*i+2; j++ {
			y := mid + i - j
			x := mid + i + 1
			idxVal++
			if idxVal >= input {
				fmt.Println(abs(x-mid) + abs(y-mid))
				return
			}
			s[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid - i - 1
			x := mid + i - j
			idxVal++
			if idxVal >= input {
				fmt.Println(abs(x-mid) + abs(y-mid))
				return
			}
			s[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid - i + j
			x := mid - i - 1
			idxVal++
			if idxVal >= input {
				fmt.Println(abs(x-mid) + abs(y-mid))
				return
			}
			s[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid + i + 1
			x := mid - i + j
			idxVal++
			if idxVal >= input {
				fmt.Println(abs(x-mid) + abs(y-mid))
				return
			}
			s[y][x] = idxVal
		}
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
