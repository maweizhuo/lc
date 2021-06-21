class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
      $symbolValues=[
        'I'=>1,
        'V'=>5,
        'X'=>10,
        'X'=>10,
        'L'=>50,
        'C'=>100,
        'D'=>500,
        'M'=>1000,
      ];
     $ans=0;
      $n=strlen($s);
      for($i=0;$i<$n;$i++){
        $val=$symbolValues[$s[$i]];
        if($i<$n-1 &&  $val<$symbolValues[$s[$i+1]]){
           $ans-=$val;
        }else{
           $ans+=$val;
        }
      }
     return $ans;
    }
}