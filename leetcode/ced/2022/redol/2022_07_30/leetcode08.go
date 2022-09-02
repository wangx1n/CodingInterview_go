package _022_07_30

import (
	"math"
)

type fms interface {
	getState(s byte) int
	getSum(s byte)
}

type fmsImpl struct {
	IfmServie fms
}

type fmsSerice struct {
	sign  int
	sum   int64
	table map[string][]string
	state string
}

func myAtoi(s string) int {
	f := fmsSerice{
		sign: 1,
		sum:  0,
		table: map[string][]string{
			"Start":  {"Start", "Sign", "Append", "End"},
			"Sign":   {"End", "End", "Append", "End"},
			"Append": {"End", "End", "Append", "End"},
			"End":    {"End", "End", "End", "End"},
		},
		state: "Start",
	}
	for i := 0; i < len(s); i++ {
		f.getSum(s[i])
	}
	return int(int64(f.sign) * f.sum)
}

func (fs *fmsSerice) getState(s byte) int {
	if s == ' ' {
		return 0
	} else if s == '+' || s == '-' {
		return 1
	} else if s >= '0' && s <= '9' {
		return 2
	} else {
		return 3
	}
}

func (fs *fmsSerice) getSum(s byte) {
	fs.state = fs.table[fs.state][fs.getState(s)]
	if fs.state == "Sign" {
		if s == '-' {
			fs.sign = -1
		}
	} else if fs.state == "Append" {
		fs.sum *= 10
		fs.sum += int64(s - '0')
		if fs.sign == -1 {
			fs.sum = min_(fs.sum, -1*math.MinInt32)
		} else {
			fs.sum = min_(fs.sum, math.MaxInt32)
		}
	}
}

func min_(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
