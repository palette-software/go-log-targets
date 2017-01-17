[![Build Status](https://travis-ci.org/palette-software/go-log-targets.svg?branch=master)](https://travis-ci.org/palette-software/go-log-targets)

# go-log-targets
Logging utility that supports multiple log targets with different levels. The aim of this module to log records simulteanausly on various log targets. Log targets may have different logging levels.

# How to use go-log-targets
Anything can be a log target that implements the `io.Writer` interface.

For example, this is how you can set up a console and file log target:
```golang
import log "github.com/palette-software/go-log-targets/logging"

func main() {
  ...

  // Create console logging with Debug level
  log.AddTarget(os.Stdout, log.LevelDebug)

  logFile, err := os.OpenFile("judgement_day.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    fmt.Println("Failed to open log file! ", err)
    panic(err)
  }

  // Close log file on exit and check for its returned error
  defer func() {
    if err := logFile.Close(); err != nil {
      fmt.Println("Failed to close log file! ", err)
      panic(err)
    }
  }()

  // Create log file with Info level
  log.AddTarget(logFile, log.LevelInfo)

  log.Info("Skynet is booting") 
  
  ...
}
```

## Splunk log target via HTTP(S)
In this repo you can find a [Splunk] log target. Although there are several implementations out there with which you can log into [Splunk], but we haven't found anything which does it over HTTP(S), thus we implemented one for ourselves. Here you can read a bit about [Splunk HTTP Event Collector](http://dev.splunk.com/view/event-collector/SP-CAAAE7F).

Here is how you can add a Splunk logger to your logs:
```golang
splunkLogger, err := log.NewSplunkTarget(SplunkServerAddress, SplunkHTTPEventCollectorToken, owner)
if err != nil {
  defer splunkLogger.Close()
  log.AddTarget(splunkLogger, log.LevelWarning)
}
```

## Contribution

### Building locally

```
go get ./...
go build -v
```

### Testing

```
go get -t ./...
go test ./... -v
```

[Splunk]: https://www.splunk.com/
