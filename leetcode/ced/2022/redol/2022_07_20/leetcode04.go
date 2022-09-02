package _022_07_20

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	sumLen := n1 + n2
	if sumLen%2 == 1 {
		return float64(findMid(nums1, nums2, sumLen/2+1))
	} else {
		return float64(findMid(nums1, nums2, sumLen/2)+findMid(nums1, nums2, sumLen/2+1)) / 2.0
	}
}

func findMid(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		h := k / 2
		newIdx1 := min(idx1+h, len(nums1)) - 1
		newidx2 := min(idx2+h, len(nums2)) - 1
		if nums1[newIdx1] <= nums2[newidx2] {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newidx2 - idx2 + 1
			idx2 = newidx2 + 1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
