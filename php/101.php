/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    function isSymmetric($root) {
      return $this->check($root,$root);
    }

    function check($p,$q){
       if($p== null && $q==null){
         return true;
       }
       if($p==null || $q==null){
         return false;
       }
       return $p->val ==$q->val && $this->check($p->left,$q->right) && $this->check($p->right,$q->left);
    }

}