func numOfUnplacedFruits(fruits []int, baskets []int) int {
    left := len(fruits) 
    for i:=0; i < len(fruits);i++ {
        for j:=0; j < len(fruits);j++ {
            if fruits[i] <= baskets[j]{
                baskets[j] = 0
                left--
                break
            }
        }
    }
    return left
}