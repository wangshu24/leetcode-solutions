public class Solution {
    static int MaxInArray(int[] nums)
    {
        int mx = nums[0];
        foreach (int num in nums)
        {
            if (num > mx) mx = num;
        }
        return mx;
    }
    
    public int MaximumBeauty(int[] nums, int k) {
        int mx = MaxInArray(nums);
        int[] diff = new int[mx + 2];

        foreach (int x in nums)
        {
            diff[Math.Max(x - k, 0)]++;
            diff[Math.Min(x + k + 1, mx + 1)]--;
        }

        int count = 0, ans = 0;
        foreach (int x in diff)
        {
            count += x;
            ans = Math.Max(ans, count);
        }

        return ans;
    }
}