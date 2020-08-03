package string

func addString(num1, num2 string) string {
	var (
		ans              string
		flag             = false
		length1, length2 = len(num1), len(num2)
		index1, index2   = length1 - 1, length2 - 1
	)
	for index1 >= 0 && index2 >= 0 {
		tmpA := int(num1[index1] - '0')
		tmpB := int(num2[index2] - '0')
		tmpSum := tmpA + tmpB
		if flag {
			tmpSum++
		}
		if tmpSum >= 10 {
			flag = true
		} else {
			flag = false
		}
		ans = string((tmpSum%10)+'0') + ans
		index1--
		index2--
	}
	if index1 >= 0 {
		for flag && index1 >= 0 {
			tmpNum := int(num1[index1] - '0')
			tmpNum++
			if tmpNum >= 10 {
				flag = true
			} else {
				flag = false
			}
			ans = string((tmpNum%10)+'0') + ans
			index1--
		}
		if flag {
			ans = "1" + ans
		} else {
			ans = num1[:index1+1] + ans
		}
	} else if index2 >= 0 {
		for flag && index2 >= 0 {
			tmpNum := int(num2[index2] - '0')
			tmpNum++
			if tmpNum >= 10 {
				flag = true
			} else {
				flag = false
			}
			ans = string((tmpNum%10)+'0') + ans
			index2--
		}
		if flag {
			ans = "1" + ans
		} else {
			ans = num2[:index2+1] + ans
		}
	} else if index1 < 0 && index2 < 0 && flag {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
