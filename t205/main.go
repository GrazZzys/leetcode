package main

func main() {
	isIsomorphic("aa", "ab")
}

func isIsomorphic(s string, t string) bool {
	st := make(map[rune]rune)
	ts := make(map[rune]rune)

	for i, val := range t {
		if _, ok := ts[val]; !ok {
			ts[val] = rune(s[i])

			if _, ok := st[rune(s[i])]; ok {
				return false
			}

			st[rune(s[i])] = val
			continue
		}

		if rune(s[i]) != ts[val] {
			return false
		}
	}

	return true
}
