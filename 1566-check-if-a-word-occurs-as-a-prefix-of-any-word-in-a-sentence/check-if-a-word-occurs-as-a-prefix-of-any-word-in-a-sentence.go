func isPrefixOfWord(sentence string, search string) int {
    words := strings.Split(sentence, " ")
    res := -1
    for ind,word := range words {
        if strings.HasPrefix(word, search){
            res = ind + 1
            break
        }
    }

    return res
}