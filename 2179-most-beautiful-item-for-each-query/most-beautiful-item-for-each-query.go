func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] < items[j][0]
    })
    prefMaxBeauty := make([]int, len(items))
    for i, item := range items {
        prefMaxBeauty[i] = item[1]
        if i > 0 && prefMaxBeauty[i - 1] > prefMaxBeauty[i] {
            prefMaxBeauty[i] = prefMaxBeauty[i - 1]
        }
    }
    answers := make([]int, 0, len(queries))
    for _, query := range queries {
        pos, _ := slices.BinarySearchFunc(items, query + 1, func(el []int, target int) int {
            return el[0] - target
        })
        answer := 0
        if pos > 0 {
            answer = prefMaxBeauty[pos - 1]
        }

        answers = append(answers, answer)
    }
    return answers
}