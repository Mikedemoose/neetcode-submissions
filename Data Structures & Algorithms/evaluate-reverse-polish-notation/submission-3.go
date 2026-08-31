func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	if a, ok := isOperand(tokens[0]); !ok {
		stack = append(stack, a)
	}

	for i:=1; i<len(tokens); i++ {
		if a, ok := isOperand(tokens[i]); !ok {
			stack = append(stack, a)
			continue
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		stack = append(stack, doOperation(a, b, tokens[i]))
	}

	return stack[0]
}

func isOperand(a string) (int, bool) {
	if a == "+" || a == "-" || a == "*" || a == "/" {
		return 0, true
	}
	finalInt := 0
	isNegative := false
	for i:= range len(a) {
		if a[i] == '-' {
			isNegative = true
			continue
		}
		finalInt*=10
		finalInt+= int(a[i]-'0')
	}
	if isNegative {
		finalInt = -finalInt
	}
	return finalInt, false
}

func doOperation(a, b int, op string) int {
	if op == "+" {
		return a+b
	} else if op == "-" {
		return a-b
	} else if op == "*" {
		return a*b
	} 
	return a/b
}
