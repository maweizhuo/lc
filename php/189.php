<?php

class Solution {

        /**
        * @param Integer[] $nums
        * @param Integer $k
        * @return NULL
        */
       public function rotate(&$nums, $k)
          {
              $len = count($nums);
              $k%=$len;
              $nums = $this->reverse($nums, 0, $len - 1);
              $nums = $this->reverse($nums, 0, $k - 1);
              $nums = $this->reverse($nums, $k, $len - 1);
              return $nums;
          }

          private function reverse($arr, $left, $right)
          {
              while ($left <= $right) {
                  $tmp = $arr[$left];
                  $arr[$left] = $arr[$right];
                  $arr[$right] = $tmp;
                  $left++;
                  $right--;
              }

              return $arr;
          }

}