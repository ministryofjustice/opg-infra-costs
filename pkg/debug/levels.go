package debug

type DebugLevel int

var LEVEL DebugLevel = ALL
var DEPTH int = 0

const (
	ALL     DebugLevel = 50000
	INFO    DebugLevel = 40
	LIMITED DebugLevel = 30
	NONE    DebugLevel = 0
)
