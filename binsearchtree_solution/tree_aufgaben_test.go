package binsearchtree

import (
	"testing"
)

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

func TestTree_NodeCount(t *testing.T) {
	// Testcase: empty tree
	tree := NewTree()
	if tree.NodeCount() != 0 {
		t.Error("NodeCount should return 0 for an empty tree")
	}

	// Testcase: tree with one element
	tree = NewTree()
	tree.Insert(10)
	if tree.NodeCount() != 1 {
		t.Errorf("NodeCount should return 1 for a tree with one element, but returned %d", tree.NodeCount())
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
	if tree.NodeCount() != 7 {
		t.Errorf("NodeCount should return 7 for the given tree, but returned %d", tree.NodeCount())
	}
}

func TestTree_IsSearchTree(t *testing.T) {
	// Testcase: empty tree
	tree := NewTree()
	if !tree.IsSearchTree() {
		t.Error("IsSearchTree should return true for an empty tree")
	}

	// Testcase: tree with one element
	tree = NewTree()
	tree.Insert(10)
	if !tree.IsSearchTree() {
		t.Error("IsSearchTree should return true for a tree with one element")
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
	if !tree.IsSearchTree() {
		t.Error("IsSearchTree should return true for the given tree")
	}

	// Testcase: tree with wrong order
	tree = NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Root.Right.SetValue(3)
	if tree.IsSearchTree() {
		t.Error("IsSearchTree should return false for the given tree")
	}
}

func TestElement_Height(t *testing.T) {
	root := NewElement()

	actualheight := root.Height()
	expectedheight := 0
	if actualheight != expectedheight {
		t.Errorf("Height should return %d, but returned %d", expectedheight, actualheight)
	}

	root.Insert(10)
	actualheight = root.Height()
	expectedheight = 1
	if actualheight != expectedheight {
		t.Errorf("Height should return %d, but returned %d", expectedheight, actualheight)
	}

	root.Insert(5)

	actualheight = root.Height()
	expectedheight = 2
	if actualheight != expectedheight {
		t.Errorf("Height should return %d, but returned %d", expectedheight, actualheight)
	}

	root.Insert(15)

	actualheight = root.Height()
	expectedheight = 2
	if actualheight != expectedheight {
		t.Errorf("Height should return %d, but returned %d", expectedheight, actualheight)
	}

	root.Insert(12)

	actualheight = root.Height()
	expectedheight = 3
	if actualheight != expectedheight {
		t.Errorf("Height should return %d, but returned %d", expectedheight, actualheight)
	}
}

func TestElement_BalanceFactor(t *testing.T) {
	root := NewElement()

	actualbalance := root.BalanceFactor()
	expectedbalance := 0
	if actualbalance != expectedbalance {
		t.Errorf("BalanceFactor should return %d, but returned %d", expectedbalance, actualbalance)
	}

	root.Insert(10)
	actualbalance = root.BalanceFactor()
	expectedbalance = 0
	if actualbalance != expectedbalance {
		t.Errorf("BalanceFactor should return %d, but returned %d", expectedbalance, actualbalance)
	}

	root.Insert(5)
	actualbalance = root.BalanceFactor()
	expectedbalance = -1
	if actualbalance != expectedbalance {
		t.Errorf("BalanceFactor should return %d, but returned %d", expectedbalance, actualbalance)
	}

	root.Insert(15)
	actualbalance = root.BalanceFactor()
	expectedbalance = 0
	if actualbalance != expectedbalance {
		t.Errorf("BalanceFactor should return %d, but returned %d", expectedbalance, actualbalance)
	}

	root.Insert(12)
	actualbalance = root.BalanceFactor()
	expectedbalance = 1
	if actualbalance != expectedbalance {
		t.Errorf("BalanceFactor should return %d, but returned %d", expectedbalance, actualbalance)
	}
}
