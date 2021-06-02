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
     * @param ListNode $head
     * @return Boolean
     */
    function isPalindrome($head) {
        $vals=[];
        while($head !=null){
           $vals[]=$head->val;
           $head=$head->next;
        }
        $num=count($vals);
        for($i=0;$i<$num/2;$i++){
           if($vals[$i] !=$vals[$num-$i-1]){
               return false;
           }
        }
        return true;
    }

     /**
     * @param ListNode $head
     * @return Boolean
     */
    function isPalindrome($head) {
       if($head ==null){
         return true;
       }
       $slow=$fast=$head;
       while($fast !=null && $fast->next !=null){
          $slow=$slow->next;
          $fast=$fast->next->next;
       }
       if($fast !=null){
          $slow=$slow->next; // 奇数取下个值
       }
       $pre=null;
       while($slow !=null){ // 反转后半段
          $next=$slow->next;
          $slow->next=$pre;
          $pre=$slow;
          $slow=$next;
       }

       while($pre !=null){
          if($pre->val !=$head->val){
             return false;
          }
          $pre=$pre->next;
          $head=$head->next;
       }
       return true;

    }


}