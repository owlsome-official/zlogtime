# zlogtime

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

zlogtime is a logging library, which measure the elapsed time of services on-demand.

## Table of Contents

- [zlogtime](#zlogtime)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Signatures](#signatures)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Dependencies](#dependencies)
  - [Example Usage](#example-usage)

## Installation

```bash
  go get -u github.com/owlsome-official/zlogtime
```

## Signatures

```go
var timeTracker zlogtime.ZLogTime = zlogtime.New()
```

or with configuration

```go
var timeTrackerWithConfig zlogtime.ZLogTime = zlogtime.New(
  zlogtime.Config{
    LogLevel: zerolog.DebugLevel.String()
  }
)
```

## Examples

```go
// Step 1: Defined as global (tools) variable
var (
  timeTracker zlogtime.ZLogTime = zlogtime.New()
)

// Step 2: Call TimeTrack within a function
func FuncName() {
  // NOTE: Always used with "defer"
  defer timeTracker.TimeTrack("NAME", time.Now())
  ...
}
```

## Config

```go
type Config struct {

  // Optional. Default: false
  Hidden bool

  // Optional. Default: "info"
  LogLevel string

  // Optional. Default: "milli". Possible Value: ["nano", "micro", "milli"]
  ElapsedTimeUnit string
}
```

## Default Config

```go
var ConfigDefault = Config{
  Hidden:          false,
  LogLevel:        "info",
  ElapsedTimeUnit: "milli",
}
```

## Dependencies

- [Zerolog](https://github.com/rs/zerolog)

## Example Usage

Please go to [example/main.go](./example/main.go)

> Made with ❤️ by watsize
