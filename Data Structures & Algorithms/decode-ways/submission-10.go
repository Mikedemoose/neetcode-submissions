func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
    numWays := make([]int, len(s)+1)
	numWays[0], numWays[1] = 1, 1
	
	for i:=1; i<len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			numWays[i+1] = numWays[i-1]
		} else if s[i-1] == '1' {
			numWays[i+1] = numWays[i] + numWays[i-1]
		} else if s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			numWays[i+1] = numWays[i] + numWays[i-1]
		} else {
			numWays[i+1] = numWays[i]
		}
	}

	return numWays[len(s)]
}
