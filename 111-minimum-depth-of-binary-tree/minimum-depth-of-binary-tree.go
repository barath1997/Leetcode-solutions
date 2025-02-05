/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// discussed solution in class
func minDepth(root *TreeNode) int {
   
   if root == nil {
    return 0
   }

   left := minDepth(root.Left)
   right := minDepth(root.Right)

   // skewed tree condition check
   if min(left,right) == 0 {
    return 1 + max(left,right)
   }
   
   // for normal tree
   return 1 + min(left,right)
}

// own solution 
/*
func minDepth(root *TreeNode) int {
    
    if root == nil {
        return 0
    }

    if root.Left == nil {
        rh := minDepth(root.Right)
        return rh + 1
    } else if root.Right == nil {
        lh := minDepth(root.Left)
        return lh + 1
    }

    lh := minDepth(root.Left)
    rh := minDepth(root.Right)
    
    return 1 + min(lh,rh)
}*/