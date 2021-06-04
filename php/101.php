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

    // 迭代方式
     function isSymmetric($root) {
           if($root ==null){
             return true;
           }
           return $this->isMirror($root->left,$root->right);
      }

      function isMirror($left,$right){
        $q=[$left,$right];
        while(!empty($q)){
           $n1=array_shift($q);
           $n2=array_shift($q);
           if($n1==null &&$n2==null){
             continue;
           }
            if($n1==null || $n2==null){
                return false;
            }
            if($n1->val!=$n2->val){
               return false;
            }
           array_push($q,$n1->left,$n2->right);
           array_push($q,$n1->right,$n2->left);
        }
        return true;
      }

}