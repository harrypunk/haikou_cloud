package mock_data

import (
	"testing"
)

func familiesN(n int, t *testing.T) {
	gen := NewWithSeed(100)
	ch := gen.GetFamilyNames()
	for i := 0; i < n; i++ {
		family := <-ch
		t.Logf("families: %v", family)
	}
}

func TestFamilies5(t *testing.T) {
	familiesN(5, t)
}

func TestFamilies20(t *testing.T) {
	familiesN(20, t)
}

func TestFamilies40(t *testing.T) {
	familiesN(40, t)
}

func nameListN(n int, t *testing.T) {
	gen := NewWithSeed(101)
	ch := gen.RandomNameList()
	for i := 0; i < n; i++ {
		t.Log(<-ch)
	}
}

func TestNames5(t *testing.T) {
	nameListN(5, t)
}

func TestNames10(t *testing.T) {
	nameListN(10, t)
}

func TestPhone10(t *testing.T) {
	gen := NewWithSeed(100)
	ch := gen.RandomPhone()
	for i := 0; i < 10; i++ {
		t.Log(<-ch)
	}
}
