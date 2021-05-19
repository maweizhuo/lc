<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
         $slow =0;
         $length = count($nums);
         for ($fast=1; $fast<$length; $fast++){
            if($nums[$fast] !=$nums[$slow]){
            $slow++;
              $nums[$slow]=$nums[$fast];
            }
         }
         return $slow+1;
          }
     }

    }

}