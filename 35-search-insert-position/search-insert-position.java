class Solution {
    public int searchInsert(int[] nums, int target) {
        List<Integer> arrlist 
              = new ArrayList<>();
        
        for (int i : nums) {
            arrlist.add(i);  
        }       
        arrlist.add(target);  
        Collections.sort(arrlist);

         return arrlist.indexOf(target);
    }
}