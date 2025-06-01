package main

func main() {
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1R := m - 1
	nums2R := n - 1
	nums1I := m + n - 1

	for nums2R >= 0 {
		if nums1R >= 0 && nums1[nums1R] > nums2[nums2R] {
			nums1[nums1I] = nums1[nums1R]
			nums1R--
		} else {
			nums1[nums1I] = nums2[nums2R]
			nums2R--
		}
		nums1I--
	}
}
