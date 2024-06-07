package main

import (
	"fmt"
)

// Fungsi untuk membuat string menjadi palindrome dengan minimal perubahan
func makePalindrome(s []rune, k int) ([]rune, int) {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			if k == 0 {
				return s, -1
			}
			if s[left] > s[right] {
				s[right] = s[left]
			} else {
				s[left] = s[right]
			}
			k--
		}
		left++
		right--
	}

	return s, k
}

// Fungsi untuk memaksimalkan nilai palindrome dengan sisa perubahan
func maximizePalindrome(s []rune, k int) []rune {
	left := 0
	right := len(s) - 1

	for left <= right {
		if left == right { // Tengah-tengah string
			if k > 0 {
				s[left] = '9'
			}
		} else {
			if s[left] < '9' {
				if k >= 2 && s[left] == s[right] { // Tidak perlu perubahan sebelumnya
					s[left], s[right] = '9', '9'
					k -= 2
				} else if k >= 1 && s[left] != s[right] { // Perubahan sebelumnya sudah terjadi
					s[left], s[right] = '9', '9'
					k--
				}
			}
		}
		left++
		right--
	}

	return s
}

// Fungsi utama untuk menemukan highest palindrome
func highestPalindrome(s string, k int) string {
	runes := []rune(s)

	// Buat string menjadi palindrome dengan minimal perubahan
	palindrome, remainingK := makePalindrome(runes, k)
	if remainingK == -1 {
		return "-1"
	}

	// Maksimalkan nilai palindrome dengan sisa perubahan
	highest := maximizePalindrome(palindrome, remainingK)

	return string(highest)
}

func main() {
	fmt.Println(highestPalindrome("3943", 1))   // Output: 3993
	fmt.Println(highestPalindrome("932239", 2)) // Output: 992299
}
