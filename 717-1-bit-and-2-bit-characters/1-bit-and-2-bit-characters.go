func isOneBitCharacter(bits []int) bool {
    if len(bits) == 1 {
        if bits[0] == 1 {return false} else {return true}
    }
    
    ind :=0
    for i:=0; i < len(bits)-1; i++ {
        if bits[i] == 1 {i++}
        ind = i
    }
    if ind == len(bits)-1 {return false}
   
    return true
}