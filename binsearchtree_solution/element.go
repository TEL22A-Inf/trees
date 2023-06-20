package binsearchtree

// Element is a node in a binary search tree.
type Element struct {
	Value int
	Left  *Element
	Right *Element
}

// NewElement creates a new empty element.
func NewElement() *Element {
	return &Element{}
}

// IsEmpty returns true if the element is empty.
func (e *Element) IsEmpty() bool {
	return e.Left == nil && e.Right == nil
}

// SetValue sets the value of the element.
// If the element is empty, two new children are created.
func (e *Element) SetValue(value int) {
	if e.IsEmpty() {
		e.Left = NewElement()
		e.Right = NewElement()
	}
	e.Value = value
}

// Insert inserts a new element into the tree.
// Returns a pointer to the new element.
func (e *Element) Insert(value int) *Element {
	if e.IsEmpty() {
		e.SetValue(value)
		return e
	}
	if value < e.Value {
		return e.Left.Insert(value)
	}
	return e.Right.Insert(value)
}

// GetNode returns the node with the given value.
// If the node is not found, nil is returned.
func (e *Element) GetNode(value int) *Element {
	if e.IsEmpty() {
		return nil
	}
	if e.Value == value {
		return e
	}
	if value < e.Value {
		return e.Left.GetNode(value)
	}
	return e.Right.GetNode(value)
}

// InOrder returns the elements of the tree in order.
func (e *Element) InOrder() []int {
	if e.IsEmpty() {
		return []int{}
	}
	return append(append(e.Left.InOrder(), e.Value), e.Right.InOrder()...)
}
