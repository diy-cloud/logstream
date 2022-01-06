# recordable

## Stdout

### New

```go
func New(ctx context.Context, level loglevel.LogLevel, converter func(log.Log) string) log.Writable
```

`New` is constructor of `Stdout` instance.  
`ctx` is recomended to use `context.WithCancel` to cancel `Stdout` instance.  
`level` is criterion of this `log.Writable` instance.  
`converter` is converting `log.Log` to string when writing.  
if `converter` is `nil`, will write time formated RFC3339Nano and log's message.

### Write

```go
func (s *Stdout) Write(value log.Log) error
```

write `value` to `stdout` with `converter`.

## Nats

`DO NOT USE THIS`
