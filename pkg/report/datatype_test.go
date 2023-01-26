package report

import "testing"

func TestDataType(t *testing.T) {
	var dt ValueDataType
	var err error

	valid := map[string]ValueDataType{
		"test":    DataIsAString,
		"123":     DataIsANumber,
		"123.01":  DataIsANumber,
		"=SUM()":  DataIsAFormula,
		"bool":    DataIsAString,
		"2022-01": DataIsADate,
	}

	for test, tt := range valid {
		dt, err = DataType(test)
		if err != nil {
			t.Errorf("unexpected err [%v]", err)
		}
		if dt != tt {
			t.Errorf("expected [%v] type [%v], actual [%v]", test, tt, dt)
		}

	}

}
