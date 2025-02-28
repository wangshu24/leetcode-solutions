func shortestCommonSupersequence(str1 string, str2 string) string {
	// Make sure str1 is the longest
	if len(str1) < len(str2) {
		return shortestCommonSupersequence(str2, str1)
	}

	dp := make([][]int, len(str1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(str2)+1)
	}

	for i := len(str1) - 1; i >= 0; i-- {
		for j := len(str2) - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	var res strings.Builder
    i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			// Characters are the same; copy character to result and advance both strings.
			res.WriteByte(str1[i])
			i++
			j++
		} else if dp[i][j+1] > dp[i+1][j] {
			// Picking str2[j], because the str1[i:] and str2[j+1:] have longest subsequence
			res.WriteByte(str2[j])
			j++
		} else {
			// if equal or less, we want to advance str1 first as it is bigger
			res.WriteByte(str1[i])
			i++
		}
	}

	// Simply copy the remaining characters from both strings
	res.WriteString(str1[i:])
	res.WriteString(str2[j:])
	return res.String()
}