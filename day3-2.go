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
	spiral := make([][]int, a)
	for i := range s {
		spiral[i] = make([]int, a)
	}
	spiral[mid][mid] = 1
	l := (a - 1) * (a - 1)
	for i := 0; i < l; i++ {
		for j := 0; j < 2*i+2; j++ {
			y := mid + i - j
			x := mid + i + 1
			idxVal := sumAdj(spiral, y, x)
			if idxVal > input {
				fmt.Println(idxVal)
				return
			}
			spiral[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid - i - 1
			x := mid + i - j
			idxVal := sumAdj(spiral, y, x)
			if idxVal > input {
				fmt.Println(idxVal)
				return
			}
			spiral[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid - i + j
			x := mid - i - 1
			idxVal := sumAdj(spiral, y, x)
			if idxVal > input {
				fmt.Println(idxVal)
				return
			}
			spiral[y][x] = idxVal
		}
		for j := 0; j < 2*i+2; j++ {
			y := mid + i + 1
			x := mid - i + j
			idxVal := sumAdj(spiral, y, x)
			if idxVal > input {
				fmt.Println(idxVal)
				return
			}
			spiral[y][x] = idxVal
		}
	}
}

func sumAdj(spiral [][]int, r, c int) int {
	ret := spiral[r][c+1] +
		spiral[r-1][c+1] +
		spiral[r-1][c] +
		spiral[r-1][c-1] +
		spiral[r][c-1] +
		spiral[r+1][c-1] +
		spiral[r+1][c] +
		spiral[r+1][c+1]
	return ret
}
