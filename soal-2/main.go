package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to check if brackets are balanced
func isBalanced(s string) string {
	stack := []rune{}
	matchingBracket := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matchingBracket[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(isBalanced(input))
}
