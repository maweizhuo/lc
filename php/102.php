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

   private $res=[];
    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    function levelOrder($root) {
        // 递归
        $this->dfs($root,0);
        return  $this->res;
    }

    function dfs($root,$level){
        if($root==null){
            return $this->res;
        }
        $this->res[$level][]=$root->val;
        $this->dfs($root->left,$level+1);
        $this->dfs($root->right,$level+1);

    }

      // 迭代
      function levelOrder($root) {
         $result=[];
         if($root==null){
           return $result;
         }
         $queue=[$root];

         for($level=0;0<count($queue);$level++){
            $next=[];
            for($j=0;$j<count($queue);$j++){
               $node=$queue[$j];
               $result[$level][]=$node->val;
               if($node->left !=null){
                 $next[]=$node->left;
               }
               if($node->right !=null){
                $next[]=$node->right;
               }
            }
            $queue=$next;
         }
         return $result;
      }

}