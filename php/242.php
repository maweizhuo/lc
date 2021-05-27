<?php

class Solution {

      /**
           * @param String $s
           * @param String $t
           * @return Boolean
           */
          function isAnagram($s, $t) {
             if(strlen($s) !=strlen($t)){
               return false;
             }
             $hash =[];
             for($i=0;$i<strlen($s);$i++){
               $hash[$s[$i]]+=1;
               $hash[$t[$i]]-=1;
             }
              foreach($hash as $v){
                 if($v!=0){
                  return false;
                 }
              }
              return true;
          }

}