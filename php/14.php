<?php

class Solution {

        /**
        * @param String[] $strs
        * @return String
        */
       function longestCommonPrefix($strs) {
           if(count($strs)==0){
             return "";
           }
          $prefix=$strs[0];
          foreach($strs as $str){
              while($prefix!='' && strpos($str,$prefix) !==0){
                  if(strlen($prefix)==0){
                   return "";
                 }
                 $prefix=substr($prefix,0,strlen($prefix)-1);
              }
          }
           return $prefix;
       }



}