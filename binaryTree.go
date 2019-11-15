package binary_tree2

import (
	"fmt"
)

type NotFound struct {
	x int
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("element %d not found", n.x)
}

type BTree struct {
	Element int
	Left    *BTree
	Right   *BTree
}

func NewBtree(x int) *BTree {
	b := new(BTree)
	b.Element = x
	return b
}

func Find(x int, b *BTree) *BTree {
	if x == b.Element {
		return b
	} else if x < b.Element {
		if b.Left != nil {
			return Find(x, b.Left)
		} else {
			return nil
		}
	} else {
		if b.Right != nil {
			return Find(x, b.Right)
		} else {
			return nil
		}
	}
}

func FindMin(b *BTree) *BTree {
	if b.Left != nil {
		return FindMin(b.Left)
	} else {
		return b
	}
}

func FindMax(b *BTree) *BTree {
	if b.Right != nil {
		return FindMax(b.Right)
	} else {
		return b
	}
}

func Insert(x int, b *BTree) {
	if x == b.Element {
		return
	} else if x < b.Element {
		if b.Left != nil {
			Insert(x, b.Left)
		} else {
			b.Left = NewBtree(x)
		}
	} else {
		if b.Right != nil {
			Insert(x, b.Right)
		} else {
			b.Right = NewBtree(x)
		}
	}
}

func PopMin(b *BTree) (int, *BTree) {
	x := 0
	if b.Left != nil {
		x, b.Left = PopMin(b.Left)
		return x, b.Left
	} else {
		return b.Element, nil
	}
}

func Delete(x int, b *BTree) (*BTree, error) {
	var err error
	if x == b.Element {
		if b.Left == nil && b.Right == nil {
			b = nil
		} else if b.Left != nil && b.Right != nil {
			b.Element, b.Right = PopMin(b.Right)
		} else {
			if b.Left == nil {
				b = b.Right
			} else {
				b = b.Left
			}
		}
		return b, nil
	} else if x < b.Element {
		if b.Left != nil {
			b.Left, err = Delete(x, b.Left)
			return b, err
		} else {
			return b, &NotFound{x: x}
		}
	} else {
		if b.Right != nil {
			b.Right, err = Delete(x, b.Right)
			return b, err
		} else {
			return b, &NotFound{x: x}
		}
	}
}
