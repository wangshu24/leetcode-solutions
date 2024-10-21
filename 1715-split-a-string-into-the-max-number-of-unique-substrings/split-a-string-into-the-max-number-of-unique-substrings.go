// Go Solution
func split(s string, st map[string]bool) int {
    if len(s) == 0 { // Base case: when the string is empty
        return len(st) // Return the size of the set (unique substrings)
    }

    mx := len(st)
    // Try all possible prefixes from length 1 to the full string
    for i := 1; i <= len(s); i++ {
        first := s[:i] // Get the current prefix
        if !st[first] { // If the prefix is unique
            st[first] = true // Insert it into the set
            mx = max(mx, split(s[i:], st)) // Recurse on the remaining string
            delete(st, first) // Backtrack: remove the prefix from the set
        }
    }

    return mx // Return the maximum number of unique substrings
}

func maxUniqueSplit(s string) int {
    st := make(map[string]bool) // Set to store unique substrings
    return split(s, st) // Call the helper function
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}