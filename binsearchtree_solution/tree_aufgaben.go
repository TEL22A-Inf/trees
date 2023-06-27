package binsearchtree

/*
Diese Datei enthält die Aufgaben für den Binärbaum.
Es geht darum, weitere Methoden für den Binärbaum zu implementieren.
*/

// RemoveValue entfernt ein Element mit dem Wert value aus dem Baum.
// Wenn das Element nicht gefunden wird, passiert nichts.
func (t *Tree) RemoveValue(value int) {
	el := t.Root.GetNode(value)
	if el == nil {
		return
	}
	el.RemoveRoot()
}

// NodeCount gibt die Anzahl der Elemente im Baum zurück.
func (t *Tree) NodeCount() int {
	return t.Root.NodeCount()
}

// Hilfsfunktion, um NodeCount rekursiv in Element umsetzen zu können.
func (e *Element) NodeCount() int {
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.Left.NodeCount() + e.Right.NodeCount()
}

// IsSearchTree gibt true zurück, wenn der Baum ein Suchbaum ist.
func (t *Tree) IsSearchTree() bool {
	return t.Root.IsSearchTree()
}

// Hilfsfunktion, um IsSearchTree rekursiv in Element umsetzen zu können.
func (e *Element) IsSearchTree() bool {
	if e.IsEmpty() {
		return true
	}
	if !e.Left.IsEmpty() && e.Left.Value > e.Value {
		return false
	}
	if !e.Right.IsEmpty() && e.Right.Value < e.Value {
		return false
	}
	return e.Left.IsSearchTree() && e.Right.IsSearchTree()
}

/** ACHTUNG: Ab hier sind die Funktionen alle in Element definiert.
 * Das ist eigentlich ein Fehler, sie sollten besser in element.go stehen.
 * Aber da ich die Aufgaben schon mehrfach angepasst habe und weitere Verwirrung
 * vermeiden möchte, lasse ich es jetzt so.
**/

// Height gibt die Höhe des Baums zurück.
func (t *Element) Height() int {
	if t.IsEmpty() {
		return 0
	}
	return 1 + max(t.Left.Height(), t.Right.Height())
}

// max gibt das Maximum von a und b zurück.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// BalanceFactor gibt den Balancefaktor des Knotens zurück.
func (e *Element) BalanceFactor() int {
	if e.IsEmpty() {
		return 0
	}
	return e.Right.Height() - e.Left.Height()

	// Anmerkung: Diese Funktion ist zwar korrekt, aber ineffizient.
	// Sie ruft Height auf, das den Baum rekursiv durchläuft.
	// D.h. sie braucht in jedem Fall O(n) Zeit.
	// Besser wäre es, die Höhe bei der Einfügeoperation zu berechnen
	// und in jedem Knoten zu speichern.
}

// RotateLeft rotiert den Baum nach links.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeft(root *Element) *Element {
	// Spezialfall: Leerer Baum. Hier kann nichts rotiert werden.
	if root.IsEmpty() {
		return root
	}

	// Ausgangssituation:
	//          A
	//         / \
	//        B   C
	//           / \
	//          D   E
	A := root
	C := root.Right
	D := C.Left
	// (B und E müssen wir uns nicht merken, ihre Position ändert sich nicht)

	// Ergebnis nach der Rotation:
	//          C
	//         / \
	//        A   E
	//       / \
	//      B   D
	C.Left = A
	A.Right = D

	return C
}

// RotateRight rotiert den Baum nach rechts.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRight(root *Element) *Element {

	// Spezialfall: Leerer Baum. Hier kann nichts rotiert werden.
	if root.IsEmpty() {
		return root
	}

	// Ausgangssituation:
	//          A
	//         / \
	//        B   C
	//       / \
	//      D   E
	A := root
	B := root.Left
	E := B.Right
	// (D und C müssen wir uns nicht merken, ihre Position ändert sich nicht)

	// Ergebnis nach der Rotation:
	//          B
	//         / \
	//        D   A
	//           / \
	//          E   C
	B.Right = A
	A.Left = E

	return B
}

// RotateLeftRight führt eine Doppelrotation nach links aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeftRight(root *Element) *Element {
	root.Left = RotateLeft(root.Left)
	return RotateRight(root)
}

// RotateRightLeft führt eine Doppelrotation nach rechts aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRightLeft(root *Element) *Element {
	root.Right = RotateRight(root.Right)
	return RotateLeft(root)
}
