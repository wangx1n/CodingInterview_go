package _8

import (
	"fmt"
	"testing"
)

func TestLC27(t *testing.T) {
	a := []int{3, 2, 2, 3}
	removeElement(a, 3)
	fmt.Println(a)
}
