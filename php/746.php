class Solution {

    /**
     * @param Integer[] $cost
     * @return Integer
     */
    function minCostClimbingStairs($cost) {
        foreach($cost as $val){
           $next=min($prev+$val,$curr+$val);
           $prev=$curr;
           $curr=$next;
        }
        return min($prev,$curr);
    }

     function minCostClimbingStairs($cost) {
        $len=count($cost);
        $dp=[];
        $dp[0]=$cost[0];
        $dp[1]=$cost[1];
        for($i=2;$i<=$len;$i++){
           $dp[$i]=min($dp[$i-1]+$cost[$i],$dp[$i-2]+$cost[$i]);
        }
        return $dp[$len];
     }

}