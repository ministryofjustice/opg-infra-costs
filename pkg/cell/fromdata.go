package cell

import (
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
	"opg-infra-costs/pkg/cell/formuladatatype"
)

// FromData inspects the data passed as an interface
// and uses IsA to determine which CellData struct is
// correct to return
// If nothing matches, then it returns nil & error
func FromData(data interface{}) (CellData, error) {
	var err error
	var c CellData

	// Most specific to least
	// -- formula as json
	// -- formula as FormulaDate
	// -- float
	// -- float as int
	// -- date
	if IsA[*formuladatatype.FormulaDataType[string]](data) {
		c = &formuladatatype.FormulaDataType[string]{}
	} else if IsA[*formuladatatype.FormulaDataType[formuladatatype.FormulaData]](data) {
		c = &formuladatatype.FormulaDataType[formuladatatype.FormulaData]{}
	} else if IsA[*floatdatatype.FloatDataType[float64]](data) {
		c = &floatdatatype.FloatDataType[float64]{}
	} else if IsA[*floatdatatype.FloatDataType[int]](data) {
		c = &floatdatatype.FloatDataType[int]{}
	} else if IsA[*datedatatype.DateDataType[string]](data) {
		c = &datedatatype.DateDataType[string]{}
	}

	return c, err

}

// IsA uses the Parse method on T to decide if
// data is of that type
func IsA[T CellData](data interface{}) bool {
	var t T
	if _, err := t.Parse(data); err == nil {
		return true
	}
	return false

}
