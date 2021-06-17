class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function countPrimes($n) {
       $cnt=0;
       $isPrime=[];
       for($i=0;$i<$n;$i++){
         $isPrime[$i]=true;
       }
       for($i=2;$i<$n;$i++){
          if($isPrime[$i]){
             $cnt++;
             for($j=2*$i;$j<$n;$j+=$i){
               $isPrime[$j]=false;
             }
          }
       }
       return $cnt;
    }

    function countPrimes($n) {
           $isPrime=$primes=[];
           for($i=0;$i<$n;$i++){
             $isPrime[$i]=true;
           }
           for($i=2;$i<$n;$i++){
              if($isPrime[$i]){
                $primes[]=$i;
              }
              foreach($primes as $p){
                if($i*$p>=$n)break;
                $isPrime[$i*$p]=false;
                if($i%$p==0)break;
              }
           }
           return count($primes);
        }

}