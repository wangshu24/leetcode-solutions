func minLength(s string) int {
	for {
		ok, subStr := hasEither(s)
		if !ok {
			break
		}
		s = strings.Replace(s, subStr, "", 1)
	}

	return len(s)
}

func hasEither(s string) (bool, string) {
	if strings.Contains(s, "AB") {
		return true, "AB"
	} else if strings.Contains(s, "CD") {
		return true, "CD"
	} else {
		return false, ""
	}
}

// func minLength(s string) int {
//     if !strings.Contains(s, "AB") && !strings.Contains(s, "CD") {return len(s)}
//     found := true
//     for found {
//         if strings.Contains(s, "AB"){
//             s = strings.Replace(s, "AB", "", 1)
//         } else if strings.Contains(s, "CD"){
//             s = strings.Replace(s, "CD", "", 1)
//         } else  { found = false}
//     }

//     return len(s)
// }