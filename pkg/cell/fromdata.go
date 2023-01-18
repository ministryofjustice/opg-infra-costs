package cell

import (
	"fmt"
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
)

// FromData inspects the data passed as an interface
// using the various datatypes using each `.Parse`
// function
// If nothing matches, then it returns nil & error
func FromData(data interface{}) (CellData, error) {
	var err error
	var c CellData

	dateType := &datedatatype.DateDataType[string]{}
	floatType := &floatdatatype.FloatDataType[float64]{}

	if _, err = floatType.Parse(data); err == nil {
		c = floatType
	} else if _, err = dateType.Parse(data); err == nil {
		c = dateType
	} else {
		c = nil
		err = fmt.Errorf("could not match data to cell type")
	}

	return c, err

}
