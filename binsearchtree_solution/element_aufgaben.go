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
	if t.Left.IsEmpty() {
		return t
	}
	return t.Left.GetMinNode()
}

// GetMaxNode gibt das Element mit dem größten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (t *Element) GetMaxNode() *Element {
	if t.IsEmpty() {
		return nil
	}
	if t.Right.IsEmpty() {
		return t
	}
	return t.Right.GetMaxNode()
}

// RemoveRoot removes the root element from the tree.
// If the root element is a leaf, the tree becomes empty.
func (el *Element) RemoveRoot() {
	// Anmerkung: Die Methode entfernt genau das Element, auf dem sie aufgerufen wird.

	// Wenn das Element leer ist, wird die Methode beendet.
	if el.IsEmpty() {
		return
	}

	// Wenn das Element ein Blatt ist, wird es "geleert" und die Methode beendet.
	if el.IsLeaf() {
		el.Left = nil
		el.Right = nil
		return
	}

	// Bestimme das größte Element im linken und das kleinste Element im rechten Teilbaum.
	// Diese Elemente werden benötigt, um das Element zu ersetzen, das entfernt werden soll.
	// Beachte: Falls der linke oder rechte Teilbaum leer ist, ist das jeweilige Element nil.
	leftMax := el.Left.GetMaxNode()
	rightMin := el.Right.GetMinNode()

	// Wenn leftMax existiert, ersetze das zu entfernende Element durch leftMax.
	// Anschließend entferne leftMax aus seinem Teilbaum.
	if leftMax != nil {
		el.SetValue(leftMax.Value)
		leftMax.RemoveRoot()
		return
	}

	// leftMax nicht existiert, muss rightMin existieren.
	// Ersetze das zu entfernende Element durch rightMin.
	// Anschließend entferne rightMin aus seinem Teilbaum.
	el.SetValue(rightMin.Value)
	rightMin.RemoveRoot()
}
