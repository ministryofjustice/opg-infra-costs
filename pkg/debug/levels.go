package debug

type DebugLevel int

var LEVEL DebugLevel = ALL

const (
	ALL     DebugLevel = 50000
	INFO    DebugLevel = 3
	LIMITED DebugLevel = 2
	NONE    DebugLevel = 0
)
