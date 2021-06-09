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
    function postorderTraversal($root) {
      $res=[];
      $this->helper($root,$res);
      return $res;
    }

// 递归
    function helper($root,&$res){
      if($root==null)return;
      $this->helper($root->left,$res);
      $this->helper($root->right,$res);
      $res[]=$root->val;
    }

    // 迭代
     function postorderTraversal($root) {
          $res=[];
          $stack=[];
          $prev=null;
          while($root !=null || !empty($stack)){
             while($root !=null){
                $stack[]=$root;
                $root=$root->left;
             }
             $root=array_pop($stack);
             if($root->right ==null || $root->right ==$prev){
                $res[]=$root->val;
                $prev=$root;
                $root=null;
             }else{
                $stack[]=$root;
                $root=$root->right;
             }
          }
          return $res;
     }

}