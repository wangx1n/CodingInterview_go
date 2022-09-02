package _8

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	contactMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var nowRes string
	getRes(contactMap, digits, 0, nowRes)
	defer func() { res = make([]string, 0) }()
	return res
}

func getRes(contactMap map[byte]string, digits string, index int, nowRes string) {
	if index == len(digits) {
		res = append(res, nowRes)
		return
	}
	c := digits[index]
	nowNum := contactMap[c]
	n := len(nowNum)
	for i := 0; i < n; i++ {
		nowRes += string(nowNum[i])
		getRes(contactMap, digits, index+1, nowRes)
		nowRes = nowRes[0 : len(nowRes)-1]
	}
}
