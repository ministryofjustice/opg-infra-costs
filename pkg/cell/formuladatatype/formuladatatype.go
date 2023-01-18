package formuladatatype

import (
	"encoding/json"
	"fmt"
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

type FormulaDataType[T FormulaData | string | []byte] struct{ values []string }

// Parse checks the string if its FormulaData struct or a Json string
func (c *FormulaDataType[T]) Parse(v interface{}) (interface{}, error) {

	switch v.(type) {
	case FormulaData:
		if f := v.(FormulaData); len(f.Label) > 0 && len(f.Formula) > 0 {
			return f, nil
		}
	case []byte:
		// check if its a json encoded string
		var obj FormulaData
		if err := json.Unmarshal(v.([]byte), &obj); err == nil {
			return obj, nil
		}
	case string:
		// check if its a json encoded string
		var obj FormulaData
		if err := json.Unmarshal([]byte(v.(string)), &obj); err == nil {
			return obj, nil
		}
	}

	return nil, fmt.Errorf("failed to parse [%v] to FormulaData", v)
}

// func (c *FormulaDataType[T]) Set(values ...interface{}) error {
// 	var err error
// 	// for _, v := range values {
// 	// 	if val, parseErr := c.Parse(v); parseErr == nil {
// 	// 		c.values = append(c.values, val.(string))
// 	// 	} else {
// 	// 		err = fmt.Errorf("failed to convert [%v] to a date format", v)
// 	// 	}
// 	// }
// 	return err
// }

// // GetAll iterates over all the current `.values` and appends
// // them as an interface{} to a slice which is then returned
// func (c *FormulaDataType[T]) GetAll() ([]interface{}, error) {
// 	interfaces := []interface{}{}
// 	for _, v := range c.values {
// 		interfaces = append(interfaces, v)
// 	}
// 	return interfaces, nil
// }

// // Get returns the first entry of `.values` only
// func (c *FormulaDataType[T]) Get() (interface{}, error) {
// 	if len(c.values) > 0 {
// 		return c.values[0], nil
// 	}
// 	return "", nil
// }
