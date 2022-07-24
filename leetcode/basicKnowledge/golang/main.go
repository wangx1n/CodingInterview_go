package main

import "fmt"

func main() {
	fmt.Println(addOne())
}

func addOne() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}
