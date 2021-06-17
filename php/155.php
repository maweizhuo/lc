class MinStack {

    protected $stack=[];
    protected $minStack=[];
    /**
     * initialize your data structure here.
     */
    function __construct() {

    }

    /**
     * @param Integer $val
     * @return NULL
     */
    function push($val) {
     array_push($this->stack,$val);
     if(empty($this->minStack) || end($this->minStack)>=$val){
       array_push($this->minStack,$val);
     }

    }

    /**
     * @return NULL
     */
    function pop() {
        $tmp=array_pop($this->stack);
        if($tmp==end($this->minStack)){
        array_pop($this->minStack);
        }
    }

    /**
     * @return Integer
     */
    function top() {
        return $this->stack[count($this->stack)-1];
    }

    /**
     * @return Integer
     */
    function getMin() {
       return $this->minStack[count($this->minStack)-1];
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * $obj = MinStack();
 * $obj->push($val);
 * $obj->pop();
 * $ret_3 = $obj->top();
 * $ret_4 = $obj->getMin();
 */