func lengthOfLastWord(s string) int {
    trimmed := strings.Trim(s, " ")
    stringArr := strings.Split(trimmed, " ")
    return len(stringArr[len(stringArr)-1])
}