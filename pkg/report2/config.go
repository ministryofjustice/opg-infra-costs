package report2

import (
	"io/ioutil"
	"opg-infra-costs/pkg/debugger"
	"os"

	"gopkg.in/yaml.v2"
)

// -- Errors
const (
	ConfigurationReportHasNoColumns       string = "no columns data for this report"
	ConfigurationColumnDefinitionNotFound string = "cannot find a defintition for [%v]"
)

// -- Yaml structures
type Configuration struct {
	Reports             map[string]ReportConfiguration `yaml:"reports"`
	ColumnDefinitions   map[string]ColumnDefinition    `yaml:"column_definitions"`
	ExtraRowDefinitions map[string]ExtraRowDefinition  `yaml:"extra_row_definitions"`
}

// -- New / Creation functions

// unmarshalConfig handles converting the yaml to a struct
func unmarshalConfig(content []byte) (cfg Configuration, err error) {
	defer debugger.Log("unmarshalConfig.", debugger.VVERBOSE)()
	cfg = Configuration{}
	err = yaml.Unmarshal(content, &cfg)
	return
}

// NewConfiguration loads the config for the report from the file passed
func NewConfiguration(
	file string,
) (cfg Configuration, err error) {
	defer debugger.Log("Configuration loaded.", debugger.DETAILED)()

	if f, err := os.Open(file); err == nil {
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		cfg, err = unmarshalConfig(content)
	}

	return
}
