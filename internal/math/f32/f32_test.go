package mathf32_test

import (
	"testing"

	mathf32 "github.com/bloodmagesoftware/architect/internal/math/f32"
)

func TestFloor(t *testing.T) {
	a := float32(1.125)
	b := float32(1.0)

	if mathf32.Floor(a) != b {
		t.Errorf("Floor(%f) = %f; want %f", a, mathf32.Floor(a), b)
	}
}
