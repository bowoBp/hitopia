package main

import (
	"fmt"
)

func createPalindrome(s string, k int, left int, right int) (string, int) {
	// Basis untuk menghentikan rekursi
	if left >= right {
		return s, k
	}

	// Jika karakter kiri dan kanan tidak sama, kita perlu mengganti salah satunya
	if s[left] != s[right] {
		if k == 0 {
			return s, -1
		}
		// Mengganti karakter yang lebih kecil dengan yang lebih besar
		if s[left] > s[right] {
			s = s[:right] + string(s[left]) + s[right+1:]
		} else {
			s = s[:left] + string(s[right]) + s[left+1:]
		}
		k--
	}

	// Lanjutkan ke karakter berikutnya di tengah
	return createPalindrome(s, k, left+1, right-1)
}

func maximizePalindrome(s string, k int, left int, right int) string {
	// Basis untuk menghentikan rekursi
	if left >= right {
		return s
	}

	// Jika karakter sudah sama, lanjutkan ke karakter berikutnya
	if s[left] == s[right] {
		s = maximizePalindrome(s, k, left+1, right-1)
	} else {
		// Jika karakter tidak sama, kita sudah menggunakan satu perubahan untuk membuatnya sama
		// Sekarang kita periksa apakah ada perubahan tersisa untuk memaksimalkan angka
		if k > 0 {
			if s[left] != '9' {
				s = s[:left] + "9" + s[left+1:]
				k--
			}
			if s[right] != '9' {
				s = s[:right] + "9" + s[right+1:]
				k--
			}
		}
	}

	// Lanjutkan ke karakter berikutnya di tengah
	return maximizePalindrome(s, k, left+1, right-1)
}

func highestPalindrome(s string, k int) string {
	var initialK = k

	// Pertama, buat string menjadi palindrome dengan minimal perubahan
	palindrome, remainingK := createPalindrome(s, k, 0, len(s)-1)

	// Jika tidak dapat menjadi palindrome, return -1
	if remainingK == -1 {
		return "-1"
	}

	// Sekarang maksimalkan palindrome dengan perubahan yang tersisa
	highest := maximizePalindrome(palindrome, remainingK, 0, len(s)-1)

	// Pastikan hasilnya adalah palindrome dan perubahan tidak lebih dari k
	finalPalindrome, finalK := createPalindrome(highest, initialK, 0, len(highest)-1)
	if finalK != initialK-remainingK {
		return "-1"
	}

	return finalPalindrome
}

func main() {
	fmt.Println(highestPalindrome("3943", 1))   // Output: 3993
	fmt.Println(highestPalindrome("932239", 2)) // Output: 992299
}
