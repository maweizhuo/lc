<?php

class Solution {

           /**
           * @param Integer[] $nums
           * @return Boolean
           */
          function containsDuplicate($nums) {
             $num_arr = [];
             for($i=0;$i<count($nums);$i++){
                        if(isset($num_arr[$nums[$i]])){
                           return true;
                        }
                        $num_arr[$nums[$i]]=1;
             }
             return false;
          }

}