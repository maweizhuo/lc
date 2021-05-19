<?php

class Solution {

        /**
        * @param Integer[] $prices
        * @return Integer
        */
       function maxProfit($prices) {
           // 贪心算法  ，后一天收益比前一天收益大，就锁定利润
           $profit =0;
           for($i=0;$i<count($prices)-1;$i++){
               if($prices[$i]<$prices[$i+1]){
                 $profit +=$prices[$i+1]-$prices[$i];
               }
           }
           return $profit;
       }

}