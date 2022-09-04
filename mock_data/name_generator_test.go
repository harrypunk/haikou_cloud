package mockdata

import (
	"testing"
)

func namesN(n int, t *testing.T) {
	gen := NewWithLen(n)
	count := 0
	for names := range gen.GetNames() {
		t.Logf("names: %v", names)
		count += 1
	}
	if count != n {
		t.Fail()
	}
}

func TestNames5(t *testing.T) {
	namesN(5, t)
}

func TestNames20(t *testing.T) {
	namesN(20, t)
}

func TestNames40(t *testing.T) {
	namesN(40, t)
}
