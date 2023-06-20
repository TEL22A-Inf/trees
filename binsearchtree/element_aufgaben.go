package binsearchtree

// IsLeaf returns true if the element is a leaf.
func (el *Element) IsLeaf() bool {
	return !el.IsEmpty() && el.Left.IsEmpty() && el.Right.IsEmpty()
}

// GetMinNode gibt das Element mit dem kleinsten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (t *Element) GetMinNode() *Element {
	if t.IsEmpty() {
		return nil
	}
	current := t
	for !current.Left.IsEmpty() {
		current = current.Left
	}
	return current
}

// GetMaxNode gibt das Element mit dem größten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (t *Element) GetMaxNode() *Element {
	if t.IsEmpty() {
		return nil
	}
	current := t
	for !current.Right.IsEmpty() {
		current = current.Right
	}
	return current
}
