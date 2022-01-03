# queue

this is wrapper of priority queue by `github.com/Workiva/go-datastructures/queue`

## New

```go
func New(size int) *LogQueue
```

create new `LogQueue` instance with given size.

## Push

```go
func (lq *LogQueue) Push(log log.Log) error
```

push `log.Log` to `LogQueue`.

## Pop

```go
func (lq *LogQueue) Pop() (log.Log, error)
```

pop `log.Log` from `LogQueue`.  
if `LogQueue` is empty, do not return error.  
if you receive error, you have some problems with `LogQueue` or computer.

## priority queue LICENSE

Copyright 2014 Workiva, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.