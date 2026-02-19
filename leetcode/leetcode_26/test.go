package main

import (
	"reflect"
	"testing"
)

// TestRemoveDuplicates tests the removeDuplicates function with different test cases
// It verifies both the returned length k and the modified array content
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string // Test case name
		input    []int  // Input array for the test
		expected []int  // Expected array after removing duplicates (first k elements)
		k        int    // Expected length of the unique elements
	}{
		{
			name:     "basic test",
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
			k:        2,
		},
		{
			name:     "multiple duplicates",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
			k:        5,
		},
	}

	// Iterate through all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of input to avoid modifying the original test data
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)

			// Execute the function under test
			k := removeDuplicates(nums)

			// Verify the returned length k is correct
			if k != tt.k {
				t.Errorf("expected length %d, got %d", tt.k, k)
			}

			// Verify the first k elements of the array are the expected unique values
			if !reflect.DeepEqual(nums[:k], tt.expected) {
				t.Errorf("expected array %v, got %v", tt.expected, nums[:k])
			}
		})
	}
}
