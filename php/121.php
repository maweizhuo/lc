class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
      $length=count($prices);
      if($length<2){
       return 0;
      }
      $min=$prices[0];
      $ans=0;
      for($i=1;$i<$length;$i++){
         $tmp=$prices[$i]-$min;
         if($tmp>$ans){
           $ans=$tmp;
         }else if ($prices[$i]<$min){
            $min=$prices[$i];
         }
      }
      return $ans;
    }
}