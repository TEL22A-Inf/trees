package binsearchtree

type Tree struct {
	Root *Element
}

// NewTree creates a new empty tree.
func NewTree() *Tree {
	return &Tree{Root: NewElement()}
}

// IsEmpty returns true if the tree is empty.
func (t *Tree) IsEmpty() bool {
	return t.Root.IsEmpty()
}

// Insert inserts a new element into the tree.
// Returns a pointer to the new element.
func (t *Tree) Insert(value int) *Element {
	return t.Root.Insert(value)
}
