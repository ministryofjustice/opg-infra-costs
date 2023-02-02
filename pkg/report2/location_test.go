package report2

import "testing"

func TestLocation(t *testing.T) {
	defer _reset()
	var location Location
	var expected, actual string

	location = Location{Row: 1, Col: 1, Sheet: "tt"}
	expected = "A1"
	actual = location.String()
	if actual != expected {
		t.Errorf("expected [%v], actual [%v]", expected, actual)
	}

	location = Location{Row: 20, Col: 80, Sheet: "tt"}
	expected = "CB20"
	actual = location.String()
	if actual != expected {
		t.Errorf("expected [%v], actual [%v]", expected, actual)
	}

}
