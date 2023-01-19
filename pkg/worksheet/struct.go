package worksheet

type SheetInterface interface {
	SetName(n string) error
	GetName() (string, error)

	SetIndex(i int) error
	GetIndex() (int, error)

	SetVisible(b bool) error
	GetVisible() (bool, error)

	SetActive(b bool) error
	GetActive() (bool, error)

	GetTableConfiguration() (SheetTableConfigurationInterface, error)
	GetFilterConfiguration() (SheetFilterConfigurationInterface, error)
	GetPaneConfiguration() (SheetPaneConfigurationInterface, error)
}

type SheetTableConfigurationInterface interface {
	GetEnabled() (bool, error)
	GetRange() (string, string, error)
}

type SheetFilterConfigurationInterface interface {
	GetEnabled() (bool, error)
	GetRange() (string, string, error)
}

type SheetPaneConfigurationInterface interface {
	GetEnabled() (bool, error)
	GetLocation() (int, int, error)
}
