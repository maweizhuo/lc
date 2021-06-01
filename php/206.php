/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
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
    function reverseList($head) {
        if($head==null || $head->next==null){
          return $head;
        }
        $newHead=$this->reverseList($head->next);
        $head->next->next=$head;
        $head->next=null;
        return $newHead;
    }

}