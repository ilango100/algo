package data

/*A BinNode represents a single node in Binary Tree.*/
type BinNode struct {
	Key   int
	Left  *BinNode
	Right *BinNode
	Par   *BinNode
}

/*BinTree represents  Binary Tree. It points to the root *BinNode.*/
type BinTree *BinNode

/*Search searches through the Binary Tree and returns the node.*/
func (b *BinNode) Search(k int) *BinNode {
	if b.Key == k {
		return b
	}
	if k < b.Key {
		if b.Left == nil {
			return nil
		}
		return b.Left.Search(k)
	} else if b.Right != nil {
		return b.Right.Search(k)
	}
	return nil
}

/*SearchIt searches through the Binary Tree iteratively, instead of recursively.*/
func (b *BinNode) SearchIt(k int) *BinNode {
	var x *BinNode
	for x = b; x != nil && x.Key != k; {
		if k < x.Key {
			x = x.Left
			continue
		}
		x = x.Right
	}
	return x
}
