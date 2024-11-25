// Solution 2:
// Using S1+S2=S2+S1, Commutative property for strings!
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    } else {
        return str1[:gcd(len(str1), len(str2))]
    }
}