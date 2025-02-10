package logic

import "testing"

func TestFindLargest(t *testing.T) {
	result := FindLargest(12, 45, 23)
	if result != 45 {
		t.Errorf("FindLargest(12, 45, 23) = %d; want 45", result)
	}
}
