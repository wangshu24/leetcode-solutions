func longestDiverseString(a int, b int, c int) string {
	s := ""
	i := 0
	for {
		if a == 0 && b == 0 && c == 0 {
			//no characters available, break
			break
		}
		flagA, flagB, flagC := false, false, false
		if i >= 2 && s[i-1] == s[i-2] {
			flagA = s[i-1] == 'a'
			flagB = s[i-1] == 'b'
			flagC = s[i-1] == 'c'
		}
		char := getNextChar(a, b, c, flagA, flagB, flagC)
		if char == 0 {
			break
		} else {
            //add a new char to the string, subtract 1 from the group
			s += string(char)
			if char == 'a' {
				a--
			}
			if char == 'b' {
				b--
			}
			if char == 'c' {
				c--
			}
		}
        //move i one step forward
		i++
	}
	return s
}

func getNextChar(a, b, c int, flagA, flagB, flagC bool) byte {
	//mask out character we don't want to pick from
	if flagA {
		a = 0
	}
	if flagB {
		b = 0
	}
	if flagC {
		c = 0
	}
	mostChar := charWithMax(a, b, c)
	if mostChar == 0 {
		//no character available
		return 0
	}
	return mostChar
}

func charWithMax(a, b, c int) byte {
	if a >= b && a >= c && a > 0 {
		return 'a'
	} else if b >= a && b >= c && b > 0 {
		return 'b'
	}
	if c > 0 {
		return 'c'
	}
	//no character available
	return 0
}