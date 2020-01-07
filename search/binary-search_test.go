package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var list = []string{"Adriana", "Andre", "Carolina", "Eliane", "Marcos", "Ofelia", "Pedro"}

// TestBinarySearch run tests for BinarySearch
func TestBinarySearch(t *testing.T) {
	index := BinarySearch(list, "Ofelia")
	assert.Equal(t, index, 5, "Ofelia must be in position 5")
}

func TestBinarySearchNotFound(t *testing.T) {
	index := BinarySearch(list, "Lucas")
	assert.Equal(t, index, -1, "Lucas must not be found")
}
