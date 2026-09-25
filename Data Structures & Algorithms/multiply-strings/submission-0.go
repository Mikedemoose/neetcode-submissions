func multiply(num1 string, num2 string) string {
	product := "0"
	if num1 == "0" || num2 == "0" {
		return product
	}

	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}

	short := len(num2)

	for i := range short {
		num := getNumericVal(num2[short-i-1])
		currProd := multiplyStringWithNum(num1, num)

		for range i {
			currProd += "0"
		}

		product = addStrings(product, currProd)
	}

	return product
}

func multiplyStringWithNum(num1 string, num2 int) string {
	if num2 == 0 {
		return "0"
	}

	product := num1

	for range num2-1 {
		product = addStrings(product, num1)
	}

	return product
}

func addStrings(num1, num2 string) string {
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}

	long, short := len(num1), len(num2)

	res := ""
	carry := 0

	for i := range short {
		sum := getNumericVal(num1[long-1-i]) + getNumericVal(num2[short-1-i]) + carry
		carry = sum/10
		sum = sum%10

		res = getStringVal(sum) + res
	}

	for i := range long-short {
		ind := long-short-1-i

		sum := getNumericVal(num1[ind]) + carry
		carry = sum/10
		sum = sum % 10

		res = getStringVal(sum) + res
	}

	if carry > 0 {
		res = getStringVal(carry) + res
	}

	return res
}

func getNumericVal(char byte) int {
	return int(char-'0')
}

func getStringVal(val int) string {
	return string('0'+val)
}
