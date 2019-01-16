package data

import (
	"testing"
)

var b = &BinTree{Key: 3}

func TestBinTreeInsert(t *testing.T) {
	b.Insert(2, nil)
	b.Insert(1, nil)
	b.Insert(4, nil)
	b.Insert(5, nil)
}
func TestBinTreeMinMax(t *testing.T) {
	if k := b.Min().Key; k != 1 {
		t.Errorf("Expected min key 1, got %d", k)
		return
	}
	if k := b.Max().Key; k != 5 {
		t.Errorf("Expected max key 5, got %d", k)
		return
	}
}
func TestBinTreeSearch(t *testing.T) {
	if k := b.Search(2); k.Key != 2 {
		t.Errorf("Expected key 2, got %d", k.Key)
		return
	}
	if k := b.Search(4); k.Key != 4 {
		t.Errorf("Expected key 4, got %d", k.Key)
		return
	}
	if k := b.SearchIt(2); k.Key != 2 {
		t.Errorf("Expected key 2, got %d", k.Key)
		return
	}
	if k := b.SearchIt(4); k.Key != 4 {
		t.Errorf("Expected key 4, got %d", k.Key)
		return
	}
}

func TestBinTreeNotFound(t *testing.T) {
	if k := b.Search(0); k != nil {
		t.Errorf("Expected nil, got %#v", k)
		return
	}
	if k := b.SearchIt(0); k != nil {
		t.Errorf("Expected nil, got %#v", k)
		return
	}
	if k := b.Search(10); k != nil {
		t.Errorf("Expected nil, got %#v", k)
		return
	}
	if k := b.SearchIt(10); k != nil {
		t.Errorf("Expected nil, got %#v", k)
		return
	}
}
