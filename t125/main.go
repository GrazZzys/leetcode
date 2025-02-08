package main

func main() {
	isPalindrome(" ")
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		val1 := s[i]
		val2 := s[j]

		if (val1 < 'a' || val1 > 'z') && (val1 < 'A' || val1 > 'Z') && (val1 < '0' || val1 > '9') {
			i++
			continue
		}

		if (val2 < 'a' || val2 > 'z') && (val2 < 'A' || val2 > 'Z') && (val2 < '0' || val2 > '9') {
			j--
			continue
		}

		if val1 >= 'A' && val1 <= 'Z' {
			val1 = uint8(rune(val1) - ('A' - 'a'))
		}
		if val2 >= 'A' && val2 <= 'Z' {
			val2 = uint8(rune(val2) - ('A' - 'a'))
		}

		if val1 != val2 {
			return false
		}

		i, j = i+1, j-1
	}

	return true
}
