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

	"github.com/snowmerak/logstream"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
	"github.com/snowmerak/logstream/log/recordable"
	"github.com/snowmerak/logstream/logqueue/logbuf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ls := logstream.New(ctx, logbuf.New(8), 8)
	ls.ObserveTopic("A", recordable.NewStdout(ctx, loglevel.All, nil))
	ls.ObserveTopic("B", recordable.NewStdout(ctx, loglevel.All, nil))

	ls.Write("A", log.New(loglevel.Fatal, "a fatal log").End())
	ls.Write("B", log.New(loglevel.Fatal, "b fatal log").End())
	ls.Write("A", log.New(loglevel.Error, "a error log").End())
	ls.Write("B", log.New(loglevel.Error, "b error log").End())
	ls.Write("A", log.New(loglevel.Warn, "a warn log").End())
	ls.Write("B", log.New(loglevel.Warn, "b warn log").End())
	ls.Write("A", log.New(loglevel.Info, "a info log").End())
	ls.Write("B", log.New(loglevel.Info, "b info log").End())
	ls.Write("A", log.New(loglevel.Debug, "a debug log").End())
	ls.Write("B", log.New(loglevel.Debug, "b debug log").End())
	time.Sleep(150 * time.Microsecond)
}
```

```bash
2022-01-03T13:40:13.707361+09:00 [FATAL] b fatal log
2022-01-03T13:40:13.707363+09:00 [ERROR] b error log
2022-01-03T13:40:13.707364+09:00 [WARN] b warn log
2022-01-03T13:40:13.707365+09:00 [INFO] b info log
2022-01-03T13:40:13.707371+09:00 [DEBUG] b debug log
2022-01-03T13:40:13.707358+09:00 [FATAL] a fatal log
2022-01-03T13:40:13.707362+09:00 [ERROR] a error log
2022-01-03T13:40:13.707363+09:00 [WARN] a warn log
2022-01-03T13:40:13.707364+09:00 [INFO] a info log
2022-01-03T13:40:13.70737+09:00 [DEBUG] a debug log
```
