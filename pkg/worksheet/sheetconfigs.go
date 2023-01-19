package worksheet

// -- SheetFilterConfigurationInterface
type EnabledFilter struct{}
type DisabledFilter struct{}

func (f *EnabledFilter) GetEnabled() (bool, error) {
	return true, nil
}

func (f *DisabledFilter) GetEnabled() (bool, error) {
	return false, nil
}

// -- SheetPaneConfigurationInterface
type EnabledPane struct{}
type DisabledPane struct{}

func (p *EnabledPane) GetEnabled() (bool, error) {
	return true, nil
}

func (p *DisabledPane) GetEnabled() (bool, error) {
	return false, nil
}

// -- SheetTableConfigurationInterface
var standardStyle string = "TableStyleMedium9"

type EnabledTable struct{}
type DisabledTable struct{}

func (t *EnabledTable) GetEnabled() (bool, error) {
	return true, nil
}
func (t *EnabledTable) GetStyleName() (string, error) {
	return standardStyle, nil
}

func (t *DisabledTable) GetEnabled() (bool, error) {
	return false, nil
}
func (t *DisabledTable) GetStyleName() (string, error) {
	return standardStyle, nil
}
