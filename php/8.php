<?php

class Solution {

         /**
          * @param String $s
          * @return Integer
          */
         function myAtoi($s) {
           $s=trim($s);
           $result=0;
           $sign=1;
           for ($i=0;$i<strlen($s);$i++){
              if(is_numeric($s[$i])){
                 $result=$result*10+$s[$i];
              }elseif($i==0 && $s[$i]=='-'){
                 $sign=-1;
              }elseif($i==0 && $s[$i]=='+'){
                 $sign=1;
              }else{
                 break;
              }
              $maxInt32=pow(2,31)-1;
              if($result>$maxInt32){
                 if($sign==-1){
                   return pow(-2,31);
                 }
                 return $maxInt32;
              }

           }
           return $sign*$result;
         }

         // 正则匹配
         function myAtoi_2($s) {
            preg_match("/^[\+\-]?\d+/",trim($s),$match);
            if(!count($match)) return 0;
            $intval=(Integer)$match[0];
            if($intval>=pow(2,31)-1)return pow(2,31)-1;
            if($intval<=pow(-2,31))return pow(-2,31);
            return $intval;
         }

}