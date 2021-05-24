<?php

class Solution {

                       /**
                       * @param Integer[] $nums
                       * @param Integer $target
                       * @return Integer[]
                       */
                      function twoSum($nums, $target) {
                            $found = [];
                            foreach($nums as $k=>$val){
                                $diff = $target-$val;
                              if(!isset($found[$diff])){
                                 $found[$val]=$k;
                                 continue;
                              }
                             return [$found[$diff],$k];
                            }
                      }


}