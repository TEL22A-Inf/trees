package binsearchtree

import "testing"

// TestElement_IsLeaf tests the method IsLeaf.
// It checks the following:
// - emtpy element is not a leaf
// - element with value 10 is a leaf
// - element with value 10 and left child: element is not a leaf but left child is
func TestElement_IsLeaf(t *testing.T) {
	// emtpy element is not a leaf
	e1 := NewElement()
	if e1.IsLeaf() {
		t.Error("NewElement should not be a leaf")
	}
	// element with value 10 is a leaf
	e2 := NewElement()
	e2.SetValue(10)
	if !e2.IsLeaf() {
		t.Error("Element with value 10 should be a leaf")
	}
	// element with value 10 and left child: element is not a leaf but left child is
	e3 := NewElement()
	e3.SetValue(10)
	e3.Left.SetValue(5)
	if e3.IsLeaf() {
		t.Error("Element with value 10 should not be a leaf")
	}
	if !e3.Left.IsLeaf() {
		t.Error("Left child of element with value 10 should be a leaf")
	}
}
