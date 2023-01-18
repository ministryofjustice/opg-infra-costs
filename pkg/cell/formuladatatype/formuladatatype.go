package formuladatatype

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// FormulaData struct handles how we return data for cells
// that contain formula data - a mix of label for a heading
// and formula for cell content
// .Formula is a string ready for string substitution
//   - =SUM(${col}${row}:${col}${row})
type FormulaData struct {
	Label   string
	Formula string // FormulaFormat contains
}

type FormulaDataType[T FormulaData | string | []byte] struct {
	values      []FormulaData
	rowIsHeader *bool
}

// validateFormulaData checks if the FormulaData casted version of
// v has a Label & Formula property set with a length > 0 (ie not empty)
// If its empty, an error is returned
func validateFormulaData(v interface{}) (interface{}, error) {
	if f := v.(FormulaData); len(f.Label) > 0 && len(f.Formula) > 0 {
		return f, nil
	}
	return nil, fmt.Errorf("failed to validate [%v] as FormulaData", v)
}

// Parse checks the v interface{} type to determine how to convert.
//   - For a string or []byte the data it tries to unmarshall, presuming
//     a json string version of the struct
//   - If the type matches the FormulaData struct then it checks to
//     ensure properties are set
//
// If a type match is not found, an error is returned
func (c *FormulaDataType[T]) Parse(v interface{}) (interface{}, error) {
	var obj FormulaData

	switch v.(type) {
	case FormulaData:
		return validateFormulaData(v)
	case []byte:
		// check if its a json encoded string
		if err := json.Unmarshal(v.([]byte), &obj); err == nil {
			return validateFormulaData(obj)
		}
	case string:
		// check if its a json encoded string
		if err := json.Unmarshal([]byte(v.(string)), &obj); err == nil {
			return validateFormulaData(obj)
		}
	}

	return nil, fmt.Errorf("failed to parse [%v] to FormulaData", v)
}

// Set takes the incoming values, checks that they Parse correctly
// and those are do are added (as FormulaData) to the `.values` slice.
// If any fail, they are skipped & the last failures error message is
// returned.
func (c *FormulaDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			c.values = append(c.values, val.(FormulaData))
		} else {
			err = fmt.Errorf("failed to convert [%v] to a FormulaData", v)
		}
	}
	return err
}

// GetAll iterates over all the current `.values` and appends
// them as an interface{} to a slice which is then returned
func (c *FormulaDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the first Formula value and returns it as an interface
func (c *FormulaDataType[T]) Get() (interface{}, error) {

	if len(c.values) > 0 {
		var i interface{}
		if c.GetIsRowAHeader() {
			i = c.values[0].Label
		} else {
			i = c.values[0].Formula
		}
		return i, nil
	}
	return nil, nil
}

// Return if the row is a header
func (c *FormulaDataType[T]) GetIsRowAHeader() bool {
	if c.rowIsHeader != nil {
		return *c.rowIsHeader
	}
	return false
}

// SetIsRowAHeader sets the flag
func (c *FormulaDataType[T]) SetIsRowAHeader(b *bool) {
	c.rowIsHeader = b
}

// Type returns the full type of c, so should
// be a pointer like *DateDataType[string]
func (c *FormulaDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T, so for this
// struct it would be string
func (c *FormulaDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
