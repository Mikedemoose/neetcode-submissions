func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	
	for i:=0; i<len(tokens); i++ {
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
	val, _ := strconv.Atoi(a)
	return val, false
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
