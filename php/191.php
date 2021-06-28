class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function hammingWeight($n) {
       $ones=0;
       for($i=0;$i<32;$i++){
          if((1<<$i&$n)>0){
            $ones++;
          }
       }
       return $ones;
    }

    function hammingWeight($n) {
          $ones=0;
         while($n>0 ){
             $n&=$n-1;
              $ones++;
         }
     return $ones;
    }
}