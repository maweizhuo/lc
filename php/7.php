<?php

class Solution {

         /**
         * @param Integer $x
         * @return Integer
         */
        function reverse($x) {
           $res = 0;
           while($x !=0){
              $res=$res*10+$x%10;
              $x=intval($x/10);
           }
           if($res>pow(2,31)-1 || $res<pow(-2,31)){
             return 0;
           }
           return $res;
        }

}