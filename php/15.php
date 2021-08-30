<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function threeSum($nums) {
       sort($nums);
       $res=[];
       $n=count($nums);
       for($i=0;$i<$n-2;$i++){
          $n1=$nums[$i];
          if($n1>0){
            break;
          }
          if($i>0 && $n1==$nums[$i-1]){
           continue;
          }
          $l=$i+1;
          $r=$n-1;
          while($l<$r){
            $n2=$nums[$l];
            $n3=$nums[$r];
            if($n1+$n2+$n3==0){
              $res[]=[$n1,$n2,$n3];
              while($l<$r && $n2==$nums[$l]){
                $l++;
              }
              while($l<$r && $n2==$nums[$l]){
               $r--;
            }
            }elseif($n1+$n2+$n3<0){
            $l++;
            }else{
            $r--;
            }
          }
       }
       return $res;
    }
}

