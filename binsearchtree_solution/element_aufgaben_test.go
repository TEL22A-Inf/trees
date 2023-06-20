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

// TestElement_GetMinNode tests the method GetMinNode.
func TestElement_GetMinNode(t *testing.T) {
	// Testcase: empty tree
	e1 := NewElement()
	minNode := e1.GetMinNode()
	if minNode != nil {
		t.Error("GetMinNode should return nil for empty tree")
	}

	// Testcase: tree with one element
	e2 := NewElement()
	e2.SetValue(10)
	minNode = e2.GetMinNode()
	if minNode != e2 {
		t.Error("GetMinNode should return root element for tree with one element")
	}

	// Testcase: tree with two elements on the left side
	e3 := NewElement()
	e3.SetValue(10)
	e3.Left.SetValue(5)
	minNode = e3.GetMinNode()
	if minNode != e3.Left {
		t.Error("GetMinNode should return leftmost element for tree with two elements on the left side")
	}

	// Testcase: tree with two elements on either sides
	e4 := NewElement()
	e4.SetValue(10)
	e4.Left.SetValue(5)
	e4.Right.SetValue(15)
	minNode = e4.GetMinNode()
	if minNode != e4.Left {
		t.Error("GetMinNode should return left element for tree with two elements on either sides")
	}

	// Testcase:
	//    10
	//   /  \
	//  5   15
	//   \
	//    7
	e5 := NewElement()
	e5.SetValue(10)
	e5.Left.SetValue(5)
	e5.Left.Right.SetValue(7)
	e5.Right.SetValue(15)
	minNode = e5.GetMinNode()
	if minNode != e5.Left {
		t.Error("GetMinNode should return left child of root")
	}
}

// TestElement_GetMaxNode tests the method GetMaxNode.
func TestElement_GetMaxNode(t *testing.T) {
	// Testcase: empty tree
	e1 := NewElement()
	maxNode := e1.GetMaxNode()
	if maxNode != nil {
		t.Error("GetMaxNode should return nil for empty tree")
	}

	// Testcase: tree with one element
	e2 := NewElement()
	e2.SetValue(10)
	maxNode = e2.GetMaxNode()
	if maxNode != e2 {
		t.Error("GetMaxNode should return root element for tree with one element")
	}

	// Testcase: tree with two elements on the right side
	e3 := NewElement()
	e3.SetValue(10)
	e3.Right.SetValue(15)
	maxNode = e3.GetMaxNode()
	if maxNode != e3.Right {
		t.Error("GetMaxNode should return rightmost element for tree with two elements on the right side")
	}

	// Testcase: tree with two elements on either sides
	e4 := NewElement()
	e4.SetValue(10)
	e4.Left.SetValue(5)
	e4.Right.SetValue(15)
	maxNode = e4.GetMaxNode()
	if maxNode != e4.Right {
		t.Error("GetMaxNode should return right element for tree with two elements on either sides")
	}

	// Testcase:
	//    10
	//   /  \
	//  5   15
	//     /
	//    12
	e5 := NewElement()
	e5.SetValue(10)
	e5.Left.SetValue(5)
	e5.Right.SetValue(15)
	e5.Right.Left.SetValue(12)
	maxNode = e5.GetMaxNode()
	if maxNode != e5.Right {
		t.Error("GetMaxNode should return right child of root")
	}
}

// TestElement_RemoveRoot tests the method RemoveRoot.
func TestElement_RemoveRoot(t *testing.T) {
	// Testcase: empty tree
	e1 := NewElement()
	e1.RemoveRoot()
	if !e1.IsEmpty() {
		t.Error("RemoveRoot should not change empty tree")
	}

	// Testcase: tree with one element
	e2 := NewElement()
	e2.SetValue(10)
	e2.RemoveRoot()
	if !e2.IsEmpty() {
		t.Error("RemoveRoot should remove root element")
	}

	// Testcase: more complex tree
	//             10
	//           /    \
	//          5     15
	//         / \    / \
	//        3   7  12  17
	e3 := NewElement()
	e3.SetValue(10)
	e3.Left.SetValue(5)
	e3.Right.SetValue(15)
	e3.Left.Left.SetValue(3)
	e3.Left.Right.SetValue(7)
	e3.Right.Left.SetValue(12)
	e3.Right.Right.SetValue(17)
	e3.RemoveRoot()

	// check tree structure
	if e3.Value != 7 {
		t.Error("RemoveRoot should set new root element")
	}
	if e3.Left.Value != 5 {
		t.Error("RemoveRoot should leave element \"l\" unchanged")
	}
	if e3.Right.Value != 15 {
		t.Error("RemoveRoot should leave element \"r\" unchanged")
	}
	if e3.Left.Left.Value != 3 {
		t.Error("RemoveRoot should leave element \"ll\" unchanged")
	}
	if !e3.Left.Right.IsEmpty() {
		t.Error("RemoveRoot should remove element \"lr\"")
	}
	if e3.Right.Left.Value != 12 {
		t.Error("RemoveRoot should leave element \"rl\" unchanged")
	}
	if e3.Right.Right.Value != 17 {
		t.Error("RemoveRoot should leave element \"rr\" unchanged")
	}
}
