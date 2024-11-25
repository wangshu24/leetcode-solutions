import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
    var res strings.Builder
    res.Grow(len(word1) + len(word2))
    lengthMin := min(len(word1), len(word2))
    
    for i:= 0; i < lengthMin; i++{
        res.WriteString(string(word1[i]))
        res.WriteString(string(word2[i]))
    }

    if len(word1) < len(word2){
        res.WriteString(string(word2[len(word1):]))
    } else {
        res.WriteString(string(word1[len(word2):]))
    }

    return res.String()
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
