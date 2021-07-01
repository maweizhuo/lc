class Solution {

    /**
     * @param Integer $numRows
     * @return Integer[][]
     */
    function generate($numRows) {
       $ans=[];
       for($i=0;$i<$numRows;$i++){
          for($j=0;$j<=$i;$j++){
            if($j==0 || $j==$i){
            $ans[$i][]=1;
            }else{
            $ans[$i][$j]=$ans[$i-1][$j]+$ans[$i-1][$j-1];
            }
          }
       }
       return $ans;
    }
}