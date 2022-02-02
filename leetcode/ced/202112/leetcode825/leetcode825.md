## leetcode 825 双指针
这道题的关键在于双指针的left, right怎么变（废话）

首先看left怎么变：

以``agex``为中心遍历，如果左端点的值小于等于``agex * 0.5 + 7``，那么就说明这个左端点的值不符合y给x发消息的条件，left往右扩，left++

再看right怎么变：

以``agex``为中心遍历，如果右端点的值小于等于``agex``，符合y加x好友的条件，那么右端点就往右扩，right++

这道题复杂在以谁为中心看。

一开始我把中心定在了age[y]上，因为觉得age [y]在两个公式中间，那么就应该用agey[y]为中心。
```go
for _, agey := range ages {
		for 2 * agey - 14<= ages[left] {
			left++
		}
        // [1]
		for right + 1 < len(ages) && agey <= ages[right + 1] {
			right++
		}
		ans += right - left
	}
```

如果以age[y]为中心看，那么视角为“会给我发好友申请的人数”，而不是“我会给谁发好好友申请”，那么在标记[1]处就会造成，只要往右看age比我大的都会给我发好友申请，但是这只是判断了一半。也就是说，“他age比我大，但是我不一定会满足0.5*age[x] + 7 大于他”。这样左右指针的伸缩就没有意义了。

```go
	for _, agex := range ages {
		if agex < 15 {
			continue
		}
		for ages[left]*2 <= agex+14 {
			left++
		}
		for right + 1 < len(ages) && agex >= ages[right + 1] {
			right++
		}
		ans += right - left
	}
```

正确的解法如上所示。以age[x]为中心，选择满足条件的age[y]的位置。

**还需要再理解**