package main

import "fmt"

func main() {
	area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	fmt.Println(area)
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxHeight := min(height[left], height[right]) * (right - left)

	for left < (right - 1) {
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

		currHeight := min(height[left], height[right]) * (right - left)
		if maxHeight < currHeight {
			maxHeight = currHeight
		}
	}

	return maxHeight
}
