<?php

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
        * @param Integer $n
        * @return ListNode
        */
       function removeNthFromEnd($head, $n) {
          $dummy=new ListNode();
          $dummy->next=$head;
          $slow=$fast=$dummy; // 快慢指针
          for($i=0;$i<$n+1;$i++){
             $fast=$fast->next;
          }
          while($fast !=null){
            $fast=$fast->next;
            $slow=$slow->next;
          }
          $slow->next=$slow->next->next;
          return $dummy->next;

       }

       /**
       * @param ListNode $head
       * @param Integer $n
       * @return ListNode
       */
      function removeNthFromEnd($head, $n) {
         $dummy=new ListNode();
         $dummy->next=$head;
         $slow=$fast=$dummy; // 快慢指针

         while($fast !=null){
           $fast=$fast->next;
           $i++;
           if($i>$n+1){
              $slow=$slow->next;
           }
         }
         $slow->next=$slow->next->next;
         return $dummy->next;

      }

   }

