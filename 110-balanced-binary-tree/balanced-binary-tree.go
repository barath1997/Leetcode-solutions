/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// discussed in class on 19-jan-2025 
var result bool = true
func isBalanced(root *TreeNode) bool {
    
    result = true
    maxHeight(root)
    return result
    
}

func maxHeight(root *TreeNode) int {

   if result == false { return 0 }
   if root == nil { return 0 }

   left_height := maxHeight(root.Left)
   right_height := maxHeight(root.Right)

   if math.Abs(float64(left_height - right_height)) > 1 { result = false }

   return 1 + max(left_height,right_height) 

}