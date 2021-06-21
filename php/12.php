class Solution {

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num) {

      $symbolValues=[
      1000=>"M",
      900=>"CM",
      500=>"D",
      400=>"CD",
      100=>"C",
      90=>"XC",
      50=>"L",
      40=>"XL",
      10=>"X",
      9=>"IX",
      5=>"V",
      4=>"IV",
      1=>"I",
      ];
      $str=[];
      foreach($symbolValues as $val=>$symbol){
         while($num>=$val){
           $num-=$val;
           $str[]=$symbol;
         }

      }
       return implode('',$str);
      }

}