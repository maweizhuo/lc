<?php

class Solution {

         /**
         * @param Integer[][] $matrix
         * @return NULL
         */
        function rotate(&$matrix) {
            $length=count($matrix);
            $len_i=floor($length/2);
             for ($i=0;$i<$len_i;$i++) {
                for($j=$i;$j<$length-$i-1;$j++) {
                   $m=$length-$j-1;
                   $n=$length-$i-1;
                   $temp=$matrix[$i][$j];
                   $matrix[$i][$j]=$matrix[$m][$i];
                   $matrix[$m][$i]=$matrix[$n][$m];
                   $matrix[$n][$m]=$matrix[$j][$n];
                   $matrix[$j][$n]=$temp;
                }
             }
        }

}