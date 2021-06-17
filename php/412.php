class Solution {

    /**
     * @param Integer $n
     * @return String[]
     */
    function fizzBuzz($n) {
       $arr = [];
       for($i=1;$i<=$n;$i++){
           switch($i){
             case $i %3 ==0 && $i%5==0:
              $arr[]='FizzBuzz';
               break;
            case $i %3 ==0:
              $arr[]='Fizz';
               break;
            case $i%5==0:
              $arr[]='Buzz';
              break;
             default:
              $arr[]=(string)$i;
           }
       }
       return $arr;
    }
}