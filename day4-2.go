package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
			for j := range password {
				if j == i {
					continue
				}
				if isAnagram(password[i], password[j]) {
					exists = true
					fmt.Println("Anagram found: ", password[i], password[j])
					break
				}
			}
			if exists {
				break
			}
			encountered[password[i]] = 1
		}
		if exists {
			fmt.Println("Bad password Sum:", sum)
			continue
		} else {
			sum++
			fmt.Println("Good password Sum:", sum)
		}
	}
}

func isAnagram(old, newer string) bool {
	if len(old) != len(newer) {
		return false
	}
	old = SortString(old)
	newer = SortString(newer)
	if strings.Compare(old, newer) == 0 {
		return true
	}
	return false

}

func SortString(s string) string {
	newS := strings.Split(s, "")
	sort.Strings(newS)
	return strings.Join(newS, "")
}
