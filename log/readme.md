# log

a package for Log structure.

## Log

```go
type Log struct {
	Message string
	Level   loglevel.LogLevel
	Time    time.Time
}
```

`Log` has `Message`, `Level`, and `Time` members.  
[loglevel](loglevel/readme.md) is a package for log level.  

## LogFactory

```go
type LogFactory struct {
	Time    time.Time
	Message strings.Builder
	Level   loglevel.LogLevel

	hasParam bool
}
```

`Log` is made by `LogFactory`.

## New

```go
func New(level loglevel.LogLevel, message string) *LogFactory
```

`New` constructor return `LogFactory` instance with `level` and `message`.

## AddParam

```go
func (l *LogFactory) AddParamString(key string, value string) *LogFactory

func (l *LogFactory) AddParamInt(key string, value int) *LogFactory

func (l *LogFactory) AddParamUint(key string, value uint) *LogFactory

func (l *LogFactory) AddParamBool(key string, value bool) *LogFactory

func (l *LogFactory) AddParamFloat(key string, value float64) *LogFactory

func (l *LogFactory) AddParamComplex(key string, value complex128) *LogFactory
```

`LogFactory` can receive some primitive type parameters.  
this methods concatenate `key` and `value` to `message` and return `LogFactory` instance.

## End

```go
func (l *LogFactory) End() Log
```

finally, `End` method returns `Log` instance.

## Compare

```go
func (l Log) Compare(other queue.Item) int
```

`Compare` is implementation of `Comparable` interface.  
if `l.Time` is earlier than `other.Time`, return -1.  
if `l.Time` is later than `other.Time`, return 1.  
else, return 0.  
