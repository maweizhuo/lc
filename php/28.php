<?php

class Solution {

          /**
          * @param String $haystack
          * @param String $needle
          * @return Integer
          */
         function strStr($haystack, $needle) {
            if($needle=='')return 0;
            $slen=strlen($haystack);
            $nlen=strlen($needle);
            for ($i=0;$i<=$slen-$nlen;$i++){
               if(substr($haystack,$i,$nlen)==$needle)return $i;
            }
            return -1;
         }

          function strStr_2($haystack, $needle) {
             // 暴力
             if($needle=='')return 0;
             $slen=strlen($haystack);
             $nlen=strlen($needle);
             $i=$j=0;
             while($i<$slen && $j<$nlen){
                if($haystack[$i]==$needle[$j]){
                  $i++;
                  $j++;
                }else{
                  $i=$i-$j+1;
                  $j=0;  //置为起点
                }
                if($j==$nlen){
                   return $i-$j;
                }
             }
             return -1;
          }

}