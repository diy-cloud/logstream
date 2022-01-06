# logstream

logstream is simple library for logging

## install

`go get github.com/snowmerak/logstream`

## use

```go
package main

import (
	"context"
	"time"

	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/logbuffer/logqueue"
	"github.com/snowmerak/logstream/log/logbuffer/logring"
	"github.com/snowmerak/logstream/log/logbuffer/logstream/globalque"
	"github.com/snowmerak/logstream/log/loglevel"
	"github.com/snowmerak/logstream/log/writable/stdout"
)

func main() {
	const aTopic = "A"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// using proirity queue for log buffer
	ls := globalque.New(ctx, logqueue.New, 8)
	ls.ObserveTopic(aTopic, stdout.New(ctx, loglevel.All, nil))

	ls.Write(aTopic, log.New(loglevel.Debug, "a debug message").End())
	ls.Write(aTopic, log.New(loglevel.Info, "a info message").AddParamInt("int parameter", 99).End())

	// using ringbuffer for log buffer
	ls = globalque.New(ctx, logring.New, 8)
	ls.ObserveTopic("A", stdout.New(ctx, loglevel.All, nil))

	ls.Write(aTopic, log.New(loglevel.Fatal, "a fatal message").End())
	ls.Write(aTopic, log.New(loglevel.Error, "a error message").End())

	time.Sleep(1 * time.Second)
}
```

```bash
2022-01-06T20:44:58.971991+09:00 [DEBUG] a debug message
2022-01-06T20:44:58.972037+09:00 [INFO] a info message ? int parameter=99
2022-01-06T20:44:58.972074+09:00 [FATAL] a fatal message
2022-01-06T20:44:58.972075+09:00 [ERROR] a error message
```
