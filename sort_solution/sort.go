package sort_solution

func insertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		preIndex := i - 1
		temp := arr[i]

		for preIndex >= 0 && arr[preIndex] > temp {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}

		arr[preIndex+1] = temp
	}

	return arr
}

func shellSort(arr []int) []int {
	n := len(arr)
	half := n / 2

	for half > 0 {
		for i := half; i < n; i++ {
			preIndex := i - half
			temp := arr[i]
			for preIndex >= 0 && arr[preIndex] > temp {
				arr[preIndex+half] = arr[preIndex]
				preIndex -= half
			}
			arr[preIndex+half] = temp
		}
		half /= 2
	}
	return arr
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2

	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(nums1, nums2 []int) []int {
	m, n, index1, index2 := len(nums1), len(nums2), 0, 0
	ret := make([]int, 0, m+n)

	for index1 < m || index2 < n {
		if index1 == m {
			ret = append(ret, nums2[index2])
			index2++
		} else if index2 == n {
			ret = append(ret, nums1[index1])
			index1++
		} else if nums1[index1] < nums2[index2] {
			ret = append(ret, nums1[index1])
			index1++
		} else {
			ret = append(ret, nums2[index2])
			index2++
		}
	}
	return ret
}

func quickSort(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(nums []int, begin, end int) {
	if begin < end {
		index := findIndex(nums, begin, end)
		quick(nums, begin, index-1)
		quick(nums, begin+1, end)
	}
}

func findIndex(nums []int, begin, end int) int {
	pivot := nums[begin]

	for begin < end {
		for begin < end && nums[end] >= pivot {
			end--
		}
		nums[begin] = nums[end]
		for begin < end && nums[begin] <= pivot {
			begin++
		}
		nums[end] = nums[begin]
	}
	nums[begin] = pivot
	return begin
}

func heapSort(nums []int) []int {
	n := len(nums)

	for i := n/2 - 1; i >= 0; i-- {
		ajustHeap(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		swap(nums, 0, i)
		ajustHeap(nums, 0, i)
	}
	return nums
}

func ajustHeap(nums []int, rootIndex int, n int) {
	rootVal := nums[rootIndex]
	for i := rootIndex*2 + 1; i < n; i = i*2 + 1 {
		if i+1 < n && nums[i] < nums[i+1] {
			i += 1
		}
		if nums[i] <= rootVal {
			break
		}
		nums[rootIndex] = nums[i]
		rootIndex = i
	}
	nums[rootIndex] = rootVal
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
