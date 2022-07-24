package _7

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	sumLen := n + m
	if sumLen%2 == 1 {
		return float64(getMedian(nums1, nums2, sumLen/2+1))
	} else {
		return float64(getMedian(nums1, nums2, sumLen/2)+getMedian(nums1, nums2, sumLen/2+1)) / 2.0
	}
}

func getMedian(nums1 []int, nums2 []int, k int) int {
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
		h := k / 2
		newIndex1 := min(index1+h, len(nums1)) - 1
		newIndex2 := min(index2+h, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
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
