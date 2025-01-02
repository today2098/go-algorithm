# go-algorithm

[![Go](https://github.com/today2098/go-algorithm/actions/workflows/go.yml/badge.svg)](https://github.com/today2098/go-algorithm/actions/workflows/go.yml)
[![verify](https://github.com/today2098/go-algorithm/actions/workflows/verify.yml/badge.svg)](https://github.com/today2098/go-algorithm/actions/workflows/verify.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/today2098/go-algorithm)](https://goreportcard.com/report/github.com/today2098/go-algorithm)
[![Go Reference](https://pkg.go.dev/badge/github.com/today2098/go-algorithm.svg)](https://pkg.go.dev/github.com/today2098/go-algorithm)
[![GitHub Tag](https://img.shields.io/github/v/tag/today2098/go-algorithm)](https://github.com/today2098/go-algorithm/tags)

My library for competitive programming by Go.


## Installation

```bash
go get -u github.com/today2098/go-algorithm
```


## Descriptions

### Data structure

| Name                                    | Summary                                    | Doucument                             |
| --------------------------------------- | ------------------------------------------ | ------------------------------------- |
| [Stack](./algorithm/stack.go)           | A data structure about simple stack (LIFO) | [stack.md](./docs/stack.md)           |
| [Queue](./algorithm/queue.go)           | A data structure about simple queue (FIFO) | [queue.md](./docs/queue.md)           |
| [Deque](./algorithm/deque.go)           | Double-ended queue                         | -                                     |
| [BinaryHeap](./algorithm/binaryheap.go) | A data structure about priority queue      | [binaryheap.md](./docs/binaryheap.md) |
| [UnionFind](./algorithm/unionfind.go)   | Disjoint-set data structure                | [unionfind.md](./docs/unionfind.md)   |


### Utilities

| Name                | Summary                     | Doucument |
| ------------------- | --------------------------- | --------- |
| [IO](./utils/io.go) | A helper structure about IO | -         |
