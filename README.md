# error-stream

error-stream is simple tool for global error handling with topic.

## install

`go get github.com/snowmerak/error-stream`

## use

```go
package main

import (
    "errors"

    "github.com/snowmerak/error-stream"
)

func main() {
    const normalStream = "normal"

    es := errorstream.New(8)
    es.EnQueue(normalStream, errors.New("new error"))

    err, b := es.DeQueue(normalStream)
    if !b {
        panic("not exist topic")
    }

    println(err)
}
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