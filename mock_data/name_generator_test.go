package mockdata

import (
	"testing"
)

func TestNameGenerate(t *testing.T) {
	gen := NewWithLen(10)
	for names := range gen.GetNames() {
		t.Logf("names: %v", names)
	}
}
