class Solution {

    /**
     * @param Integer $x
     * @param Integer $y
     * @return Integer
     */
    function hammingDistance($x, $y) {
       $ans=0;
       $s=$x^$y;
       while($s>0){
          $ans+=$s&1;
          $s>>=1;
       }
       return $ans;
    }
}