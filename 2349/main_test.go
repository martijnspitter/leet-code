package main

import (
	"testing"
)

func TestNumberContainers(t *testing.T) {
	nc := Constructor()

	// Test case 1: find(10) should return -1 as there is no index filled with number 10
	if result := nc.Find(10); result != -1 {
		t.Errorf("expected -1, got %d", result)
	}

	// Test case 2: change(2, 10)
	nc.Change(2, 10)

	// Test case 3: change(1, 10)
	nc.Change(1, 10)

	// Test case 4: change(3, 10)
	nc.Change(3, 10)

	// Test case 5: change(5, 10)
	nc.Change(5, 10)

	// Test case 6: find(10) should return 1 as the smallest index filled with number 10 is 1
	if result := nc.Find(10); result != 1 {
		t.Errorf("expected 1, got %d", result)
	}

	// Test case 7: change(1, 20)
	nc.Change(1, 20)

	// Test case 8: find(10) should return 2 as the smallest index filled with number 10 is now 2
	if result := nc.Find(10); result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}
