<?php

class Solution {

//  先转为小写。
// 正则匹配，非数字和小写字母跳过

       /**
       * @param String $s
       * @return Boolean
       */
      function isPalindrome($s) {
          $s=strtolower($s);
           $l=0;
           $r=strlen($s)-1;
           while($l<$r){
              if(!$this->isValid($s[$l]) ){
                  $l++;
                 continue;
              }
              if(!$this->isValid($s[$r]) ){
                  $r--;
                 continue;
              }
              if($s[$l] !=$s[$r]){
               return false;
              }
              $l++;
              $r--;
           }
           return true;
        }

      function isValid($c){
         return  preg_match("/[\da-z]/",$c)==1;
      }


}