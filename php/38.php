<?php

class Solution {

    /**
     * @param Integer $n
     * @return String
     */
    function countAndSay($n) {
       $result="1";
       for ($i=1;$i<$n;$i++){
          $temp="";

          for ($j=0;$j<strlen($result);$j++){
             $count=1;
             while ($j+1<strlen($result) && $result[$j]==$result[$j+1]){
              $count++;
              $j++;
             }
             $temp.=$count.$result[$j];
          }
          $result=$temp;
       }
       return $result;
    }

    // 递归
    function countAndSay_2($n) {
           $r=$s="";
           $count=1;
           if($n==1){
             return "1";
            }
            if($n==2){
             return "11";
            }
            $s=$this->countAndSay_2($n-1);
            for($i=1;$i<strlen($s);$i++){
               if($s[$i]!=$s[$i-1]){
                 $r .=$count,$s[$i-1]
                 $count=1;
               }else{
                 $count++;
               }
               // 判断是否到末尾，打印字符
               if($i+1==strlen($s)){
                  $r.=$count,$s[$i]
                  return $r;
               }
            }
            return $s;

        }

}