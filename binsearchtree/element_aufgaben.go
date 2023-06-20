package binsearchtree

// IsLeaf returns true if the element is a leaf.
func (el *Element) IsLeaf() bool {
	return !el.IsEmpty() && el.Left.IsEmpty() && el.Right.IsEmpty()
}

