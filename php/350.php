<?php

class Solution {

                 /**
                * @param Integer[] $nums1
                * @param Integer[] $nums2
                * @return Integer[]
                */
               function intersect($nums1, $nums2) {
                  $set = $arr = [];
                  foreach($nums1 as $num1){
                     $set[$num1]+=1;
                   }
                   foreach($nums2 as $num2){
                      if($set[$num2] >0){
                         $set[$num2]-=1;
                         $arr[]=$num2;
                      }
                   }
                   return $arr;
               }

                function intersect_2($nums1, $nums2) {
                          $i = $j = $k = 0;
                          sort($nums1);
                          sort($nums2);
                          $num1_len = count($nums1);
                          $num2_len = count($nums2);
                          $arr = [];
                          while ($i < $num1_len && $j < $num2_len) {
                              if ($nums1[$i] == $nums2[$j]) {
                                  $arr[] = $nums1[$i];
                                  $i++;
                                  $j++;
                              } else {
                                  $nums1[$i] > $nums2[$j] ? $j++ : $i++;
                              }
                          }
                          return $arr;
                 }

}