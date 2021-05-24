<?php

class Solution {

                       /**
                       * @param Integer[] $nums
                       * @return NULL
                       */
                      function moveZeroes(&$nums) {
                        $noZero = 0;
                        foreach($nums as $key => $val){
                            if($val != 0){
                                $nums[$key] = 0;
                                $nums[$noZero] = $val;
                                $noZero++;
                            }
                        }

                      }


}