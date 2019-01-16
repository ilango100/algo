package data

/*A BinTree represents a Binary Tree.*/
type BinTree struct {
	Key   int
	Value interface{}
	Left  *BinTree
	Right *BinTree
	Par   *BinTree
}

/*Search searches through the Binary Tree and returns the node.*/
func (b *BinTree) Search(k int) *BinTree {
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
func (b *BinTree) SearchIt(k int) *BinTree {
	for b != nil && b.Key != k {
		if k < b.Key {
			b = b.Left
			continue
		}
		b = b.Right
	}
	return b
}
