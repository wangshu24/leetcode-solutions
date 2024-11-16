class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function resultsArray($nums, $k) {
        $res = [];
    $l = 0;
    $consecCnt = 1;

    for ($r = 0; $r < count($nums); $r++) {
        if ($r > 0 && $nums[$r - 1] + 1 === $nums[$r]) {
            $consecCnt++;
        }

        if ($r - $l + 1 > $k) {
            if ($nums[$l] + 1 === $nums[$l + 1]) {
                $consecCnt--;
            }
            $l++;
        }

        if ($r - $l + 1 === $k) {
            $res[] = $consecCnt === $k ? $nums[$r] : -1;
        }
    }

    return $res;
    }
}