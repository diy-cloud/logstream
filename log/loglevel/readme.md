# loglevel

a enum package of loglevel

## levels

there are 7 levels.

```go
const (
	All LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Off
)
```

All is contains all.  
Debug is contains itself, Info, Warn, Error, Fatal.  
...  
Fatal is contains itself.  
Off has nothing.  

## functions

### WrapColor

```go
func WrapColor(level LogLevel, message string) string
```

`WrapColor` is paint color to message.  
if `Fatal` is violet.  
if `Error` is red.  
if `Warn` is yellow.  
if `Info` is green.  
if `Debug` is blue.  

### Avaliable

```go
func Available(criterion, loglevel LogLevel) bool {
	return criterion <= loglevel
}
```

check `loglevel` is available or not.
