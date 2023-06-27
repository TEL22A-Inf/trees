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

func TestRotateLeft_simple(t *testing.T) {
	root := NewElement()

	// Structure before rotation:
	//      10
	//        \
	//         15
	//           \
	//            20
	root.Insert(10)
	root.Insert(15)
	root.Insert(20)

	root = RotateLeft(root)

	// Expected structure after rotation:
	//      15
	//     /  \
	//    10  20

	// Check structure:
	if root.Value != 15 {
		t.Errorf("root should be 15, but is %d", root.Value)
	}
	if root.Left.Value != 10 {
		t.Errorf("root.Left should be 10, but is %d", root.Left.Value)
	}
	if root.Right.Value != 20 {
		t.Errorf("root.Right should be 20, but is %d", root.Right.Value)
	}
}

func TestRotateLeft_complex(t *testing.T) {
	root := NewElement()

	// Structure before rotation:
	//      10
	//     /  \
	//    5    15
	//        /  \
	//       12   20
	root.Insert(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(20)
	root.Insert(12)

	root = RotateLeft(root)

	// Expected structure after rotation:
	//      15
	//     /  \
	//    10  20
	//   /  \
	//  5    12

	// Check structure:
	if root.Value != 15 {
		t.Errorf("root should be 15, but is %d", root.Value)
	}
	if root.Left.Value != 10 {
		t.Errorf("root.Left should be 10, but is %d", root.Left.Value)
	}
	if root.Right.Value != 20 {
		t.Errorf("root.Right should be 20, but is %d", root.Right.Value)
	}
	if root.Left.Left.Value != 5 {
		t.Errorf("root.Left.Left should be 5, but is %d", root.Left.Left.Value)
	}
	if root.Left.Right.Value != 12 {
		t.Errorf("root.Left.Right should be 12, but is %d", root.Left.Right.Value)
	}
}

func TestRotateRight_simple(t *testing.T) {
	root := NewElement()

	// Structure before rotation:
	//      15
	//     /
	//    10
	//   /
	//  5
	root.Insert(15)
	root.Insert(10)
	root.Insert(5)

	root = RotateRight(root)

	// Expected structure after rotation:
	//      10
	//     /  \
	//    5   15

	// Check structure:
	if root.Value != 10 {
		t.Errorf("root should be 10, but is %d", root.Value)
	}
	if root.Left.Value != 5 {
		t.Errorf("root.Left should be 5, but is %d", root.Left.Value)
	}
	if root.Right.Value != 15 {
		t.Errorf("root.Right should be 15, but is %d", root.Right.Value)
	}
}

func TestRotateRight_complex(t *testing.T) {
	root := NewElement()

	// Structure before rotation:
	//      15
	//     /  \
	//    10  20
	//   /  \
	//  5    12
	root.Insert(15)
	root.Insert(10)
	root.Insert(20)
	root.Insert(5)
	root.Insert(12)

	root = RotateRight(root)

	// Expected structure after rotation:
	//      10
	//     /  \
	//    5   15
	//       /  \
	//      12  20

	// Check structure:
	if root.Value != 10 {
		t.Errorf("root should be 10, but is %d", root.Value)
	}
	if root.Left.Value != 5 {
		t.Errorf("root.Left should be 5, but is %d", root.Left.Value)
	}
	if root.Right.Value != 15 {
		t.Errorf("root.Right should be 15, but is %d", root.Right.Value)
	}
	if root.Right.Left.Value != 12 {
		t.Errorf("root.Right.Left should be 12, but is %d", root.Right.Left.Value)
	}
	if root.Right.Right.Value != 20 {
		t.Errorf("root.Right.Right should be 20, but is %d", root.Right.Right.Value)
	}
}

func TestElement_RotateLeftRight(t *testing.T) {
	root := NewElement()

	// Expected rotation:
	//      15               15                  12
	//     /  \             /  \               /    \
	//    10  20           12  20            10      15
	//   /  \      -->    /  \        -->   /  \    /  \
	//  5    12          10   14           5    11 14   20
	//      /  \        / \
	//     11  	14     5   11
	root.Insert(15)
	root.Insert(10)
	root.Insert(20)
	root.Insert(5)
	root.Insert(12)
	root.Insert(11)
	root.Insert(14)

	root = RotateLeftRight(root)

	// Check structure:
	if root.Value != 12 {
		t.Errorf("root should be 12, but is %d", root.Value)
	}
	if root.Left.Value != 10 {
		t.Errorf("root.Left should be 10, but is %d", root.Left.Value)
	}
	if root.Right.Value != 15 {
		t.Errorf("root.Right should be 15, but is %d", root.Right.Value)
	}
	if root.Left.Left.Value != 5 {
		t.Errorf("root.Left.Left should be 5, but is %d", root.Left.Left.Value)
	}
	if root.Left.Right.Value != 11 {
		t.Errorf("root.Left.Right should be 11, but is %d", root.Left.Right.Value)
	}
	if root.Right.Left.Value != 14 {
		t.Errorf("root.Right.Left should be 14, but is %d", root.Right.Left.Value)
	}
	if root.Right.Right.Value != 20 {
		t.Errorf("root.Right.Right should be 20, but is %d", root.Right.Right.Value)
	}
}

func TestElement_RotateRightLeft(t *testing.T) {
	root := NewElement()

	// Expected rotation:
	//      15               15                      17
	//     /  \             /  \                   /    \
	//    10  20           10  17               15        20
	//       /  \    -->      /  \     -->     /  \      /  \
	//      17   25          16  20          10   16   18   25
	//     / \                  /  \
	//    16  18               18  25
	root.Insert(15)
	root.Insert(10)
	root.Insert(20)
	root.Insert(17)
	root.Insert(25)
	root.Insert(16)
	root.Insert(18)

	root = RotateRightLeft(root)

	// Check structure:
	if root.Value != 17 {
		t.Errorf("root should be 17, but is %d", root.Value)
	}
	if root.Left.Value != 15 {
		t.Errorf("root.Left should be 15, but is %d", root.Left.Value)
	}
	if root.Right.Value != 20 {
		t.Errorf("root.Right should be 20, but is %d", root.Right.Value)
	}
	if root.Left.Left.Value != 10 {
		t.Errorf("root.Left.Left should be 10, but is %d", root.Left.Left.Value)
	}
	if root.Left.Right.Value != 16 {
		t.Errorf("root.Left.Right should be 16, but is %d", root.Left.Right.Value)
	}
	if root.Right.Left.Value != 18 {
		t.Errorf("root.Right.Left should be 18, but is %d", root.Right.Left.Value)
	}
	if root.Right.Right.Value != 25 {
		t.Errorf("root.Right.Right should be 25, but is %d", root.Right.Right.Value)
	}
}
