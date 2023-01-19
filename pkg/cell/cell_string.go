package cell

import (
	"fmt"
	"reflect"
)

type StringDataType[T string] struct {
	rowIsHeader bool
	values      []string
}

// Parse checks the v.(type) of the interface
func (c *StringDataType[T]) Parse(v interface{}) (interface{}, error) {
	var val string
	var err error

	switch v.(type) {
	case string:
		val = v.(string)
	default:
		err = fmt.Errorf("failed to parse [%v] to a string", v)
	}

	if err != nil {
		return nil, err
	}
	return val, err

}

// Set pushes items from values param into the `.values` if they
// pass the Parse check.
// Ones that fail are skipped and a single error returned. The
// error will be the last failed Parse
func (c *StringDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			c.values = append(c.values, val.(string))
		} else {
			err = fmt.Errorf("failed to convert [%v] to a date format", v)
		}
	}
	return err
}

// GetAll converts the `.values` into a slice of interfaces
// and returns that
func (c *StringDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the total value of all stored `.values`
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

// Type returns this cells structure type, so you should get a pointer
// like *StringDataType[string]
func (c *StringDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T for this cell, so
// should return float64
func (c *StringDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func (c *StringDataType[T]) AcceptsValuesOf(t reflect.Type) bool {
	return t == c.TType()
}
