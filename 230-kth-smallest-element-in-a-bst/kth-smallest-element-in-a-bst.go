/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    
    arr := []int{}
    inorderRes(root,&arr)
    for idx,v := range arr {
		if idx + 1 == k { return v }
	}
    
	return 0
}

func inorderRes(node *TreeNode,arr *[]int){

      if node == nil { return }
	  inorderRes(node.Left,arr)
	  *arr = append(*arr, node.Val)
	  inorderRes(node.Right,arr)
}