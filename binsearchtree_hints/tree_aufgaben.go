package binsearchtree

/*
Diese Datei enthält die Aufgaben für den Binärbaum.
Es geht darum, weitere Methoden für den Binärbaum zu implementieren.
*/

// RemoveValue entfernt ein Element mit dem Wert value aus dem Baum.
// Wenn das Element nicht gefunden wird, passiert nichts.
func (t *Tree) RemoveValue(value int) {
	// Knoten mit Wert value finden.
	// Diesen Knoten aus seinem Teilbaum löschen, falls er nicht nil ist.
	// TODO
}

// NodeCount gibt die Anzahl der Elemente im Baum zurück.
func (t *Tree) NodeCount() int {
	// TODO
	return 0
}

// IsSearchTree gibt true zurück, wenn der Baum ein Suchbaum ist.
func (t *Tree) IsSearchTree() bool {

	// Ein Baum ist genau dann *kein* Suchbaum, wenn es ein Element gibt,
	// das größer ist als sein rechtes Kind oder kleiner als sein linkes Kind.
	// D.h. wir müssen prüfen, ob es ein solches Element gibt.

	// Das geht am einfachsten rekursiv, also mit einer Hilfsfunktion in Element.
	// Diese muss prüfen, ob das Element leer ist, dann ist es ein Suchbaum.
	// Ansonsten muss es prüfen, ob das Element größer ist als sein linkes Kind
	// oder kleiner als sein rechtes Kind. Wenn ja, ist es kein Suchbaum.
	// Dabei darf diese Prüfung nur ausgeführt werden, wenn das Kind nicht leer ist.
	// Ansonsten ist es ein Suchbaum, wenn beide Kinder Suchbäume sind.

	// TODO
	return false
}

// Hilfsfunktion, um IsSearchTree rekursiv in Element umsetzen zu können.
func (e *Element) IsSearchTree() bool {
	// TODO
	return false
}

// Height gibt die Höhe des Baums zurück.
func (t *Element) Height() int {
	// TODO
	return 0
}

// BalanceFactor gibt den Balancefaktor des Knotens zurück.
func (e *Element) BalanceFactor() int {
	// TODO
	return 0
}

// RotateLeft rotiert den Baum nach links.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeft(root *Element) *Element {
	// TODO
	return nil
}

// RotateRight rotiert den Baum nach rechts.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRight(root *Element) *Element {
	// TODO
	return nil
}

// RotateLeftRight führt eine Doppelrotation nach links aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeftRight(root *Element) *Element {
	// TODO
	return nil
}

// RotateRightLeft führt eine Doppelrotation nach rechts aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRightLeft(root *Element) *Element {
	// TODO
	return nil
}
