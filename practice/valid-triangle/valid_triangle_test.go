package validtriangle

import (
	"reflect"
	"testing"
)

func TestValidTriangle(t *testing.T) {
	if !reflect.DeepEqual(is_valid_triangle([]int{9, 9, 9}), "Yes") {
		t.Errorf("expected Yes, got %v", is_valid_triangle([]int{9, 9, 9}))
	}
}
