class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
       $n=strlen($s);
       $pairs=[
        '}'=>'{',
        ')'=>'(',
        ']'=>'[',
       ];
       $stack=[];
       for($i=0;$i<$n;$i++){
          if(isset($pairs[$s[$i]])){
            if(count($stack) ==0 || $pairs[$s[$i]] != $stack[count($stack)-1]){
              return false;
            }
            array_pop($stack);
          }else{
            array_push($stack,$s[$i]);
          }
       }
       return $stack ? false : true;
    }
}