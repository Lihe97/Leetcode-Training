package main
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if root == nil{
		return node
	}
	res := root

	for root != nil{
		if root.Val > val{
			if root.Left == nil{
				root.Left = node
				break
			}else{
				root = root.Left
			}
		}else if root.Val < val{
			if root.Right == nil{
				root.Right = node
				break
			}else{
				root = root.Right
			}
		}
	}
	return res
}


func main() {
	
}
