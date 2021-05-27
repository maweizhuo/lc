<?php

class Solution {

         /**
         * @param String $s
         * @return Integer
         */
        function firstUniqChar($s) {
            $hash=[];
            for($i=0;$i<strlen($s);$i++){
              $hash[$s[$i]][]=$i;
            }
            foreach($hash as $v){
              if(count($v) ==1){
                return $v[0];
              }
            }
            return -1;
        }


}