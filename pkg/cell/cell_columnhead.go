package cell

import (
	"fmt"
	"reflect"
)

type ColumnHeadData struct {
	Display string
	Key     string
}

type ColumnHeadDataType[T ColumnHeadData] struct {
	rowIsHeader bool
	values      []ColumnHeadData
}

// Parse handles string or []byte and returns a string interface from them
func (c *ColumnHeadDataType[T]) Parse(v interface{}) (i interface{}, err error) {

	switch v.(type) {
	case ColumnHeadData:
		if f := v.(ColumnHeadData); len(f.Display) > 0 || len(f.Key) > 0 {
			i = f
		}
	}

	if i == nil {
		err = fmt.Errorf("failed to parse [%v] to a stringdata", v)
	}
	return i, err
}

// Set takes a series of interfaces, checks each one (via Parse) and
// those that work are added to the list of values.
// If a parse fails, a error is set and this will be returned at the
// end. This way all valid items are added, but it does mean the error
// message gets overwritten
func (c *ColumnHeadDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			c.values = append(c.values, val.(ColumnHeadData))
		} else {
			err = fmt.Errorf("failed to convert [%v] to a ColumnHeadData format", v)
		}
	}
	return err
}

// GetAll interates over all the current `.values` and appends
// them as an interface{} to a slice which is then returned
func (c *ColumnHeadDataType[T]) GetAll() ([]interface{}, error) {
	interfaces := []interface{}{}
	for _, v := range c.values {
		interfaces = append(interfaces, v)
	}
	return interfaces, nil
}

// Get returns the first values Display value
func (c *ColumnHeadDataType[T]) Get() (interface{}, error) {
	var i interface{}

	if len(c.values) > 0 {
		first := c.values[0]
		if len(first.Display) > 0 {
			return first.Display, nil
		} else {
			return first.Key, nil
		}
	}

	return i, nil
}

// Return if the row is a header
func (c *ColumnHeadDataType[T]) GetIsRowAHeader() bool {
	return c.rowIsHeader
}

// SetIsRowAHeader sets the flag
func (c *ColumnHeadDataType[T]) SetIsRowAHeader(b bool) {
	c.rowIsHeader = b
}

// Type returns the full type of c, so should
// be a pointer like *ColumnHeadDataType[string]
func (c *ColumnHeadDataType[T]) Type() reflect.Type {
	return reflect.TypeOf(c)
}

// TType returns just the type of T, so for this
// struct it would be string
func (c *ColumnHeadDataType[T]) TType() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func (c *ColumnHeadDataType[T]) AcceptsValuesOf(t reflect.Type) bool {
	return t == c.TType()
}
