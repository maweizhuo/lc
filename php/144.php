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
    function preorderTraversal($root) {
      $res=[];
      $this->helper($root,$res);
      return $res;
    }

    // 递归
    function helper($root,&$res){
       if($root==null)return;
       $res[]=$root->val;
       $this->helper($root->left,$res);
       $this->helper($root->right,$res);
    }

    // 迭代
      function preorderTraversal($root) {
          $res=[];
          $stack=[];
          while($root !=null || !empty($stack)){
            while($root !=null){
              $res[]=$root->val;
              $stack[]=$root;
              $root=$root->left;
            }
            $node=array_pop($stack);
            $root=$node->right;
          }
          return $res;
      }

}