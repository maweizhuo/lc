class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function reverseBits($n) {
        $rev=0;
        for($i=0;$i<32;$i++){
           $rev= ($rev <<1)+($n&1);
           $n>>=1;
        }
        return $rev;
    }
}