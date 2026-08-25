func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
    numWays := make([]int, len(s))
	numWays[0] = 1

	for i:=1; i<len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			if i>1 {
				numWays[i] = numWays[i-2]
			} else {
				numWays[i] = 1
			}
		} else if s[i-1] == '1' {
			numWays[i] = numWays[i-1] 
			if i>1 {
				numWays[i] += numWays[i-2]
			} else {
				numWays[i] += 1
			}
		} else if s[i-1] == '2' && (s[i] >= '1' && s[i] <= '6') {
			if i>1 {
				numWays[i] = numWays[i-2] + numWays[i-1]
			} else {
				numWays[i] = numWays[i-1] + 1
			}
		} else {
			numWays[i] = numWays[i-1]
		}
	}

	return numWays[len(s)-1]
}
