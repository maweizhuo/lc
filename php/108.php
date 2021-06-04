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
     * @param Integer[] $nums
     * @return TreeNode
     */
    function sortedArrayToBST($nums) {
       if(empty($nums)){
          return null;
       }
       $mid=floor(count($nums)/2);
       $left = array_slice($nums,0,$mid);
       $right = array_slice($nums,$mid+1,$mid);
       return new TreeNode($nums[$mid],$this->sortedArrayToBST($left),$this->sortedArrayToBST($right));
    }
}