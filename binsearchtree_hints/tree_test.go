package binsearchtree

import "testing"

func TestTree_NewTree(t *testing.T) {
	tree := NewTree()
	if !tree.IsEmpty() {
		t.Error("NewTree should return an empty tree")
	}
}

func TestTree_Insert(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	if tree.IsEmpty() {
		t.Error("Tree should not be empty")
	}
	if tree.Root.Value != 10 {
		t.Error("Root should have value 10, but has", tree.Root.Value)
	}
}
