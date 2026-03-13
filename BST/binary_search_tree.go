package bst

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	return root
}

func Search(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return Search(root.Left, val)
	}
	return Search(root.Right, val)
}

func findMin(root *TreeNode) *TreeNode{
    for root.Left != nil{
        root = root.Left
    }
    return root
}

func Delete(root *TreeNode, val int) *TreeNode {
	if root == nil{
        return nil
    }
    
    if val < root.Val{
        root.Left = Delete(root.Left, val)
    }else if val > root.Val{
        root.Right = Delete(root.Right, val)
    }else{
        if root.Left == nil{
            return root.Right
        }
        if root.Right == nil{
            return root.Left
        }
        
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = Delete(root.Right, minNode.Val)
    }
    return root
}

func Inorder(root *TreeNode) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.Val, " ")
	Inorder(root.Right)
}

func Preorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}

func Postorder(root *TreeNode) {
	if root == nil {
		return
	}

	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.Val, " ")
}
