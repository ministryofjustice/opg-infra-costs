package cell

import "reflect"

type CellData interface {
	Parse(v interface{}) (interface{}, error)

	Set(values ...interface{}) error
	Get() (interface{}, error)
	GetAll() ([]interface{}, error)

	SetIsRowAHeader(b *bool)
	GetIsRowAHeader() bool

	Type() reflect.Type
	TType() reflect.Type
}
