package _022_07_30

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	sumLen := len1 + len2
	if sumLen%2 == 0 {
		return float64(findMid(nums1, nums2, sumLen/2)+findMid(nums1, nums2, sumLen/2+1)) / 2.0
	} else {
		return float64(findMid(nums1, nums2, sumLen/2+1))
	}
}

func findMid(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		p := k / 2
		newIndex1 := min(index1+p, len(nums1)) - 1
		newIndex2 := min(index2+p, len(nums2)) - 1
		if nums1[newIndex1] < nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
