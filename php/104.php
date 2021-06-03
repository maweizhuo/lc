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
     * dfs
     * @param TreeNode $root
     * @return Integer
     */
    function maxDepth($root) {
     if($root==null) return 0;
     $left=$this->maxDepth($root->left);
     $right=$this->maxDepth($root->right);
     return max($left,$right)+1;
    }

     // bfs
    function maxDepth($root) {
      if($root==null) return 0;
      $queue=[$root];
      $depth=0;
      while(count($queue)>0){
        $size=count($queue);
        for($i=0;$i<$size;$i++){
            $s=array_shift($queue);
            if($s->left !=null){
               array_push($queue,$s->left);
            }
            if($s->right !=null){
               array_push($queue,$s->right);
            }
        }
        $depth++;
      }
      return $depth;
    }
}