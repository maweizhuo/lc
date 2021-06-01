<?php

class Solution {

     /**
      * Definition for a singly-linked list.
      * class ListNode {
      *     public $val = 0;
      *     public $next = null;
      *     function __construct($val) { $this->val = $val; }
      * }
      */
     function deleteNode($node) {
         $node->val=$node->next->val;
         $node->next=$node->next->next;
     }

}