package main

import "fmt"
type ImageData struct {
	src    string
	tp     string
	title  string
	width  int
	height int
}

var ImageDatas []ImageData


func main() {

	test := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	n := len(test) // rowsnum
	var a11 []int
	for i := 0; i < n; i++ {
		a11 = append(a11, test[i][0:1][0])
	}
	fmt.Println(a11)

	imageData := ImageData{"src", "tp", "title", 1, 2}

	ImageDatas = append(ImageDatas, imageData)
}
