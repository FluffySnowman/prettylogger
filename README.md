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

Initialise the pretty logger with a config (can be SIMPLE, SIMPLE2 or TIMEBASED)
which is automatically used everywhere once set.

Basic example:

```go 
package main

import (
  pl "github.com/fluffysnowman/prettylogger"
)

func main() {
  pl.InitPrettyLogger("SIMPLE")       // basic 
  pl.LogInfo("Hello World")
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
  pl.Log("Hello World")
}
```

Using the new `SIMPLE2` format:

```go
package main

import (
  pl "github.com/fluffysnowman/prettylogger"
)

func main() {
  pl.InitPrettyLogger("SIMPLE2")    // detailed and concise format (new)
  pl.Log("Hello World")
}
```

> [!IMPORTANT]  
> Note that `SIMPLE`, `SIMPLE2` and `TIMEBASED` cannot be used together.
> Changing or setting the InitPrettyLogger() multiple times may cause problems.

Multiple arguments:

```go
pl.LogDebug("this is a debug log %v", "foo bar")
```

## All Logging Functions

Below is a list of all the available functions.

```go
pl.InitPrettyLogger(opts)  // Accepts "SIMPLE", "SIMPLE2" or "TIMEBASED"

pl.Log()          // green
pl.LogDebug()     // cyan 
pl.LogError()     // red
pl.LogInfo()      // cyan
pl.LogWarn()      // yellow
pl.LogFatal()     // red
pl.LogSuccess()   // green
pl.LogFailure()   // yellow
pl.LogOK()        // green
pl.LogErrorBG()   // red background, white text
pl.LogFailureBG() // yellow background, white text
```


