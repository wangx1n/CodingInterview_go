package graph

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		a []int
		b = make([][]int, 3)
		c = new([]int)
	)
	fmt.Println(a, b, c)
	fmt.Println(unsafe.Sizeof(c))
	a = append(a, 1,2,3,4,5,6,7,8,9,10)
	*c = append(*c, 1,2,3,4,5,6,7,8,9,10)
	fmt.Println(a, c)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(*c))
}
