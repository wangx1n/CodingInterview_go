package main

import (
	"math"
)

type point struct {
	i int
	j int
}

type max []point
type pointList []point

func main() {
	matrix := [][]int{
	 {1,10,4,2},{9,3,8,7},{15,16,17,12},
	}
	luckyNumbers(matrix)
}
// luckyNumbers 标准答案
func luckyNumbers(matrix [][]int) (ans []int) {
	for _, row := range matrix {
	next:
		for j, x := range row {
			for _, y := range row {
				if y < x {
					continue next
				}
			}
			for _, r := range matrix {
				if r[j] > x {
					continue next
				}
			}
			ans = append(ans, x)
		}
	}
	return
}

// luckyNumbers_self 自己写的
func luckyNumbers_self (matrix [][]int) []int {
	var minArray []point
	var maxArray []point
	var result []int
	for i := 0; i < len(matrix); i++ {
		minArray = append(minArray, point{i, findMin(matrix[i])})
	}
	for j := 0; j < len(matrix[0]); j++ {
		colArray := constructColArray(matrix, j)
		maxArray = append(maxArray, point{findMax(colArray), j})
	}
	for _, resMin := range minArray {
		for _, resMax := range maxArray {
			if resMin == resMax {
				result = append(result, matrix[resMin.i][resMin.j])
			}
		}
	}
	return result
}

// TODO: 可以被用作common
func constructColArray(matrix [][]int, j int) (col []int) {
	for row := 0; row < len(matrix); row++ {
		col = append(col, matrix[row][j: j + 1][0])
	}
	return
}


//func in(target point, int_array []int, l pointList) bool {
//	sort.Sort(int_array)
//	index := sort.SearchInts(int_array, target)
//	if index < len(int_array) && int_array[index] == target {
//		return true
//	}
//	return false
//}

func findMin(arrayFromMatrix []int) (res int) {
	min := math.MaxInt32
	for i := 0; i < len(arrayFromMatrix); i++ {
		if arrayFromMatrix[i] < min {
			min = arrayFromMatrix[i]
			res = i
		}
	}
	return
}

func findMax(arrayFromMatrix []int) (res int) {
	max := math.MinInt32
	for i := 0; i < len(arrayFromMatrix); i++ {
		if arrayFromMatrix[i] > max {
			max = arrayFromMatrix[i]
			res = i
		}
	}
	return
}

func (I pointList) Len() int {
	return len(I)
}
func (I pointList) Less(i, j int) bool {
	return I[i].i < I[j].i
}
func (I pointList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
