func longestPalindrome(s string) string {
    n := len(s)

    if n == 0 {
        return ""
    }

    dp := make([][]bool, n)

    for i := range dp {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }

    start := 0
    maxLength := 1

    for length := 2; length <= n; length++ {
        for i := 0; i+length <= n; i++ {
            j := i + length - 1

            if s[i] == s[j] &&
                (length <= 2 || dp[i+1][j-1]) {

                dp[i][j] = true

                if length > maxLength {
                    start = i
                    maxLength = length
                }
            }
        }
    }

    return s[start : start+maxLength]
}