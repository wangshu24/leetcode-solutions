public class Solution {
    public bool IsPalindrome(int x) {
        var number = x.ToString();
        int l = 0, r = number.Length-1;
        while (l < r) {
            if (number[l] != number[r]) {
                return false;
            }
            l++;
            r--;
        }
        return true;
    }
}