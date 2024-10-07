// func areSentencesSimilar(sentence1 string, sentence2 string) bool {
//     s1Arr := strings.Split(sentence1, " ")
//     s2Arr := strings.Split(sentence2, " ")

//     shortArr := []string{}
//     longArr := []string{}

//     if len(s1Arr)>len(s2Arr){
//         longArr = s1Arr
//         shortArr = s2Arr
//     }else{
//         longArr = s2Arr
//         shortArr = s1Arr
//     }

//     if len(shortArr) == 1 {
//         result := shortArr[0] == longArr[0] ||  shortArr[0] == longArr[len(longArr)-1]
//         return result
//     }

//     if len(shortArr) == len(longArr) {
//         for i:=0; i< len(shortArr);i++{
//             if string(shortArr[i]) != string(longArr[i]){
//                 return false
//                 }
//         }
//     }

//     long:= ""
//     for i:=0; i < len(longArr);i++ {
//         long = long + " " + string(longArr[i])
//     }

//     for i:=0; i < len(shortArr);i++{
//         fmt.Println(shortArr[i])
//         long = strings.TrimPrefix(long, string(shortArr[i]))
//         long = strings.Trim(long, " ")
//         fmt.Println(long)
//     }

//     for i:=0; i < len(shortArr);i++{
//         long = strings.TrimSuffix(long, shortArr[len(shortArr)-1-i])
//         long= strings.Trim(long, " ")
//     }

//     longRemainArr := strings.Split(long, " ")
//     fmt.Println(longRemainArr)
//     for i:= 0; i < len(longRemainArr);i++{
//         if slices.Contains(shortArr, longRemainArr[i]){
//             return false
//         }
//     }

//     // Do count of Prefix, suffix then validate with sum of pref + suff
    
//     return true
// }



func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    s1Arr := strings.Split(sentence1, " ")
    s2Arr := strings.Split(sentence2, " ")

    length := 0
    if len(s1Arr) < len(s2Arr){
        length = len(s1Arr)
    } else {length = len(s2Arr) }

    //start, end := 0,0

    interlapse :=0
    for i:=0; i<length;i++ {
        if s1Arr[i] != s2Arr[i] {
            break
        }
        interlapse++
    }
    
    l1,l2 := len(s1Arr), len(s2Arr)
    for i:=0; i<length;i++ {
        if s1Arr[l1-1-i] != s2Arr[l2-1-i] {
            break
        }
        interlapse++
    }

    if interlapse >= length {return true}

    return false
}

