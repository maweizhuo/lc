<?php

class Solution {

            /**
            * @param String[][] $board
            * @return Boolean
            */
           function isValidSudoku($board) {
                $heng=$zong=$fang=[];
                for($i=0;$i<9;$i++){
                   for($j=0;$j<9;$j++){
                        if($board[$i][$j] =='.'){
                          continue;
                        }
                       // 横判断
                       if(isset($heng[$i][$board[$i][$j]])){
                          return false;
                       }else{
                          $heng[$i][$board[$i][$j]]=1;
                       }

                       // 纵判断
                       if(isset($zong[$j][$board[$i][$j]])){
                         return false;
                       }else{
                          $zong[$j][$board[$i][$j]]=1;
                       }

                       // 方块判断
                       $index=floor($i/3)*3+floor($j/3);
                       if(isset($fang[$index][$board[$i][$j]])){
                          return false;
                       }else{
                          $fang[$index][$board[$i][$j]]=1;
                       }
                   }
                }
              return true;
           }

}