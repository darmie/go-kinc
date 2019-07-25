package kinc

const (
	Info    LogLevel = 0
	Warning LogLevel = 1
	Error   LogLevel = 2
)

type LogLevel int



func Log(level LogLevel, format string, interface{}...) {
 // Todo: use normal Golang logger
}

