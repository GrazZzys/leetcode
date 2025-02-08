package main

func main() {
	canConstruct("aa", "ab")
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mag := make([]int, 26, 26)

	for _, val := range magazine {
		mag[val-97] = mag[val-97] + 1
	}

	for _, val := range ransomNote {
		if mag[val-97] <= 0 {
			return false
		}

		mag[val-97] = mag[val-97] - 1
	}

	return true
}
