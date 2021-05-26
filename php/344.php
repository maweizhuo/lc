<?php

class Solution {

         /**
         * @param String[] $s
         * @return NULL
         */
        function reverseString(&$s) {
//            for($l=0,$r=count($s)-1;$l<$r;$l++,$r--){
//               $temp=$s[$l];
//               $s[$l]=$s[$r];
//               $s[$r]=$temp;
//            }

            $l=0;
            $r=count($s)-1;
            while($l<$r){
                $temp=$s[$l];
                $s[$l]=$s[$r];
                $s[$r]=$temp;
                $l++;
                $r--;
            }

        }

}