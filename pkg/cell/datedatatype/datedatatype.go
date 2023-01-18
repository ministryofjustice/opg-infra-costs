package datedatatype

import (
	"fmt"
	"opg-infra-costs/pkg/formats"
	"reflect"
	"time"
)

type DateDataType[T string] struct{ values []string }

var dateFormat = formats.DATES["ym"]

// Parse converts v.(string) with time.Parse and `dateFormat` into a
// date. This date is then stored as a string.
// If string -> time conversion fails an error is returned with an empty
// string
func (c *DateDataType[T]) Parse(v interface{}) (interface{}, error) {
	if val, err := time.Parse(dateFormat, v.(string)); err == nil {
		return val.Format(dateFormat), nil
	}
	return nil, fmt.Errorf("failed to parse [%v] to a date", v)
}

// Set takes a series of interfaces, checks each one (via Parse) and
// those that work are added to the list of values.
// If a parse fails, a error is set and this will be returned at the
// end. This way all valid items are added, but it does mean the error
// message gets overwritten
func (c *DateDataType[T]) Set(values ...interface{}) error {
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

// GetAll interates over all the current `.values` and appends
// them as an interface{} to a slice which is then returned
func (c *DateDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the first entry of `.values` only
func (c *DateDataType[T]) Get() (interface{}, error) {
	if len(c.values) > 0 {
		return c.values[0], nil
	}
	return "", nil
}

// Type returns the full type of c, so should
// be a pointer like *DateDataType[string]
func (c *DateDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T, so for this
// struct it would be string
func (c *DateDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
