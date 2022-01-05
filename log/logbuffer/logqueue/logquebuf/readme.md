# logbuf

## LogBuffer

```go
type LogBuffer struct {
	trie       *ctrie.Ctrie
	bufferSize int
	signals    map[string]chan struct{}
}
```

`LogBuffer` has a `ctrie.Ctrie` instance for topic.  
And `signals` is a map of channel for topic.  
Finally, `bufferSize` is a size of queue for buffer.

## New

```go
func New(bufferSize int) *LogBuffer
```

`New` constructor is a constructor of `LogBuffer` with given buffer size.  

## AddTopic

```go
func (e *LogBuffer) AddTopic(topic string, signal chan struct{})
```

`AddTopic` add a topic to `LogBuffer` instance and initialize signal.

## RemoveTopic

```go
func (e *LogBuffer) RemoveTopic(topic string)
```

`RemoveTopic` remove a topic from `LogBuffer` instance.  
DO NOT CALL THIS MENIALY.

## Enqueue

```go
func (e *LogBuffer) EnQueue(topic string, value log.Log)
```

`Enqueue` method is push `value` to `topic` queue in `LogBuffer`.

## Dequeue

```go
func (e *LogBuffer) DeQueue(topic string) (log.Log, error)
```

`Dequeue` method is pop `value` from `topic` queue in `LogBuffer`.

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