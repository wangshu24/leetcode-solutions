impl Solution {
    pub fn maximum_beauty(nums: Vec<i32>, k: i32) -> i32 {
        use std::cmp::{max, min};

        let mx = *nums.iter().max().unwrap(); // Find the maximum element in nums
        let mut diff = vec![0; (mx + 2) as usize];

        for &x in &nums {
            diff[max(x - k, 0) as usize] += 1;
            diff[min(x + k + 1, mx + 1) as usize] -= 1;
        }

        let mut count = 0;
        let mut ans = 0;
        for &x in &diff {
            count += x;
            ans = max(ans, count);
        }

        ans
    }
}