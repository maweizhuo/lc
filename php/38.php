<?php

class Solution {

    /**
     * @param Integer $n
     * @return String
     */
    function countAndSay($n) {
       $result="1";
       for ($i=1;$i<$n;$i++){
          $temp="";

          for ($j=0;$j<strlen($result);$j++){
             $count=1;
             while ($j+1<strlen($result) && $result[$j]==$result[$j+1]){
              $count++;
              $j++;
             }
             $temp.=$count.$result[$j];
          }
          $result=$temp;
       }
       return $result;
    }

}