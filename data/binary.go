package data

/*A BinTree represents a Binary Tree.*/
type BinTree struct {
	Key    int
	Value  interface{}
	Left   *BinTree
	Right  *BinTree
	Parent *BinTree
}

func (b *BinTree) successor() *BinTree {
	if b.Right != nil {
		return b.Right.Min()
	}
	if b.Parent == nil {
		return nil
	}
	par := b.Parent
	for par != nil && par.Right != nil && par.Right.Key == b.Key {
		b = par
		par = par.Parent
	}
	return par
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

/*Min returns the minimum value from the Binary Tree.*/
func (b *BinTree) Min() *BinTree {
	for b.Left != nil {
		b = b.Left
	}
	return b
}

/*Max returns the maximum value from the Binary Tree.*/
func (b *BinTree) Max() *BinTree {
	for b.Right != nil {
		b = b.Right
	}
	return b
}

/*Insert inserts a key, value pair into the Binary Tree and returns the respective Node.*/
func (b *BinTree) Insert(key int, value interface{}) *BinTree {
	if key < b.Key {
		if b.Left == nil {
			b.Left = &BinTree{
				Key:    key,
				Value:  value,
				Parent: b,
			}
			return b.Left
		}
		return b.Left.Insert(key, value)
	}
	if b.Right == nil {
		b.Right = &BinTree{
			Key:    key,
			Value:  value,
			Parent: b,
		}
		return b.Right
	}
	return b.Right.Insert(key, value)
}

/*Delete deletes a key, with its node from the Binary Tree.
Yet to be implemented*/
func (b *BinTree) Delete(key int) {
	//TODO: Implement delete function
}
