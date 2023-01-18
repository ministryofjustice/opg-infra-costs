package floatdatatype

import (
	"fmt"
	"reflect"
	"strconv"
)

type FloatDataType[T float64 | float32 | int] struct {
	rowIsHeader *bool
	values      []float64
}

// Parse checks the v.(type) of the interface to determine how to
// convert the data in to a float. By default it assumes a string.
// If parsing fails nil and an error are returned
func (c *FloatDataType[T]) Parse(v interface{}) (interface{}, error) {
	var val float64
	var err error

	switch v.(type) {
	case int:
		val = float64(v.(int))
	case float32:
		val = float64(v.(float32))
	case float64:
		val = v.(float64)
	default:
		val, err = strconv.ParseFloat(v.(string), 64)
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
func (c *FloatDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			c.values = append(c.values, val.(float64))
		} else {
			err = fmt.Errorf("failed to convert [%v] to a date format", v)
		}
	}
	return err
}

// GetAll converts the `.values` into a slice of interfaces
// and returns that
func (c *FloatDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the total value of all stored `.values`
func (c *FloatDataType[T]) Get() (interface{}, error) {
	sum := 0.0
	for _, v := range c.values {
		sum += v
	}
	return sum, nil
}

// Return if the row is a header
func (c *FloatDataType[T]) IsHeading() bool {
	if c.rowIsHeader != nil {
		return *c.rowIsHeader
	}
	return false
}

// SetIsHeading sets the flag
func (c *FloatDataType[T]) SetIsHeading(b *bool) {
	c.rowIsHeader = b
}

// Type returns this cells structure type, so you should get a pointer
// like *FloatDataType[float64]
func (c *FloatDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T for this cell, so
// should return float64
func (c *FloatDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
