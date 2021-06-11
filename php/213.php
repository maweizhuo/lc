class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
       $n=count($nums);
       if($n==1){
         return $nums[0];
       }
       if($n==2){
         return max($nums[0],$nums[1]);
       }
       $src_nums=$nums;
//       array_pop($nums);// 移除最后一个
//       array_shift($src_nums); // 移除第一个
//       return max($this->steal($nums),$this->steal($src_nums));
       return max($this->steal(array_slice($nums,1,$n-1)),$this->steal(array_slice($nums,0,$n-1)));
    }

    function steal($nums){
      $first=$nums[0];
      $second=max($nums[0],$nums[1]);
//     for($i=2;$i<count($nums);$i++){
//        $max=max($first+$nums[$i],$second);
//        $first=$second;
//        $second=$max;
//     }
      foreach($nums as $key=>$val){
        if (in_array($key,[0,1])){
            continue;
        }
       $max=max($first+$val,$second);
       $first=$second;
       $second=$max;
     }
     return $second;
    }

     function steal($nums){
          $first=0;
          $second=0;
          foreach($nums as $key=>$val){
           $max=max($first+$val,$second);
           $first=$second;
           $second=$max;
         }
         return $second;
     }

}