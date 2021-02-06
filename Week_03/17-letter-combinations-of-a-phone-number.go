package Week_03

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	res := []string{}

	recursion(m, 0, digits, "", &res)

	return res
}

func recursion(m map[string]string, cur int, digits, path string, res *[]string) {
	// 递归完成，将 path 添加到结果列表中
	if cur == len(digits) {
		*res = append(*res, path)
		return
	}

	// 获取当前数字，并尝试所有可能
	for _, s := range m[string(digits[cur])] {
		str := string(s)

		path = path + str
		recursion(m, cur+1, digits, path, res)
		path = path[:len(path)-1]
	}
}

// 数字按钮 -> 字母
var m = map[string]string {
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",
}
