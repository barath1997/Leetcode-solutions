/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
}