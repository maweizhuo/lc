class Solution {


    /**
     * 迭代法
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        $newHead=null;
        while($head !==null){
           $next=$head->next;
           $head->next=$newHead;
           $newHead=$head;
           $head=$next;
        }
        return $newHead;
    }

     /**
     * 递归
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList_1($head) {
        if($head==null || $head->next==null){
          return $head;
        }
        $newHead=$this->reverseList($head->next);
        $head->next->next=$head;
        $head->next=null;
        return $newHead;
    }

}
