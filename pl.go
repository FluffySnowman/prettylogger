/*
# PL: A simple Pretty Logger made in go

All the code for the pretty logger is contained within this one file for
simplicity.

# Usage

Initialise the pretty logger with a config/string specifying options seperated
by commas (,)

Import alias for simplicity must be:

	import (
	    pl "github.com/fluffysnowman/prettylogger"
	)

Example:

	pl.InitPrettyLogger("V3")                                       // original (just times, no dates) time format
	pl.InitPrettyLogger("V3,TEA")                                   // TEA (bri'ish) time format
	pl.InitPrettyLogger("V3,BURGER,FILECOLOR")                      // BURGER (american) time + colored file
	pl.InitPrettyLogger("V3,READABLE,FILEONLY,FILE=logfile.log")    // readable time w/ month names etc, file-only written to logfile.log

# All available config/string options

  - TEA           : british time format (dd-mm-yy hh:mm:ss.mmm)
  - BURGER        : american time format (mm-dd-yy hh:mm:ss.mmm)
  - READABLE      : readable time format (ddmonyyyy hh:mm ss.mmm s)
  - FILECOLOR     : write colour logs to file
  - FILEONLY      : write logs only to file, no stdout output etc.  ..
  - FILE=<path>   : write logs to specified file path, if exists, if not- created

Multiple arguments:

	pl.LogDebug("this is a debug log %v", "meow meow mf")

# All available log types/methods/funcs

	pl.Log("info message")
	pl.LogDebug("debug message")
	pl.LogWarn("warning message")
	pl.LogError("error message")

Author: @FluffySnowman (GitHub)

Source: https://github.com/FluffySnowman/prettylogger

License: MIT License
*/
package prettylogger

// package main // main here for testing

import (
	"fmt"
	// "io"
	"os"
	"strings"
	"time"
)

const (
	COLOR_DEBUG = "\x1b[36m"
	COLOR_INFO  = "\x1b[32m"
	COLOR_WARN  = "\x1b[33m"
	COLOR_ERROR = "\x1b[31m"
	COLOR_RESET = "\x1b[0m"
)

var cfg struct {
	timeMode  int
	fileColor bool
	fileOnly  bool
	logFile   *os.File
}

/*
OPTS: TEA, BURGER, READABLE, FILECOLOR, FILEONLY, FILE=<path>

  - TEA           : british time format (dd-mm-yy hh:mm:ss.mmm)
  - BURGER        : american time format (mm-dd-yy hh:mm:ss.mmm)
  - READABLE      : readable time format (ddmonyyyy hh:mm ss.mmm s)
  - FILECOLOR     : write colour logs to file
  - FILEONLY      : write logs only to file, no stdout output etc.  ..
  - FILE=<path>   : write logs to specified file path, if exists, if not- created
*/
func InitPrettyLogger(opts string) {
	cfg = struct {
		timeMode  int
		fileColor bool
		fileOnly  bool
		logFile   *os.File
	}{}
	for _, o := range strings.Fields(strings.ReplaceAll(opts, ",", " ")) {
		u := strings.ToUpper(o)
		switch {
		case u == "TEA":
			cfg.timeMode = 1
		case u == "BURGER":
			cfg.timeMode = 2
		case u == "READABLE":
			cfg.timeMode = 3
		case u == "FILECOLOR":
			cfg.fileColor = true
		case u == "FILEONLY":
			cfg.fileOnly = true
		case strings.HasPrefix(u, "FILE="):
			path := o[len("FILE="):]
			if f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644); err == nil {
				cfg.logFile = f
			}
		}
	}
}

func formattedTime() string {
	now := time.Now()
	ms := now.Nanosecond() / 1e6
	switch cfg.timeMode {
	// case 1, 2:
    // british format
    case 1:
		// dd-mm-yy HH:MM:SS.mmm
		return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d.%03d",
			now.Day(), now.Month(), now.Year()%100,
			now.Hour(), now.Minute(), now.Second(), ms)
    // american format for 2 
    case 2: 
		// mm-dd-yy HH:MM:SS.mmm
		return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d.%03d",
			now.Month(), now.Day(), now.Year()%100,
			now.Hour(), now.Minute(), now.Second(), ms)
	case 3:
		// DDMonYYYY HH:MM SS.mmm s
		mon := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
			"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
		p1 := fmt.Sprintf("%02d%s%04d", now.Day(), mon[now.Month()-1], now.Year())
		p2 := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
		p3 := fmt.Sprintf("%02d.%03ds", now.Second(), ms)
		return fmt.Sprintf("%s %s %s", p1, p2, p3)
	default:
		// HH:MM:SS.mmm
		return fmt.Sprintf("%02d:%02d:%02d.%03d",
			now.Hour(), now.Minute(), now.Second(), ms)
	}
}

func logV3(levelTag, color, msg string) {
	ts := formattedTime()
	lineOut := fmt.Sprintf("%s %s%s%s %s\n", ts, color, levelTag, COLOR_RESET, msg)
	if !cfg.fileOnly {
		fmt.Fprint(os.Stdout, lineOut)
	}
	if f := cfg.logFile; f != nil {
		if cfg.fileColor {
			fmt.Fprint(f, lineOut)
		} else {
			fmt.Fprintf(f, "%s %s %s\n", ts, levelTag, msg)
		}
		_ = f.Sync()
	}
}

func LogDebug(format string, a ...interface{}) {
	logV3("DBG", COLOR_DEBUG, fmt.Sprintf(format, a...))
}

func Log(format string, a ...interface{}) {
	logV3("INF", COLOR_INFO, fmt.Sprintf(format, a...))
}

func LogWarn(format string, a ...interface{}) {
	logV3("WRN", COLOR_WARN, fmt.Sprintf(format, a...))
}

func LogError(format string, a ...interface{}) {
	logV3("ERR", COLOR_ERROR, fmt.Sprintf(format, a...))
}

func main() {

	// // default shit test
	// InitPrettyLogger("V3")
	// Log("hello there with default shit")
	// LogDebug("eeeeee")
	// LogWarn("warning here")
	// LogError("error shit: %v", "lol")

	// diff modes test
	InitPrettyLogger("V3,BURGER,FILECOLOR,FILE=app.log")
	Log("hello there")
	LogDebug("eeeeee")
	LogWarn("warning here")
	LogError("error shit: %v", "lol")

	// InitPrettyLogger("V3,READABLE,FILEONLY,FILE=readable.log")
	// Log("hello there")
	// LogDebug("eeeeee")
	// LogWarn("warning here")
	// LogError("error shit: %v", "lol")

}
