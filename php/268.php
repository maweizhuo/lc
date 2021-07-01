class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function missingNumber($nums) {
        $res=0;
        foreach($nums as $k=>$v){
          $res ^=$k^$v;
        }
        return $res ^count($nums);
    }

    function missingNumber($nums) {
       $n=count($nums);
       $sum=$n*($n+1)/2;
//       $sum1=0;
//       foreach($nums as $k=>$v){
//            $sum1+=$v;
//       }
       $sum1=array_sum($nums);
       return $sum-$sum1;
    }
}