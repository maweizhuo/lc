class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer $m
     * @param Integer[] $nums2
     * @param Integer $n
     * @return NULL
     */
    function merge(&$nums1, $m, $nums2, $n) {
   // 从后向前推进
        while ($n>0){
            if($nums1[$m-1]>$nums2[$n-1]){
              $nums1[$m+$n-1]=$nums1[$m-1];
              $m--;
            }else{
                $nums1[$m+$n-1] = $nums2[$n-1];
                $n--;
            }
        }
    }
}