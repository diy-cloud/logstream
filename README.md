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
	"github.com/snowmerak/logstream/logqueue/logquebuf"
	"github.com/snowmerak/logstream/unlock/logringbuf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// using proirity queue for log buffer
	ls := logstream.New(ctx, logquebuf.New(8), 8)
	ls.ObserveTopic("A", recordable.NewStdout(ctx, loglevel.All, nil))

	for i := 0; i < 10; i++ {
		ls.Write("A", log.New(loglevel.Fatal, "qp time test log").End())
	}

	// using ringbuffer for log buffer
	ls = logstream.New(ctx, logringbuf.New(8), 8)
	ls.ObserveTopic("A", recordable.NewStdout(ctx, loglevel.All, nil))

	for i := 0; i < 10; i++ {
		ls.Write("A", log.New(loglevel.Fatal, "rb time test log").End())
	}

	time.Sleep(1 * time.Second)
}
```

```bash
2022-01-03T14:02:14.048117+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048119+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.04812+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048121+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048121+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048122+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048123+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048123+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048124+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048168+09:00 [FATAL] qp time test log
2022-01-03T14:02:14.048331+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048335+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048335+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048336+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048336+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048337+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048337+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048337+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048338+09:00 [FATAL] rb time test log
2022-01-03T14:02:14.048354+09:00 [FATAL] rb time test log
```
