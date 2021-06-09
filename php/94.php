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
     * @return Integer[]
     */
    function inorderTraversal($root) {
      $res= [];
      $this->helper($root,$res);
      return $res;

    }

    // 递归
    function helper($root,&$res) {
      if($root==null) return;
      $this->helper($root->left,$res);
      $res[]=$root->val;
      $this->helper($root->right,$res);
    }

    // 迭代
     function inorderTraversal($root) {
          $res= [];
          $stack=[];
          while($root !=null || !empty($stack)){
             while($root !=null){
                $stack[]=$root;
                $root=$root->left;
             }

             $root=array_pop($stack);
             $res[]=$root->val;
             $root=$root->right;
          }
        return $res;
     }

}