package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		exists := false
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if strings.Compare(text, "end") == 0 {
			break
		} else if len(text) == 0 {
			continue
		}
		encountered := map[string]int{}
		password := strings.Fields(text)
		for i := range password {
			if encountered[password[i]] > 0 { // base anagram
				exists = true
				break
			}
			encountered[password[i]] = 1
		}
		if exists == true {
			continue
		} else {
			sum++
		}
		fmt.Println("Sum:", sum)
	}
	fmt.Println("Sum:", sum)
}
