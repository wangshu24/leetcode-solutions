func shortestCommonSupersequence(str1 string, str2 string) string {
   lenStr1,lenStr2 := len(str1),len(str2)
	prev := make([]string, lenStr2+1)
	for col := 0; col <= lenStr2; col++ {
		prev[col] = string(str2[:col])
	}

	for row := 1; row <= lenStr1; row++ {
		currRow := make([]string, lenStr2+1)
		currRow[0] = str1[:row]
		for col := 1; col <= lenStr2; col++ {
			if str1[row-1] == str2[col-1] {
				currRow[col] = prev[col-1] + string(str1[row-1])
			} else {
				pickS1 := prev[col]
				pickS2 := currRow[col-1]
				if len(pickS1) < len(pickS2) {
					currRow[col] = pickS1 + string(str1[row-1])
				}else{
                    currRow[col] = pickS2 + string(str2[col-1])
                }
			}
		}
		prev = currRow
      
	}

    return prev[lenStr2]

}
