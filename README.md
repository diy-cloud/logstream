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
2022-01-03T03:46:48+09:00 [DEBUG] test
2022-01-03T03:46:48+09:00 [WARN] test
2022-01-03T03:46:48+09:00 [ERROR] test
2022-01-03T03:46:48+09:00 [FATAL] test
2022-01-03T03:46:48+09:00 [WARN] test ? string=hello!
2022-01-03T03:46:48+09:00 [ERROR] test ? float=3.141592
2022-01-03T03:46:48+09:00 [DEBUG] test ? bool=true
2022-01-03T03:49:12+09:00 [FATAL] test ? int=-9999 uint=9999
```

## ctrie LICENSE

Copyright 2015 Workiva, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Package ctrie provides an implementation of the Ctrie data structure, which is
a concurrent, lock-free hash trie. This data structure was originally presented
in the paper Concurrent Tries with Efficient Non-Blocking Snapshots:
https://axel22.github.io/resources/docs/ctries-snapshot.pdf

`from github.com/Workiva/go-datastructures`
