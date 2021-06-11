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
     * @return Integer
     */
    function rob($root) {
       return max($this->dfs($root));
    }

    function dfs($root){
       if($root ==null){
         return [0,0];
       }
       $left_arr = $this->dfs($root->left);
       $right_arr=$this->dfs($root->right);
       $isRoot=$root->val+$left_arr[1]+$right_arr[1];
       $notRoot=max($left_arr)+max($right_arr);
       return [$isRoot,$notRoot];
    }

}