package binsearchtree

import "testing"

func TestTree_RemoveValue(t *testing.T) {
	// Testcase: empty tree
	tree := NewTree()
	tree.RemoveValue(10)
	if !tree.IsEmpty() {
		t.Error("RemoveValue should not change an empty tree")
	}

	// Testcase: tree with one element
	tree = NewTree()
	tree.Insert(10)
	tree.RemoveValue(10)
	if !tree.IsEmpty() {
		t.Error("RemoveValue should remove the root element")
	}

	// Testcase: more complex tree
	tree = NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(12)
	tree.Insert(17)
	tree.Insert(7)
	tree.Insert(3)
	tree.RemoveValue(10)
	if tree.Root.Value != 7 {
		t.Error("RemoveValue should set new root element")
	}
	if tree.Root.Left.Value != 5 {
		t.Error("RemoveValue should leave element \"l\" unchanged")
	}
	if tree.Root.Right.Value != 15 {
		t.Error("RemoveValue should leave element \"r\" unchanged")
	}
	if tree.Root.Left.Left.Value != 3 {
		t.Error("RemoveValue should leave element \"ll\" unchanged")
	}
	if !tree.Root.Left.Right.IsEmpty() {
		t.Error("RemoveValue should remove element \"lr\"")
	}
	if tree.Root.Right.Left.IsEmpty() {
		t.Error("RemoveValue should leave element \"rl\" unchanged")
	}
	if tree.Root.Right.Right.Value != 17 {
		t.Error("RemoveValue should leave element \"rr\" unchanged")
	}
}