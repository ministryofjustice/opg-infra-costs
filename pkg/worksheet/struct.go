package worksheet

import (
	"opg-infra-costs/pkg/cell"
)

type SheetInterface interface {
	SetName(n string) error
	GetName() (string, error)
	SetIndex(i int) error
	GetIndex() (int, error)
	GetVisible() (bool, error)
	GetActive() (bool, error)
	GetTableConfiguration() (SheetTableConfigurationInterface, error)
	GetFilterConfiguration() (SheetFilterConfigurationInterface, error)
	GetPaneConfiguration() (SheetPaneConfigurationInterface, error)
	GetRowCount() (int, error)
	GetDefaultColumns() []cell.ColumnHeadDataType[cell.ColumnHeadData]
	SetColumns(columns ...cell.ColumnHeadDataType[cell.ColumnHeadData]) error
	GetColumns() ([]cell.ColumnHeadDataType[cell.ColumnHeadData], error)

	SetRawData(rows []map[string]string) error
	GetRawData() ([]map[string]string, error)

	ConvertData() error
}

type SheetTableConfigurationInterface interface {
	GetEnabled() (bool, error)
	GetStyleName() (string, error)
}

type SheetFilterConfigurationInterface interface {
	GetEnabled() (bool, error)
}

type SheetPaneConfigurationInterface interface {
	GetEnabled() (bool, error)
}
