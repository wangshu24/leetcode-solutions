// func checkInclusion(s1 string, s2 string) bool {
    // reversed := ""

    // for i:=len(s1)-1; i > -1;i-- {
    //     reversed += string(s1[i])
    // }
    // fmt.Println(reversed)
    // return strings.Contains(s2, reversed) || strings.Contains(s2, s1)

    // s1Arr := []string{}

    // for i:=0; i< len(s1);i++{
    //     s1Arr = append(s1Arr, string(s1[i]))
    // }
    // fmt.Println(s1Arr)

    // chunks := []string{}

    // if len(s1) == len(s2) {
    //     found := 0
    //     clone := strings.Clone(s1)
    //     for i:=0;i< len(s1);i++{
    //           fmt.Println("@ char ", string(s2[i]))
    //         if strings.Contains(s1, string(s2[i])){
    //             clone = strings.Replace(clone, string(s2[i]), "" ,1)
    //             found++
    //         }
    //     }    
    //     fmt.Println("clone remain:", clone)
    //     if found == len(s1) && len(clone) == 0 {return true}   
    // }

    // for i:=0; i <= len(s2) - len(s1);i++ {
    //     if strings.Contains(s1, string(s2[i])) {
    //         clone := strings.Clone(s1)
    //         clone = strings.Replace(clone, string(s2[i]), "" ,1)
    //         found := 1
    //         for j:=1; j < len(s1);j++{
              
    //             if strings.Contains(clone, string(s2[i+j])){
    //                 clone = strings.Replace(clone, string(s2[i+j]), "" ,1)
    //                 found++
    //             }
    //         }
    //         if found == len(s1) {return true}
    //     }
    // }

    // return false
//}


func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)

    // Initialize counts for s1 and the first window of s2
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    // Slide the window over s2
    for i := 0; i < len(s2)-len(s1); i++ {
        if match(s1Count, s2Count) {
            return true
        }
        s2Count[s2[i]-'a']--
        s2Count[s2[i+len(s1)]-'a']++
    }

    // Check the last window
    return match(s1Count, s2Count)
}

func match(a, b []int) bool {
    for i := 0; i < 26; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}