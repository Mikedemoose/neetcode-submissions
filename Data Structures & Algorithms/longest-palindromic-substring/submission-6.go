func longestPalindrome(s string) string {
    n := len(s)
    dp := make([][]bool, n)
    for i := range n {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }

    maxLength, ind_left, ind_right := 1, 0, 0

    for length:=1; length<=n; length++ {
        for i:=0; i+length<n+1; i++ {
            
            j:= i+length-1

            if s[i] == s[j] {
                if length == 1 || length == 2 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
                if dp[i][j] {
                    currLength := j-i+1
                    if currLength > maxLength {
                        maxLength = currLength
                        ind_left, ind_right = i, j
                    }
                }
            }
        }
    }

    return s[ind_left:ind_right+1]
}
