<?php

class Solution {

              /**
               * @param Integer[] $nums
               * @return Integer
               */
              function singleNumber($nums) {
                $single=0;
                foreach($nums as $num){
                  $single^=$num;
                 }
                 return $single;
              }

}