package worksheet

import "testing"

func TestFilterConfigs(t *testing.T) {

	enabled := EnabledFilter{}
	ok, err := enabled.GetEnabled()

	if !ok {
		t.Errorf("expected [%T] to return true", enabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", enabled, err)
	}

	disabled := DisabledFilter{}
	ok, err = disabled.GetEnabled()

	if ok {
		t.Errorf("expected [%T] to return false", disabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", disabled, err)
	}

}

func TestPaneConfigs(t *testing.T) {

	enabled := EnabledPane{}
	ok, err := enabled.GetEnabled()

	if !ok {
		t.Errorf("expected [%T] to return true", enabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", enabled, err)
	}

	disabled := DisabledPane{}
	ok, err = disabled.GetEnabled()

	if ok {
		t.Errorf("expected [%T] to return false", disabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", disabled, err)
	}

}

func TestTableConfigs(t *testing.T) {

	enabled := EnabledTable{}
	ok, err := enabled.GetEnabled()

	if !ok {
		t.Errorf("expected [%T] to return true", enabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", enabled, err)
	}

	disabled := DisabledTable{}
	ok, err = disabled.GetEnabled()

	if ok {
		t.Errorf("expected [%T] to return false", disabled)
	}
	if err != nil {
		t.Errorf("unexpected error from [%T] [%v]", disabled, err)
	}

	if v, _ := enabled.GetStyleName(); v != standardStyle {
		t.Errorf("expected [%T] stylename to match standard [%v], actual [%v]", enabled, standardStyle, v)
	}
	if v, _ := disabled.GetStyleName(); v != standardStyle {
		t.Errorf("expected [%T] stylename to match standard [%v], actual [%v]", disabled, standardStyle, v)
	}

}
