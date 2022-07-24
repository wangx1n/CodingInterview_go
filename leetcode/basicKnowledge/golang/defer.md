# defer

1. defer执行的时机

```go
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

结果：2
```

`原因`：由于return不是原子操作，实际上，return执行要经过三个过程 set Value -> defer -> return。

`注意`：这里返回值不是匿名返回值，如果是匿名返回值则无法做到这样的效果。

`问题`：1. return执行时除了这三个步骤还有哪些步骤。