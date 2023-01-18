package cell

import "reflect"

type CellData interface {
	Parse(v interface{}) (interface{}, error)
	Set(values ...interface{}) error
	Get() (interface{}, error)
	GetAll() ([]interface{}, error)
	Type() reflect.Type
	TType() reflect.Type
	IsHeading() bool
}
