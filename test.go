package tree

import "testing"

func TestTrueCase(t *testing.T) {
	tree := TreeNode{
		HasToy: false,
		Left: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: false,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
		},
		Right: &TreeNode{
			HasToy: true,
			Left:   nil,
			Right:  nil,
		},
	}

	got := areToysBalanced(&tree)
	if got != true {
		t.Error("Got false, but true needed")
	}
}

func TestFalseCase(t *testing.T) {
	tree := TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left:   nil,
			Right:  nil,
		},
		Right: &TreeNode{
			HasToy: false,
			Left:   nil,
			Right:  nil,
		},
	}

	got := areToysBalanced(&tree)
	if got != false {
		t.Error("Got true, but false needed")
	}
}
