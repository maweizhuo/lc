class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maxSubArray($nums) {
      $max=$nums[0];
      for($i=1;$i<count($nums);$i++){
         if($nums[$i]+$nums[$i-1]>$nums[$i]){
           $nums[$i]+=$nums[$i-1];
         }
         if($nums[$i]>$max){
           $max=$nums[$i];
         }
      }
      return $max;
    }
}