func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

//Instantiate map of alphabet, with each index equivalent of byte value of a character denoted by char 'a' (which is 97 value in ASCII) so a would default to index 0
//With index being the character value, the value at index is the occurence of said character
//Recording the character occurence using array and byte value of character -> hashmap 
    s1Count := make([]int, 26)
    s2Count := make([]int, 26)

    for i:=0; i< len(s1);i++{
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    if len(s1) == len(s2) {return match(s1Count, s2Count)}

//Now we apply sliding window technique with hashmap
    for i:=0;i < len(s2)-len(s1);i++ {
        fmt.Println(i)
        if match(s1Count, s2Count) {
            return true
        }
        s2Count[s2[i]-'a']--
        s2Count[s2[i+len(s1)]-'a']++
    }
//The sliding window will stop short at the last index, ie, if s1 = "adc" and s2 = "dcda", it will stop right at "cda"
//and not execute that last comparision

    return match(s1Count, s2Count)
}

func match(s1Count, s2Count []int) bool {
    for i:=0; i< 26; i++{
        if s1Count[i] != s2Count[i]{
            return false
        }
    }
    return true
}



