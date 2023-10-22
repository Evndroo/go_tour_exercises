package geometry

import "testing"

func TestGeometryFunc(t *testing.T) {
	areas := Geometry()

	if areas.circle != 78.53981633974483 {
		t.Errorf("Circle area was incorrect, got: %f, want: %f.", areas.circle, 78.53981633974483)
	}

	if areas.rectangle != 50 {
		t.Errorf("Rectangle area was incorrect, got: %f, want: %f.", areas.rectangle, 50.0)
	}
}
