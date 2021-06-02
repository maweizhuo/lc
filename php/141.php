/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */

class Solution {
    /**
     * @param ListNode $head
     * @return Boolean
     */
    function hasCycle($head) {
    // 环形链表，快的必然追的上慢的
        if($head ==null) return false;
        $slow=$head;
        $fast=$head->next;
        while($fast !=$slow){
           if($fast ==null || $slow==null){
             return false;
           }
           $slow=$slow->next;
           $fast=$fast->next->next;
        }
        return true;
    }

     function hasCycle($head) {
        // 环形链表，快的必然追的上慢的
            if($head ==null) return false;
            $slow=$fast=$head;
            while($fast !=null && $fast->next !=null){
               $slow=$slow->next;
               $fast=$fast->next->next;
               if($fast ==$slow){
                 return true;
               }
            }
            return false;
        }

}