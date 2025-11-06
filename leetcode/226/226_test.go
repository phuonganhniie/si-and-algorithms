package leetcode_226

import (
	"fmt"
	"reflect"
	"testing"
)

// Helper function to create a tree node
func newNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

// Helper function to build a tree from level-order array (nil for missing nodes)
func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := newNode(values[0].(int))
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.leftNode = newNode(values[i].(int))
			queue = append(queue, node.leftNode)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.rightNode = newNode(values[i].(int))
			queue = append(queue, node.rightNode)
		}
		i++
	}

	return root
}

// Helper function to convert tree to level-order array for comparison
func treeToArray(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.val)
			queue = append(queue, node.leftNode)
			queue = append(queue, node.rightNode)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name  string
		input []interface{}
		want  []interface{}
	}{
		{
			name:  "Example 1: Complete binary tree",
			input: []interface{}{4, 2, 7, 1, 3, 6, 9},
			want:  []interface{}{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:  "Example 2: Small tree",
			input: []interface{}{2, 1, 3},
			want:  []interface{}{2, 3, 1},
		},
		{
			name:  "Example 3: Empty tree",
			input: []interface{}{},
			want:  []interface{}{},
		},
		{
			name:  "Single node",
			input: []interface{}{1},
			want:  []interface{}{1},
		},
		{
			name:  "Only left child",
			input: []interface{}{1, 2, nil},
			want:  []interface{}{1, nil, 2},
		},
		{
			name:  "Only right child",
			input: []interface{}{1, nil, 2},
			want:  []interface{}{1, 2},
		},
		{
			name:  "Left-skewed tree",
			input: []interface{}{1, 2, nil, 3, nil},
			want:  []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:  "Right-skewed tree",
			input: []interface{}{1, nil, 2, nil, 3},
			want:  []interface{}{1, 2, nil, 3},
		},
		{
			name:  "Three levels, complete",
			input: []interface{}{1, 2, 3, 4, 5, 6, 7},
			want:  []interface{}{1, 3, 2, 7, 6, 5, 4},
		},
		{
			name:  "Three levels, incomplete left",
			input: []interface{}{1, 2, 3, nil, 5, 6, 7},
			want:  []interface{}{1, 3, 2, 7, 6, 5},
		},
		{
			name:  "Three levels, incomplete right",
			input: []interface{}{1, 2, 3, 4, 5, nil, 7},
			want:  []interface{}{1, 3, 2, 7, nil, 5, 4},
		},
		{
			name:  "Symmetric tree",
			input: []interface{}{1, 2, 2, 3, 4, 4, 3},
			want:  []interface{}{1, 2, 2, 3, 4, 4, 3},
		},
		{
			name:  "All left children",
			input: []interface{}{1, 2, nil, 3, nil, 4, nil},
			want:  []interface{}{1, nil, 2, nil, 3, nil, 4},
		},
		{
			name:  "All right children",
			input: []interface{}{1, nil, 2, nil, 3, nil, 4},
			want:  []interface{}{1, 2, nil, 3, nil, 4},
		},
		{
			name:  "Large values",
			input: []interface{}{100, 50, 150, 25, 75, 125, 175},
			want:  []interface{}{100, 150, 50, 175, 125, 75, 25},
		},
		{
			name:  "Negative values",
			input: []interface{}{0, -5, 5, -10, -3, 3, 10},
			want:  []interface{}{0, 5, -5, 10, 3, -3, -10},
		},
		{
			name:  "Mixed positive and negative",
			input: []interface{}{1, -2, 3, -4, 5},
			want:  []interface{}{1, 3, -2, nil, nil, 5, -4},
		},
		{
			name:  "Four levels deep",
			input: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want:  []interface{}{1, 3, 2, 7, 6, 5, 4, 15, 14, 13, 12, 11, 10, 9, 8},
		},
		{
			name:  "Unbalanced tree - left heavy",
			input: []interface{}{1, 2, 3, 4, 5},
			want:  []interface{}{1, 3, 2, nil, nil, 5, 4},
		},
		{
			name:  "Unbalanced tree - right heavy",
			input: []interface{}{1, 2, 3, nil, nil, 6, 7},
			want:  []interface{}{1, 3, 2, 7, 6},
		},
		{
			name:  "Two nodes - parent and left child",
			input: []interface{}{5, 3, nil},
			want:  []interface{}{5, nil, 3},
		},
		{
			name:  "Two nodes - parent and right child",
			input: []interface{}{5, nil, 7},
			want:  []interface{}{5, 7},
		},
		{
			name:  "Complex incomplete tree",
			input: []interface{}{1, 2, 3, nil, 4, nil, 5, nil, nil, 6, 7},
			want:  []interface{}{1, 3, 2, 5, nil, 4, nil, 7, 6},
		},
		{
			name:  "Zigzag pattern",
			input: []interface{}{1, 2, nil, nil, 3, nil, 4},
			want:  []interface{}{1, nil, 2, 3, nil, 4},
		},
		{
			name:  "Single branch going deep",
			input: []interface{}{1, 2, nil, 3, nil, 4, nil, 5, nil},
			want:  []interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build input tree
			root := buildTree(tt.input)

			// Invert the tree
			result := invertTree(root)

			// Convert result to array
			got := treeToArray(result)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
				fmt.Printf("Input: %v\n", tt.input)
			} else {
				fmt.Printf("✓ %s: PASSED\n", tt.name)
			}
		})
	}
}

// Test that the original tree is modified (not a copy)
func TestInvertTreeModifiesOriginal(t *testing.T) {
	root := buildTree([]interface{}{1, 2, 3})
	result := invertTree(root)

	if root != result {
		t.Errorf("invertTree should modify and return the original tree, not create a copy")
	}
}

// Test nil input
func TestInvertTreeNil(t *testing.T) {
	result := invertTree(nil)
	if result != nil {
		t.Errorf("invertTree(nil) should return nil, got %v", result)
	} else {
		fmt.Println("✓ Nil input: PASSED")
	}
}

// Benchmark test to measure performance
func BenchmarkInvertTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Rebuild tree for each iteration since it gets modified
		testTree := buildTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
		invertTree(testTree)
	}
}

// Benchmark with large tree
func BenchmarkInvertTreeLarge(b *testing.B) {
	// Create a larger tree
	values := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testTree := buildTree(values)
		invertTree(testTree)
	}
}

// Benchmark with deep unbalanced tree
func BenchmarkInvertTreeDeep(b *testing.B) {
	// Create a deep left-skewed tree
	values := []interface{}{1, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9, nil, 10, nil}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testTree := buildTree(values)
		invertTree(testTree)
	}
}
