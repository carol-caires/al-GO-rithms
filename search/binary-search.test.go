package search_test

import (
	"github.com/carol-caires/al-GO-rithms/search"
	"testing"
)

// TestBinarySearch run tests for BinarySearch
func TestBinarySearch(*testing.T) {
	search.BinarySearch([]string{"Adriana", "Andre", "Carolina", "Eliane", "Marcos", "Ofelia", "Pedro"}, "Ofelia")
}
