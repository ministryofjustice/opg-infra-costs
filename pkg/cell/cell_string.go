package cell

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StringData struct {
	Display string
	Key     string
}

type StringDataType[T string | StringData] struct {
	rowIsHeader bool
	values      []StringData
}

// validateStringData checks if the casted version of
// v has a Display or Key property set with a length > 0 (ie not empty)
// If its empty, an error is returned
func validateStringData(v interface{}) (interface{}, error) {
	if f := v.(StringData); len(f.Display) > 0 || len(f.Key) > 0 {
		return f, nil
	}
	return nil, fmt.Errorf("failed to validate [%v] as StringData", v)
}

func validateJsonAsStringData(v interface{}) (interface{}, error) {
	var obj StringData
	if err := json.Unmarshal(v.([]byte), &obj); err == nil {
		return validateFormulaData(obj)
	}
	return nil, fmt.Errorf("failed to validate as json [%v] to StringData", v)
}

// Parse handles string or []byte and returns a string interface from them
func (c *StringDataType[T]) Parse(v interface{}) (interface{}, error) {
	var i interface{}
	var err error

	switch v.(type) {
	case StringData:
		return validateStringData(v)
	case []byte:
		// check if this is a json string version of StringData
		i, err = validateJsonAsStringData(v.([]byte))
		if err == nil {
			return i, nil
		}
		// otherwise, presume a []byte to string
		i = string(v.([]byte))
		return i, nil
	case string:
		// check this string for being json, swap to its []byte form
		i, err = validateJsonAsStringData([]byte(v.(string)))
		if err == nil {
			return i, nil
		}
		// otherwise, presume string
		i = v.(string)
		return i, nil
	}

	err = fmt.Errorf("failed to parse [%v] to a date", v)
	return nil, err
}

// Set takes a series of interfaces, checks each one (via Parse) and
// those that work are added to the list of values.
// If a parse fails, a error is set and this will be returned at the
// end. This way all valid items are added, but it does mean the error
// message gets overwritten
func (c *StringDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			switch val.(type) {
			case StringData:
				c.values = append(c.values, val.(StringData))
			default:
				c.values = append(c.values, StringData{Display: val.(string)})
			}

		} else {
			err = fmt.Errorf("failed to convert [%v] to a string format", v)
		}
	}
	return err
}

// GetAll interates over all the current `.values` and appends
// them as an interface{} to a slice which is then returned
func (c *StringDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the first entry of `.values` only
func (c *StringDataType[T]) Get() (interface{}, error) {
	if len(c.values) > 0 {
		return c.values[0], nil
	}
	return nil, nil
}

// Return if the row is a header
func (c *StringDataType[T]) GetIsRowAHeader() bool {
	return c.rowIsHeader
}

// SetIsRowAHeader sets the flag
func (c *StringDataType[T]) SetIsRowAHeader(b bool) {
	c.rowIsHeader = b
}

// Type returns the full type of c, so should
// be a pointer like *StringDataType[string]
func (c *StringDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T, so for this
// struct it would be string
func (c *StringDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
