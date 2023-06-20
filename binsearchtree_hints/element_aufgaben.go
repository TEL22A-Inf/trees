package binsearchtree

// IsLeaf returns true if the element is a leaf.
func (el *Element) IsLeaf() bool {
	// TODO
	return false
	// Prüfen Sie, ob das Blatt nicht leer ist und ob es keine Kinder hat.
}

// GetMinNode gibt das Element mit dem kleinsten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (t *Element) GetMinNode() *Element {
	// TODO
	return nil
	// Prüfen Sie, ob das Element leer ist.
	// Prüfen Sie, ib der linke Teilbaum leer ist.
	// Wenn nicht, rufen Sie die Methode rekursiv auf.
}

// GetMaxNode gibt das Element mit dem größten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (t *Element) GetMaxNode() *Element {
	// TODO
	return nil
	// Analog zu GetMinNode.
}

// RemoveRoot removes the root element from the tree.
// If the root element is a leaf, the tree becomes empty.
func (el *Element) RemoveRoot() {
	// Anmerkung: Die Methode entfernt genau das Element, auf dem sie aufgerufen wird.

	// TODO

	// Wenn das Element leer ist, wird die Methode beendet.

	// Wenn das Element ein Blatt ist, wird es "geleert" und die Methode beendet.

	// Bestimme das größte Element im linken und das kleinste Element im rechten Teilbaum.
	// Diese Elemente werden benötigt, um das Element zu ersetzen, das entfernt werden soll.
	// Beachte: Falls der linke oder rechte Teilbaum leer ist, ist das jeweilige Element nil.

	// Wenn leftMax existiert, ersetze das zu entfernende Element durch leftMax.
	// Anschließend entferne leftMax aus seinem Teilbaum.

	// leftMax nicht existiert, muss rightMin existieren.
	// Ersetze das zu entfernende Element durch rightMin.
	// Anschließend entferne rightMin aus seinem Teilbaum.
}
