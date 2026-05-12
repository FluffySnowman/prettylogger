# PL - Simple Pretty Logger

PL is a simple, easy to use pretty logger made in go.

> [!NOTE]  
> The latest version produces the same logging output as
> [elog](https://github.com/FluffySnowman/elog) along with the same available
> config/string options.


<!--toc:start-->
- [PL - Simple Pretty Logger](#pl-simple-pretty-logger)
  - [Installation](#installation)
  - [Usage (latest)](#usage-latest)
    - [Debug log configuration](#debug-log-configuration)
      - [Programmatic API for debug log config](#programmatic-api-for-debug-log-config)
      - [Notes about how things work and how they are handled in go](#notes-about-how-things-work-and-how-they-are-handled-in-go)
    - [All available log types/methods/funcs](#all-available-log-typesmethodsfuncs)
  - [Usage (old)](#usage-old)
  - [All Logging Functions](#all-logging-functions)
<!--toc:end-->


## Installation

Below are the instructions for `v0.1.2` and `v0.1.5`. `v0.1.5` is recommended.

<details open>
  <summary>Latest V3 Documentation</summary>
  Use `go get` to install with the latest tag `v0.1.5` (recommended)

  ```bash
  go get -u github.com/fluffysnowman/prettylogger@v0.1.5
  ```

  Import: 

  ```go
  import pl "github.com/fluffysnowman/prettylogger"
  ```

  ## Usage (latest)

  Specify options on init seperated by commas. 

  All available config/string options (taken from ./pl.go package doc comment): 

  - TEA           : british time format (dd-mm-yy hh:mm:ss.mmm)
  - BURGER        : american time format (mm-dd-yy hh:mm:ss.mmm)
  - READABLE      : readable time format (ddmonyyyy hh:mm ss.mmm s)
  - FILECOLOR     : write colour logs to file
  - FILEONLY      : write logs only to file, no stdout output etc.  ..
  - FILE=<path>   : write logs to specified file path, if exists, if not- created


  Example usage:

  ```go 
  pl.InitPrettyLogger("V3")                                       // original (just times, no dates) time format
  pl.InitPrettyLogger("V3,TEA")                                   // TEA (bri'ish) time format
  pl.InitPrettyLogger("V3,BURGER,FILECOLOR")                      // BURGER (american) time + colored file
  pl.InitPrettyLogger("V3,READABLE,FILEONLY,FILE=logfile.log")    // readable time w/ month names etc, file-only written to logfile.log
  ```

  ### Debug log configuration 

  You can turn **all `LogDebug` output** off in production without changing all
  function calls throughout your codebase.

  This can be done simply via environment variables when running your program or
  programatically via the api.

  | Variable | Effect |
  |----------|--------|
  | **`PLOG_DEBUG`** | If set and NON-EMPTY (after trimming): **`1`** -> debug **on**; **`0`** -> debug **off**. ANY OTHER NON-EMPTY VALUE is treated as **on** |
  | **`LOGGER_ENV`** | Used only when **`PLOG_DEBUG`** is UNSET OR EMPTY. Value is trimmed and compared (case insensitieve) (e.g. **`prod`** and **`PROD`** both work). **`PROD`** -> debug **off**; anything else (INCLUDING UNSET) -> **on**. |

  **Precedence:** non-empty **`PLOG_DEBUG`** **wins over** **`LOGGER_ENV`**.

  Examples: 

  ```bash 
  export LOGGER_ENV=PROD          # debug OFF (if PLOG_DEBUG not set)
  export PLOG_DEBUG=0             # debug OFF (overrides LOGGER_ENV)
  export PLOG_DEBUG=1             # debug ON
  ```

  Above ^ can also be used inline when running as `PLOG_DEBUG=0 go run main.go
  (or ./main)` or `LOGGER_ENV=PROD go run main.go (or ./main)` etc.

  #### Programmatic API for debug log config

  ```go
  pl.DisableDebugLogs()   // force debug OFF
  pl.EnableDebugLogs()    // force debug ON
  if pl.DebugEnabled() {  // optional if needed. (should avoiid, possibly expensive)
    pl.LogDebug("state=%+v", expensiveStruct)
  }
  ```

  #### Notes about how things work and how they are handled in go

  For re-initialisation Each call to InitPrettyLogger(...) **re applies** the
  current process environment to the debug toggle. **IF YOU NEED DEBUG OFF
  REGARDLESS OF ENV** then call pl.DisableDebugLogs() **AFTER**
  InitPrettyLogger().

  About golang itself: 

  Args to `LogDebug("msg %v", f())` **ARE STILL EVALUATED** when the call runs.
  For heavy f(), use `if pl.DebugEnabled() {...}` or move formatting inside that
  block.


  ### All available log types/methods/funcs

  ```go
  pl.Log("info message")
  pl.LogDebug("debug message")
  pl.LogWarn("warning message")
  pl.LogError("error message")
  pl.LogFatal("some fatal msg")
  ```

</details>

<details>
  <summary>Older documentation (copy pasted, not updated)</summary>
  ## Installation

  Use `go get` to install with the latest tag `v0.1.2` (recommended)

  ```bash
  go get -u github.com/fluffysnowman/prettylogger@v0.1.2
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


  ## Usage (old)

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
</details>



