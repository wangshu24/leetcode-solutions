func dividePlayers(skill []int) int64 {
    result := 0
    totalSkill := 0
    maxIndex := len(skill)-1
    if maxIndex == 1 {
        result = skill[0]*skill[1]
        return int64(result)
    }
    

    for i:=0; i<len(skill); i++ {
        totalSkill+= skill[i]
    }
    totalSkill = totalSkill / (len(skill)/2)
    slices.Sort(skill)
    fmt.Println(totalSkill)
    
    if skill[len(skill)-1] >= totalSkill {return -1}
    for i:=0; i<len(skill)/2;i++{
        if skill[i] + skill[maxIndex-i] == totalSkill {
            result += skill[i]*skill[maxIndex-i] 
        } else {return -1}
    }
    
    return int64(result)
}