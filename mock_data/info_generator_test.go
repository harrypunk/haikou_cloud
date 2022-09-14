package mock_data

import (
	"testing"
)

func familiesN(n int, t *testing.T) {
	gen := NewWithSeed(100)
	count := 0
	for families := range gen.GetFamilyNames(n) {
		t.Logf("families: %v", families)
		count += 1
	}
	if count != n {
		t.Fail()
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
	count := 0
	for name := range gen.RandomNameList(n) {
		t.Log(name)
		count += 1
	}
	if count != n {
		t.Fail()
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
	count := 0
	for phone := range gen.RandomPhone(10) {
		t.Log(phone)
		count += 1
	}
	if count != 10 {
		t.Fail()
	}
}
