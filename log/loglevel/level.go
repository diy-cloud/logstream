package loglevel

const (
	All = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Off
)

func WrapColor(level int32, message string) string {
	switch level {
	case Debug:
		return "\033[0;36m" + message + "\033[0m"
	case Info:
		return "\033[0;32m" + message + "\033[0m"
	case Warn:
		return "\033[0;33m" + message + "\033[0m"
	case Error:
		return "\033[0;31m" + message + "\033[0m"
	case Fatal:
		return "\033[0;35m" + message + "\033[0m"
	default:
		return "\033[0;37m" + message + "\033[0m"
	}
}

func Available(criterion, loglevel int32) bool {
	return criterion <= loglevel
}
