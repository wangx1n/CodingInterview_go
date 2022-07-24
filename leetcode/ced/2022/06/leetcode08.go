package main

import "math"

// NumberState 有限状态自动机定义
type NumberState struct {
	state string
	ans   int64
	sign  int
	table map[string][]string
}

// get 更新状态
func (a *NumberState) get(c byte) {
	a.state = a.table[a.state][a.getCol(c)]
	if a.state == "Append" {
		a.ans *= 10
		a.ans += int64(c - '0')
		if a.sign == -1 {
			a.ans = min(a.ans, int64(-1*math.MinInt32))
		} else {
			a.ans = min(a.ans, int64(math.MaxInt32))
		}
	} else if a.state == "Sign" {
		if c == '-' {
			a.sign = -1
		}
	}
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

// getCol 取出某一行的状态
func (a *NumberState) getCol(c byte) int {
	if c == ' ' {
		return 0
	} else if c >= '0' && c <= '9' {
		return 1
	} else if c == '-' || c == '+' {
		return 2
	} else {
		return 3
	}
}

func myAtoi(s string) int {
	a := &NumberState{
		ans:   0,
		state: "Start",
		sign:  1,
		table: map[string][]string{
			"Start":  []string{"Start", "Append", "Sign", "End"},
			"Append": []string{"End", "Append", "End", "End"},
			"Sign":   []string{"End", "Append", "End", "End"},
			"End":    []string{"End", "End", "End", "End"},
		},
	}
	length := len(s)
	for i := 0; i < length; i++ {
		a.get(s[i])
	}
	return int(int64(a.sign) * a.ans)
}
