package minheap

import (
	"testing"
)

// TestMinHeap_New_IsEmpty prüft, ob ein neuer Heap leer ist.
func TestMinHeap_New_IsEmpty(t *testing.T) {
	h := NewMinHeap()
	if !h.IsEmpty() {
		t.Errorf("New heap is not empty")
	}
}

// TestMinHeap_GetMin_Empty prüft, ob GetMin eine Panic auslöst, wenn der Heap leer ist.
func TestMinHeap_GetMin_Empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetMin did not panic")
		}
	}()
	h := NewMinHeap()
	h.GetMin()
}

// TestMinHeap_Insert_One_Element fügt ein Element in den Heap ein und prüft:
// * ob der Heap nicht mehr leer ist.
// * ob das Element an der Wurzel steht.
func TestMinHeap_Insert_One_Element(t *testing.T) {
	h := NewMinHeap()
	h.Insert(5)
	if h.IsEmpty() {
		t.Errorf("Heap is empty after insert")
	}
	min := h.GetMin()
	if min != 5 {
		t.Errorf("Expected %v at root, but got %v.", 5, min)
	}
}

// TestMinHeap_Insert_Two_Elements fügt zwei Elemente in den Heap ein und prüft:
// * ob der Heap nicht mehr leer ist.
// * ob das kleinere Element an der Wurzel steht.
func TestMinHeap_Insert_Two_Elements(t *testing.T) {
	h := NewMinHeap()
	h.Insert(5)
	h.Insert(3)
	if h.IsEmpty() {
		t.Errorf("Heap is empty after insert")
	}
	min := h.GetMin()
	if min != 3 {
		t.Errorf("Expected %v at root, but got %v.", 3, min)
	}
}

// TestMinHeap_RemoveMin_One_Element fügt ein Element in den Heap ein und entfernt es wieder.
// Es wird geprüft, ob der Heap leer ist.
func TestMinHeap_RemoveMin_One_Element(t *testing.T) {
	h := NewMinHeap()
	h.Insert(5)
	h.RemoveMin()
	if !h.IsEmpty() {
		t.Errorf("Heap is not empty after remove")
	}
}

// TestMinHeap_RemoveMin_Two_Elements fügt zwei Elemente in den Heap ein und entfernt das kleinere.
// Es wird geprüft, ob das größere Element an der Wurzel steht und ob der Heap nicht leer ist.
func TestMinHeap_RemoveMin_Two_Elements(t *testing.T) {
	h := NewMinHeap()
	h.Insert(5)
	h.Insert(3)
	h.RemoveMin()
	if h.IsEmpty() {
		t.Errorf("Heap is empty after remove")
	}
	min := h.GetMin()
	if min != 5 {
		t.Errorf("Expected %v at root, but got %v.", 5, min)
	}
}

// TestMinHeap_Remove_Empty entfernt ein Element aus einem leeren Heap.
// Es wird geprüft, ob der Heap leer bleibt.
func TestMinHeap_Remove_Empty(t *testing.T) {
	h := NewMinHeap()
	h.RemoveMin()
	if !h.IsEmpty() {
		t.Errorf("Heap is not empty after remove")
	}
}

// TestMinHeap_BubbleUp_Root baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element ans Ende ein, das kleiner ist als alle bisherigen.
// Es wird geprüft, ob die Elemente durch BubbleUp an die richtigen Positionen wandern.
func TestMinHeap_BubbleUp_Root(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{5, 7, 9, 2}
	h.BubbleUp(3) // Element an Stelle 3 (2) nach oben schieben.
	actual := h.elements
	expected := []int{2, 5, 9, 7}

	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleUp_Middle baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element ans Ende ein, das an eine mittlere Position gehört.
// Es wird geprüft, ob die Elemente durch BubbleUp an die richtigen Positionen wandern.
func TestMinHeap_BubbleUp_Middle(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{5, 7, 9, 6}
	h.BubbleUp(3) // Element an Stelle 3 (6) nach oben schieben.
	actual := h.elements
	expected := []int{5, 6, 9, 7}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleUp_End baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element ans Ende ein, das an die letzte Position gehört.
// Es wird geprüft, ob die Elemente durch BubbleUp an die richtigen Positionen wandern.
func TestMinHeap_BubbleUp_End(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{5, 7, 9, 10}
	h.BubbleUp(3) // Element an Stelle 3 (10) nach oben schieben.
	actual := h.elements
	expected := []int{5, 7, 9, 10}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleDown_Max baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element an die Wurzel ein, das größer ist als alle bisherigen.
// Es wird geprüft, ob die Elemente durch BubbleDown an die richtigen Positionen wandern.
func TestMinHeap_BubbleDown_Max(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{20, 7, 11, 9}
	h.BubbleDown(0) // Element an Stelle 0 (20) nach unten schieben.
	actual := h.elements
	expected := []int{7, 9, 11, 20}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleDown_Middle baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element an die Wurzel ein, das an eine mittlere Position gehört.
// Es wird geprüft, die Elemente durch BubbleDown an die richtigen Positionen wandern.
func TestMinHeap_BubbleDown_Middle(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{8, 7, 11, 9}
	h.BubbleDown(0) // Element an Stelle 0 (8) nach unten schieben.
	actual := h.elements
	expected := []int{7, 8, 11, 9}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleDown_End baut einen Heap mit 3 Elementen manuell auf und fügt
// ein Element an die Wurzel ein, das an die Wurzel gehört.
// Es wird geprüft, ob die Elemente durch BubbleDown an die richtigen Positionen wandern.
func TestMinHeap_BubbleDown_End(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{7, 8, 11, 9}
	h.BubbleDown(0) // Element an Stelle 0 (7) nach unten schieben.
	actual := h.elements
	expected := []int{7, 8, 11, 9}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// TestMinHeap_BubbleDown_BubbleRight baut einen Heap manuell auf und fügt
// ein Element an die Wurzel ein, das nach ganz rechts unten wandern muss.
// Es wird geprüft, ob die Elemente durch BubbleDown an die richtigen Positionen wandern.
func TestMinHeap_BubbleDown_BubbleRight(t *testing.T) {
	h := NewMinHeap()
	h.elements = []int{20, 9, 7, 11, 10, 16, 15}
	h.BubbleDown(0) // Element an Stelle 0 (20) nach unten schieben.
	actual := h.elements
	expected := []int{7, 9, 15, 11, 10, 16, 20}
	if !equalLists(expected, actual) {
		t.Errorf("Expected elements %v, but got %v.", expected, actual)
	}
}

// equalLists prüft, ob zwei Listen gleich sind (Hilfsfunktion für die Tests).
// Anmerkung: Die Funktion ist ein Generic, d.h. sie kann mit beliebigen Typen aufgerufen werden,
// solange diese vergleichbar sind (der Typparameter T muss die Eigenschaft "comparable" haben).
func equalLists[T comparable](expected, actual []T) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i, v := range expected {
		if actual[i] != v {
			return false
		}
	}
	return true
}
