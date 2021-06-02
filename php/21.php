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
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function mergeTwoLists($l1, $l2) {
       $dummyHead=new ListNode(); // 虚拟一个头结点
       $res=$dummyHead;
       while($l1|| $l2){
          // $l1不为空。并且$l2为空 或者 $l1不为空并且值小于$l2
          // 这里不能只用：$l1->val <=$l2->val,如果$l1为空，结果也是true
          if(($l1 && $l2==null) || ($l1 && $l1->val<=$2->val)){
             $res->next=$l1;
             $res=$l1;
             $l1=$l1->next;
          }else{
            $res->next=$l2;
            $res=$l2;
            $l2=$l2->next;
          }
       }
       return $dummyHead->next;
    }
}