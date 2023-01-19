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

	SetTableConfiguration(c SheetTableConfigurationInterface) error
	GetTableConfiguration() (SheetTableConfigurationInterface, error)

	SetFilterConfiguration(c SheetFilterConfigurationInterface) error
	GetFilterConfiguration() (SheetFilterConfigurationInterface, error)

	SetPaneConfiguration(c SheetPaneConfigurationInterface) error
	GetPaneConfiguration() (SheetPaneConfigurationInterface, error)
}

type SheetTableConfigurationInterface interface {
	SetEnabled(b bool) error
	GetEnabled() (bool, error)

	SetRange(min string, max string) error
	GetRange() (string, string, error)
}

type SheetFilterConfigurationInterface interface {
	SetEnabled(b bool) error
	GetEnabled() (bool, error)

	SetRange(min string, max string) error
	GetRange() (string, string, error)
}

type SheetPaneConfigurationInterface interface {
	SetEnabled(b bool) error
	GetEnabled() (bool, error)

	SetLocation(x int, y int) error
	GetLocation() (int, int, error)
}
