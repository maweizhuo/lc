<?php

class Solution {

          /**
           * @param String $s
           * @return Boolean
           */
          function isPalindrome($s) {
             $s=strtolower($s);
             for($l=0,$r=strlen($s)-1;$l<$r;$l++,$r--){
                if(!$this->isValid($s[$l]) || !$this->isValid($s[$r])){
                   continue;
                }
                if($s[$l] !=$s[$r]){
                 return false;
                }
             }
             return true;
          }

          function isValid($c){
            return preg_match('/[\da-z]/',$c);
          }


}