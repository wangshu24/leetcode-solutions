//IMOW

func decrypt(code []int, k int) []int {
    res := make([]int, len(code))
    if k == 0 {return res}

    kIndex := func(startInd, k int) int{
        if k>0 {
            return (startInd + k) % len(code)
        } else {

            newInd :=startInd -1 + len(code) + k
            if newInd >= len(code) {
                return newInd % len(code) 
            } else {
                return newInd
            }
        }
    }

    currKsum := 0
    copyK := int(math.Abs(float64(k)))
    for i:= 0; i < copyK; i++ {
        if k >0 {
            currKsum += code[i+1]
        } else {
            currKsum += code[i + len(code) + k]
        }

    }
    res[0] = currKsum
    fmt.Println(currKsum)

    for i:=1 ; i < len(code); i++ {
        if k > 0 {
            currKsum = currKsum - code[i] + code[kIndex(i, k)]
            res[i] = currKsum
        } else {
            fmt.Println("currKsum is: ", currKsum, " i is: ", i, " add " , code[i-1], " minus ", code[kIndex(i,k)], " at index: ", kIndex(i,k))
            currKsum = currKsum + code[i-1] - code[kIndex(i,k)]
            res[i] = currKsum 
        }
    }

    return res
}