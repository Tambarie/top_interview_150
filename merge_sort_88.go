package main

import (
	"sort"
)

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	mergeSortApproachTwo([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}
	sort.Ints(nums1)
}

func mergeSortApproachTwo(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n - 1; p >= 0; p-- {
		if n == 0 {
			break
		}
		if m == 0 {
			nums1[p] = nums2[n-1]
			n--
			continue
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[p] = nums1[m-1]
			m--
		} else {
			nums1[p] = nums2[n-1]
			n--
		}
	}

}

func mergeThree(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 int = m - 1, n - 1, m + n - 1
	for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr3+1], nums2[:ptr2+1])
	}
}
