package report2

import "testing"

func TestCellLocation(t *testing.T) {
	var location CellLocation
	var expected, actual string

	location = CellLocation{Row: 1, Col: 1}
	expected = "A1"
	actual = location.String()
	if actual != expected {
		t.Errorf("expected [%v], actual [%v]", expected, actual)
	}

	location = CellLocation{Row: 20, Col: 80}
	expected = "CB20"
	actual = location.String()
	if actual != expected {
		t.Errorf("expected [%v], actual [%v]", expected, actual)
	}

}
