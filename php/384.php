class Solution {
    protected $arr;
    protected $copy;
    /**
     * @param Integer[] $nums
     */
    function __construct($nums) {
      $this->arr=$nums;
      $this->copy=$nums;
    }

    /**
     * Resets the array to its original configuration and return it.
     * @return Integer[]
     */
    function reset() {
        $this->arr=$this->copy;
        return $this->arr;
    }

    /**
     * Returns a random shuffling of the array.
     * @return Integer[]
     */
    function shuffle() {
       $len_a =count($this->arr);
       for($i=0;$i<$len_a;$i++){
          $rand=rand($i,$len_a-1);
          $t=$this->arr[$i];
          $this->arr[$i]=$this->arr[$rand];
          $this->arr[$rand]=$t;
       }
       return $this->arr;
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * $obj = Solution($nums);
 * $ret_1 = $obj->reset();
 * $ret_2 = $obj->shuffle();
 */