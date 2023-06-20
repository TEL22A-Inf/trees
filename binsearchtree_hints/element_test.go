package binsearchtree

import "testing"

func TestElement_NewElement(t *testing.T) {
	e := NewElement()
	if !e.IsEmpty() {
		t.Error("NewElement should return an empty element")
	}
}

func TestElement_SetValue(t *testing.T) {
	e := NewElement()
	e.SetValue(10)
	if e.IsEmpty() {
		t.Error("SetValue should set the value of the element")
	}
	if e.Value != 10 {
		t.Error("SetValue should set the value of the element")
	}
}

func TestElement_Insert_Multiple(t *testing.T) {
	e := NewElement()
	e.Insert(10) // L
	e.Insert(5)  // LL
	e.Insert(15) // R
	e.Insert(3)  // LLL
	e.Insert(7)  // LLR
	e.Insert(42) // RR

	if e.IsEmpty() {
		t.Error("Tree should not be empty")
	}
	if e.Value != 10 {
		t.Error("Root should have value 10")
	}
	if e.Left.Value != 5 {
		t.Error("Node at position L should have value 5, but has", e.Left.Value)
	}
	if e.Right.Value != 15 {
		t.Error("Node at position R should have value 15, but has", e.Right.Value)
	}
	if e.Left.Left.Value != 3 {
		t.Error("Node at position LL should have value 3, but has", e.Left.Left.Value)
	}
	if e.Left.Right.Value != 7 {
		t.Error("Node at position LR should have value 7, but has", e.Left.Right.Value)
	}
	if e.Right.Right.Value != 42 {
		t.Error("Node at position RR should have value 42, but has", e.Right.Right.Value)
	}
}

func TestElement_InOrder(t *testing.T) {
	e := NewElement()
	e.Insert(10) // L
	e.Insert(5)  // LL
	e.Insert(15) // R
	e.Insert(3)  // LLL
	e.Insert(7)  // LLR
	e.Insert(42) // RR

	expected := []int{3, 5, 7, 10, 15, 42}
	actual := e.InOrder()
	if len(expected) != len(actual) {
		t.Error("InOrder should return", len(expected), "elements, but returned", len(actual))
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("InOrder should return", expected, "but returned", actual)
			return
		}
	}
}
