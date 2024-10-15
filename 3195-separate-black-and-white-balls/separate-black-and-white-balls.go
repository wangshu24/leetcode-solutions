// import ("heap")

// func ()

func minimumSteps(s string) int64 {
    // 1 = BLACK = 49
    // 0 = WHITE = 48
    if len(s) == 1 || !strings.Contains(s, "1") {return 0}

    charArr := []byte{}
    for i:=0; i<len(s);i++{
        charArr = append(charArr, s[i])
    }
    
    count:=0
    black := slices.Index(charArr, 49)

    leftGap:=0
    for i:=black+1;i<len(charArr);i++{
        if charArr[i] == 48 {
            count+= i - black + leftGap
            black = i 
            continue
        }
        if charArr[i] == 49 {
            leftGap+=i-black
            black  = i
            continue
        }
    }

    return int64(count)
}