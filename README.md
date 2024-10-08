# PL - Simple Pretty Logger

PL is a simple, easy to use pretty logger made in go.

<!--toc:start-->
- [Installation](#installation)
- [Usage](#usage)
- [All Logging Functions](#all-logging-functions)
<!--toc:end-->


## Installation

Use `go get` to install with the latest tag `v0.0.2` (recommended)

```bash
go get -u github.com/fluffysnowman/prettylogger@v0.0.2
```

or

Go's module suppert automatically fetches all the dependencies needed when you
import it in your code so `go get` isn't required, however for this project, the
`go get` installation is recommended with a specified tag (see the code block
above for instructions).

```go
// Importing and setting an alias to `pl` 
import pl "github.com/fluffysnowman/prettylogger"
```


## Usage

Initialise the pretty logger with a config (can be SIMPLE, TIMEBASED) which
is automatically used everywhere once set.

Basic example:

```go 
package main

import (
  pl "github.com/fluffysnowman/prettylogger"
)

func main() {
  pl.InitPrettyLogger("SIMPLE")       // basic 
  pl.Log("Hello World").Print()   // .Print() must always be called
                                      // otherwise nothing works
}
```

Using timestamps:

```go 
package main

import (
  pl "github.com/fluffysnowman/prettylogger"
)

func main() {
  pl.InitPrettyLogger("TIMEBASED")    // shows timestamps
  pl.Log("Hello World").Print()
}
```

> Please note that `SIMPLE` and `TIMEBASED` cannot be used together. Changing or
> redoing the InitPrettyLogger() will cause problems. 

Multiple arguments:

```go
pl.LogDebug("this is a debug log %v", "foo bar").Print()
```

Force a timestamp log:

This is useful when you want to log with timestamps when `InitPrettyLogger()` is
set with `SIMPLE`

```go
pl.LogDebug("this is a debug log %v", "with a timestamp").Timestamp().Print()
```

Changing SIMPLE to TIMEBASED & vice versa will cause problems hence use the
.Timestamp() and .Print() when timestamps are needed.

## All Logging Functions

Below is a list of all the available functions.

> Please note that `.Print()` **must be chained after every function**. The logs
> will not be printed without .Print() after each one. 

> For timestamp logs, `.Print()` should be called **AFTER** `.Timestamp()` for
> timestamp logs; i.e. `pl.LogDebug().Timestamp().Print()`

```go
pl.InitPrettyLogger(opts)  // Accepts "SIMPLE" or "TIMEBASED"

// Chaining functions (in order)
.Timestamp()  // adds timestamp to the log
.Print()      // prints the log (must be chained to each log for it to work)

// All functions below accept `.Timestamp()` before `.Print()`
pl.Log().Print()          // green
pl.LogDebug().Print()     // cyan 
pl.LogError().Print()     // red
pl.LogInfo().Print()      // cyan
pl.LogFatal().Print()     // red
pl.LogSuccess().Print()   // green
pl.LogFailure().Print()   // yellow
pl.LogOK().Print()        // green
pl.LogErrorBG().Print()   // red background, white text
pl.LogFailureBG().Print() // yellow background, white text
```


