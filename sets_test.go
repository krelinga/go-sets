package sets_test

import (
	"testing"

	"github.com/krelinga/go-sets"
)

func TestSet(t *testing.T) {
	// Create a new set
	s := sets.New[int]()

	// Add items to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Check if items are present
	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Error("Set should contain 1, 2, and 3")
	}

	// Check if an item not in the set is absent
	if s.Has(4) {
		t.Error("Set should not contain 4")
	}

	// Remove an item and check if it was removed
	if !s.Del(2) {
		t.Error("Set should have removed 2")
	}
	if s.Has(2) {
		t.Error("Set should not contain 2 after removal")
	}

	// Try to remove an item that is not present
	if s.Del(5) {
		t.Error("Set should not have removed 5, as it was not present")
	}

	var saw1, saw3 bool
	for item := range s.Values() {
		if item == 1 {
			saw1 = true
		} else if item == 3 {
			saw3 = true
		} else {
			t.Errorf("Unexpected item in set: %d", item)
		}
	}
	if !saw1 || !saw3 {
		t.Error("Set should only contain 1 and 3 after operations")
	}
}
