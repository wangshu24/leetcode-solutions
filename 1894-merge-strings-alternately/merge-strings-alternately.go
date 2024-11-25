import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
    var res strings.Builder
    res.Grow(len(word1) + len(word2))
    lengthMin := min(len(word1), len(word2))
    lengthMax := max(len(word1), len(word2))

    for i:= 0; i < lengthMin; i++{
        res.WriteString(string(word1[i]))
        res.WriteString(string(word2[i]))
    }

    longer := ""
    if len(word1) == lengthMax {
        longer = word1
    }  else { longer = word2 }

    for i:=lengthMin; i < lengthMax; i++{
        res.WriteString(string(longer[i]))
    }

    return res.String()
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a <b {
        return b
    }
    return a
}