package cell

import (
	"fmt"
	"reflect"
)

type StringData struct {
	Display string
	Key     string
}

type StringDataType[T StringData] struct {
	rowIsHeader bool
	values      []StringData
}

// Parse handles string or []byte and returns a string interface from them
func (c *StringDataType[T]) Parse(v interface{}) (i interface{}, err error) {

	switch v.(type) {
	case StringData:
		if f := v.(StringData); len(f.Display) > 0 || len(f.Key) > 0 {
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
func (c *StringDataType[T]) Set(values ...interface{}) error {
	var err error
	for _, v := range values {
		if val, parseErr := c.Parse(v); parseErr == nil {
			c.values = append(c.values, val.(StringData))
		} else {
			err = fmt.Errorf("failed to convert [%v] to a StringData format", v)
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

// Get returns the first values Display value
func (c *StringDataType[T]) Get() (interface{}, error) {
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
