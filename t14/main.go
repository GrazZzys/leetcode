package main

import "sort"

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	left := strs[0]
	right := strs[len(strs)-1]

	res := ""

	minLen := len(left)

	if len(right) < minLen {
		minLen = len(right)
	}

	for i := 0; i < minLen; i++ {
		if left[i] == right[i] {
			res += string(left[i])
		} else {
			break
		}
	}

	return res
}
