package cell

import (
	"fmt"
)

// FromData inspects the data passed as an interface
// and uses IsA to determine which CellInterface struct is
// correct to return
// If nothing matches, then it returns nil & error
func FromData(data interface{}) (c CellInterface, err error) {

	// Most specific to least
	// -- formula as json
	// -- formula as FormulaDate
	// -- string as ColumnHeadData
	// -- float
	// -- float as int
	// -- date
	if IsA[*FormulaDataType[string]](data) {
		c = &FormulaDataType[string]{}
	} else if IsA[*FormulaDataType[FormulaData]](data) {
		c = &FormulaDataType[FormulaData]{}
	} else if IsA[*ColumnHeadDataType[ColumnHeadData]](data) {
		c = &ColumnHeadDataType[ColumnHeadData]{}
	} else if IsA[*FloatDataType[float64]](data) {
		c = &FloatDataType[float64]{}
	} else if IsA[*FloatDataType[int]](data) {
		c = &FloatDataType[int]{}
	} else if IsA[*DateDataType[string]](data) {
		c = &DateDataType[string]{}
	} else {
		err = fmt.Errorf("failed to find matching cell data type for [%v]", data)
	}

	return

}

// IsA uses the Parse method on T to decide if
// data is of that type
func IsA[T CellInterface](data interface{}) bool {
	var t T
	if _, err := t.Parse(data); err == nil {
		return true
	}
	return false

}
