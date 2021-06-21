class Solution {

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isPowerOfThree($n) {
       if($n>1){
         while($n%3==0){
             $n/=3;
         }
       }
       return $n==1;
    }

}