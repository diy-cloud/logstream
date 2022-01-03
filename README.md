# logstream

logstream is simple library for logging

## install

`go get github.com/snowmerak/logstream`

## use

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/snowmerak/logstream"
	"github.com/snowmerak/logstream/buf/logbuf"
	"github.com/snowmerak/logstream/log"
	"github.com/snowmerak/logstream/log/loglevel"
	"github.com/snowmerak/logstream/log/recordable"
)

func main() {
	ls := logstream.New(context.Background(), logbuf.New(8))
	if err := ls.Observe("test", recordable.NewStdout(loglevel.All, true, nil)); err != nil {
		fmt.Println(err)
	}
	ls.Write("test", log.New(loglevel.Debug, "test").End())
	ls.Write("test", log.New(loglevel.Error, "test").End())
	ls.Write("test", log.New(loglevel.Fatal, "test").End())
	ls.Write("test", log.New(loglevel.Warn, "test").End())
	ls.Write("test", log.New(loglevel.Warn, "test").AddParamString("string", "hello!").End())
	ls.Write("test", log.New(loglevel.Fatal, "test").AddParamInt("int", -9999).AddParamUint("uint", 9999).End())
	ls.Write("test", log.New(loglevel.Error, "test").AddParamFloat("float", 3.141592).End())
	ls.Write("test", log.New(loglevel.Debug, "test").AddParamBool("bool", true).End())
	time.Sleep(1 * time.Second)
}
```

```bash
2022-01-03T04:17:39.292867+09:00 [DEBUG] test
2022-01-03T04:17:39.292871+09:00 [ERROR] test
2022-01-03T04:17:39.292872+09:00 [FATAL] test
2022-01-03T04:17:39.292873+09:00 [WARN] test
2022-01-03T04:17:39.292873+09:00 [WARN] test ? string=hello!
2022-01-03T04:17:39.292874+09:00 [FATAL] test ? int=-9999 uint=9999
2022-01-03T04:17:39.292879+09:00 [ERROR] test ? float=3.141592
2022-01-03T04:17:39.292892+09:00 [DEBUG] test ? bool=true
```
