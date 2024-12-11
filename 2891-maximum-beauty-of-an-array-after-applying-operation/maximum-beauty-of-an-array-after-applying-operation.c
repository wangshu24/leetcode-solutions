int maxInArray(int *nums, int size) {
    int mx = nums[0];
    for (int i = 1; i < size; i++) {
        if (nums[i] > mx) {
            mx = nums[i];
        }
    }
    return mx;
}

// Utility function to return the maximum of two integers
int max(int a, int b) {
    return (a > b) ? a : b;
}

// Utility function to return the minimum of two integers
int min(int a, int b) {
    return (a < b) ? a : b;
}

int maximumBeauty(int* nums, int numsSize, int k) {
     int mx = maxInArray(nums, numsSize);
    int *diff = (int *)calloc(mx + 2, sizeof(int));

    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        diff[max(x - k, 0)]++;
        diff[min(x + k + 1, mx + 1)]--;
    }

    int count = 0, ans = 0;
    for (int i = 0; i <= mx + 1; i++) {
        count += diff[i];
        ans = max(ans, count);
    }

    free(diff);
    return ans;
}