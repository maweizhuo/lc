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
    function isValidBST($root) {
       return $this->validBst($root,pow(-2,63),pow(2,63)-1);
    }


    function validBst($root,$min,$max){
      if($root==null){
        return true;
      }
      if($root->val<=$min || $root->val>=$max){
        return false;
      }
      return $this->validBst($root->left,$min,$root->val) && $this->validBst($root->right,$root->val,$max);
    }

}