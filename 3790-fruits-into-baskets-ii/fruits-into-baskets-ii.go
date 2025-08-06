func numOfUnplacedFruits(fruits []int, baskets []int) int {
    used := make([]bool, len(fruits))
    left := len(fruits)

    for i:=range fruits {
        for j:= range baskets {
            if used[j] || baskets[j] < fruits[i] {
                continue
            }

            used[j] = true
            left--
            break
        }
    }
    return left
}