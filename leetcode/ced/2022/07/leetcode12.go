package _7

func intToRoman(num int) string {
	var res string
	for num > 0 {
		res += getNum(&num)
	}
	return res
}

func getNum(num *int) string {
	if *num >= 1000 {
		*num -= 1000
		return "M"
	} else if *num >= 900 {
		*num -= 900
		return "CM"
	} else if *num >= 500 {
		*num -= 500
		return "D"
	} else if *num >= 400 {
		*num -= 400
		return "CD"
	} else if *num >= 100 {
		*num -= 100
		return "C"
	} else if *num >= 90 {
		*num -= 90
		return "XC"
	} else if *num >= 50 {
		*num -= 50
		return "L"
	} else if *num >= 40 {
		*num -= 40
		return "XL"
	} else if *num >= 10 {
		*num -= 10
		return "X"
	} else if *num >= 9 {
		*num -= 9
		return "IX"
	} else if *num >= 5 {
		*num -= 5
		return "V"
	} else if *num >= 4 {
		*num -= 4
		return "IV"
	} else {
		*num -= 1
		return "I"
	}
}
