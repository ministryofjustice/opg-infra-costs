package cell

import (
	"fmt"
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
	"opg-infra-costs/pkg/cell/formuladatatype"
	"opg-infra-costs/pkg/cell/stringdatatype"
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
	// -- string
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
	} else if IsA[*stringdatatype.StringDataType[string]](data) {
		c = &stringdatatype.StringDataType[string]{}
	} else {
		err = fmt.Errorf("failed to find matching cell data type for [%v]", data)
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
