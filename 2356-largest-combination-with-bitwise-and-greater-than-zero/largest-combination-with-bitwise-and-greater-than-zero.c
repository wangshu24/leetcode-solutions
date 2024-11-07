int largestCombination(int* candidates, int candidatesSize) {
    int maxSize = 0;
    for (int i = 0; i < 24; i++) {
        int count = 0;
        for (int j = 0; j < candidatesSize; j++) {
            if ((candidates[j] & (1 << i)) != 0) {
                count++;
            }
        }
        if (count > maxSize) {
            maxSize = count;
        }
    }
    return maxSize;
}