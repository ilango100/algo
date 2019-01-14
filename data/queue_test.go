package data

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	if d := q.Enqueue(0).Enqueue(1).Enqueue(2).Dequeue(); d != 0 {
		t.Errorf("Expected 0, got %d", d)
	} else if d = q.Dequeue(); d != 1 {
		t.Errorf("Expected 1, got %d", d)
	} else if d = q.Dequeue(); d != 2 {
		t.Errorf("Expected 2, got %d", d)
	}
}
