// func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    // strings.Trim(sentence1, " ")
    // strings.Trim(sentence2, " ")

    // s1Arr := strings.Split(sentence1, " ")
    // s2Arr := strings.Split(sentence2, " ")

    // shortArr := []string{}
    // longArr := []string{}

    // if len(s1Arr)>len(s2Arr){
    //     longArr = s1Arr
    //     shortArr = s2Arr
    // }else{
    //     longArr = s2Arr
    //     shortArr = s1Arr
    // }

    // if len(shortArr) == 1 {
    //     result := shortArr[0] == longArr[0] ||  shortArr[0] == longArr[len(longArr)-1]
    //     return result
    // }

    // if len(shortArr) == len(longArr) {
    //     for i:=0; i< len(shortArr);i++{
    //         if string(shortArr[i]) != string(longArr[i]){
    //             return false
    //             }
    //     }
    // }

    // long:= ""
    // for i:=0; i < len(longArr);i++ {
    //     long = long + " " + string(longArr[i])
    // }

    // for i:=0; i < len(shortArr);i++{
    //     fmt.Println(shortArr[i])
    //     long = strings.TrimPrefix(long, string(shortArr[i]))
    //     long = strings.Trim(long, " ")
    //     fmt.Println(long)
    // }

    // for i:=0; i < len(shortArr);i++{
    //     long = strings.TrimSuffix(long, shortArr[len(shortArr)-1-i])
    //     long= strings.Trim(long, " ")
    // }

    // longRemainArr := strings.Split(long, " ")
    // fmt.Println(longRemainArr)
    // for i:= 0; i < len(longRemainArr);i++{
    //     if slices.Contains(shortArr, longRemainArr[i]){
    //         return false
    //     }
    // }

    //Do count of Prefix, suffix then validate with sum of pref + suff
    
    // return true
// }



func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    s1Arr := strings.Split(sentence1, " ")
    s2Arr := strings.Split(sentence2, " ")
    
    // if len(s1Arr) ==1 && len(s2Arr) == 1 {
    //     return s1Arr[0] == s2Arr[0]
    // }

    length := 0
    if len(s1Arr) < len(s2Arr){
        length = len(s1Arr)
    } else {length = len(s2Arr) }

    start, end := 0,0

    //interlapse :=0
    for start<length && s1Arr[start] == s2Arr[start] {
        start++
        }
    
    
    l1,l2 := len(s1Arr), len(s2Arr)
    for end<length && s1Arr[l1-1-end] == s2Arr[l2-1-end] {
        fmt.Println(s1Arr[l1-1-end])
        fmt.Println(s2Arr[l2-1-end])
        end++
        }


    if start + end >= length {return true}

    return false
}


