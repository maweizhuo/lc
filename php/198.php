class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
       $n=count($nums);
       if($n==0){
         return 0;
       }
       if($n==1){
         return $nums[0];
       }
       $dp=[];
       $dp[0]=$nums[0];
       $dp[1]=max($nums[0],$nums[1]);
       for($i=2;$i<$n;$i++){
          $dp[$i]=max($dp[$i-2]+$nums[$i],$dp[$i-1]);
       }
       return $dp[$n-1];
    }

    // 优化
     function rob($nums) {
           $n=count($nums);
           if($n==0){
             return 0;
           }
           if($n==1){
             return $nums[0];
           }

           $first=$nums[0];
           $second=max($nums[0],$nums[1]);
           for($i=2;$i<$n;$i++){
              $max=max($first+$nums[$i],$second);
              $first=$second;
              $second=$max;
           }
           return $second;
        }

}