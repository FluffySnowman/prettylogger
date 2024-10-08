/*
# PL: A simple Pretty Logger made in go

All the code for the pretty logger is contained within this one file for
simplicity.

# Usage

Initialise the pretty logger with a config (can be SIMPLE, TIMEBASED) which
is automatically used everywhere once set.

Example:

	pl.InitPrettyLogger("SIMPLE")       // basic
	pl.InitPrettyLogger("TIMEBASED")    // shows timestamps
	pl.LogInfo("Hello World")

Multiple arguments:

	pl.LogDebug("this is a debug log %v", "foo bar")

Force a timestamp log:

	pl.LogDebug("this is a debug log %v", "foo bar").Timestamp().Print()

Changing SIMPLE to TIMEBASED & vice versa will cause problems hence use the
.Timestamp() and .Print() when timestamps are needed.

Author: @FluffySnowman (GitHub)

Source: https://github.com/FluffySnowman/prettylogger
*/
package main

// package main // main here for testing

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Foreground colours (only text)
const (
	// Normal forground colours
	RedFgANSI     = "\033[31m"
	GreenFgANSI   = "\033[32m"
	YellowFgANSI  = "\033[33m"
	BlueFgANSI    = "\033[34m"
	MagentaFgANSI = "\033[35m"
	CyanFgANSI    = "\033[36m"
	WhiteFgANSI   = "\033[37m"
	BlackFgANSI   = "\033[30m"

	// Bright foreground colours (might not be used anywhere but here cos yes)
	BrightBlackFgANSI   = "\033[90m"
	BrightRedFgANSI     = "\033[91m"
	BrightGreenFgANSI   = "\033[92m"
	BrightYellowFgANSI  = "\033[93m"
	BrightBlueFgANSI    = "\033[94m"
	BrightMagentaFgANSI = "\033[95m"
	BrightCyanFgANSI    = "\033[96m"
	BrightWhiteFgANSI   = "\033[97m"
)

// Background Colours (background of the text & won't affect the text content)
const (
	// Normal background colours
	RedBgANSI     = "\033[41m"
	GreenBgANSI   = "\033[42m"
	YellowBgANSI  = "\033[43m"
	BlueBgANSI    = "\033[44m"
	MagentaBgANSI = "\033[45m"
	CyanBgANSI    = "\033[46m"
	WhiteBgANSI   = "\033[47m"
	BlackBgANSI   = "\033[40m"

	// Bright background colours (might not be used since the bright ones are
	// harder to read)
	BrightBlackBgANSI   = "\033[100m"
	BrightRedBgANSI     = "\033[101m"
	BrightGreenBgANSI   = "\033[102m"
	BrightYellowBgANSI  = "\033[103m"
	BrightBlueBgANSI    = "\033[104m"
	BrightMagentaBgANSI = "\033[105m"
	BrightCyanBgANSI    = "\033[106m"
	BrightWhiteBgANSI   = "\033[107m"
)

// All text formatting ansi codes
// These are for the way that text should visually look such as italic, dim,
// bold, blink (may not be supported in many terminals) and do not affect the
// colours of the text (maybe except the Dim one)
const (
	ResetANSI               = "\033[0m" // Reset all (most used here)
	BoldFormatANSI          = "\033[1m" // Bold
	DimFormatANSI           = "\033[2m" // Dim
	ItalicFormatANSI        = "\033[3m" // Italic
	UnderlineFormatANSI     = "\033[4m" // Underline
	BlinkFormatANSI         = "\033[5m" // Blink
	ReverseFormatANSI       = "\033[7m" // Reverse (swaps bg and fg )
	HiddenFormatANSI        = "\033[8m" // John Cena
	StrikethroughFormatANSI = "\033[9m" // Strikethrough

	// Reset all formats (these are only here for specific use cases. The
	// general ResetANSI will be used everywhere)
	ResetBoldFormatANSI          = "\033[21m"
	ResetDimFormatANSI           = "\033[22m"
	ResetItalicFormatANSI        = "\033[23m"
	ResetUnderlineFormatANSI     = "\033[24m"
	ResetBlinkFormatANSI         = "\033[25m"
	ResetReverseFormatANSI       = "\033[27m"
	ResetHiddenFormatANSI        = "\033[28m"
	ResetStrikethroughFormatANSI = "\033[29m"
)

// All variants of log side/type (the text put before the thing the user will
// log).
// These will be printed before the respective logs and some can change
const (
	LogLogBasic   = "[   LOG   ] "
	DebugLogBasic = "[  DEBUG  ] "
	ErrorLogBasic = "[  ERROR  ] "
	FatalLogBasic = "[  FATAL  ] "
	InfoLogBasic  = "[  INFO   ] "
)

// Struct for all the log types so that diff log formats can be defined without
// repeating shit
type LogTypes struct {
	LogLog     string
	DebugLog   string
	ErrorLog   string
	FatalLog   string
	InfoLog    string
	SuccessLog string
	FailedLog  string
	OkayLog    string
}

// // Original simple log format where the whole thing except the log message is
// // coloured. Now changed to the struct below
// var SimpleLog = LogTypes{
// 	LogLog:     "%s[   LOG   %s]%s %v\n",
// 	DebugLog:   "%s[  DEBUG  %s]%s %v\n",
// 	ErrorLog:   "%s[  ERROR  %s]%s %v\n",
// 	FatalLog:   "%s[  FATAL  %s]%s %v\n",
// 	InfoLog:    "%s[  INFO   %s]%s %v\n",
// 	SuccessLog: "%s[ SUCCESS %s]%s %v\n",
// 	FailedLog:  "%s[ FAILURE %s]%s %v\n",
// 	OkayLog:    "%s[   OK    %s]%s %v\n",
// }

// Format for all the logs with (mostly) string format specifiers for ansi
// colours, static text indicating what log it is, the log message and then
// finally the ansi colour reset .
var SimpleLog = LogTypes{
	LogLog:     "[%s   LOG   %s%s] %v\n",
	DebugLog:   "[%s  DEBUG  %s%s] %v\n",
	ErrorLog:   "[%s  ERROR  %s%s] %v\n",
	FatalLog:   "[%s  FATAL  %s%s] %v\n",
	InfoLog:    "[%s  INFO   %s%s] %v\n",
	SuccessLog: "[%s SUCCESS %s%s] %v\n",
	FailedLog:  "[%s FAILURE %s%s] %v\n",
	OkayLog:    "[%s   OK    %s%s] %v\n",
}

// // Original timestamp log (not in use anymore)
// // Same as SimpleLog but with timestamps
// var TimestampLog = LogTypes{
// 	LogLog:     "%s[   LOG   %s]%s %v\n",
// 	DebugLog:   "%s[  DEBUG  %s]%s %v\n",
// 	ErrorLog:   "%s[  ERROR  %s]%s %v\n",
// 	FatalLog:   "%s[  FATAL  %s]%s %v\n",
// 	InfoLog:    "%s[  INFO   %s]%s %v\n",
// 	SuccessLog: "%s[ SUCCESS %s]%s %v\n",
// 	FailedLog:  "%s[ FAILURE %s]%s %v\n",
// 	OkayLog:    "%s[   OK    %s]%s %v\n",
// }

// Same as SimpleLog but with timestamps inside the []'s
var TimestampLog = LogTypes{
	LogLog:     "[%s   LOG   %s%s] %v\n",
	DebugLog:   "[%s  DEBUG  %s%s] %v\n",
	ErrorLog:   "[%s  ERROR  %s%s] %v\n",
	FatalLog:   "[%s  FATAL  %s%s] %v\n",
	InfoLog:    "[%s  INFO   %s%s] %v\n",
	SuccessLog: "[%s SUCCESS %s%s] %v\n",
	FailedLog:  "[%s FAILURE %s%s] %v\n",
	OkayLog:    "[%s   OK    %s%s] %v\n",
}

// Configuration for the logger
type PrettyLogger struct {
	writer   io.Writer
	logLevel string
	color    string
	logType  string
}

// Gets current time formatted to rfc3339 to milliseconds with Z
func getCurrentTimestamp() string {
	// return time.Now().Format(time.RFC3339Nano)[:23] + "Z"
	// return time.Now().Format(time.RFC3339);
	return time.Now().Format("2006/01/02 15:04:05")
}

// Global pretty logger instance (used to r/w config from)
var prettyLoggerConfig *PrettyLogger

// Ititializez the global pretty logger config with optional arguments
func InitPrettyLogger(prettyLogType string) {
	// If nothing is passed then its set to SIMPLE
	if len(prettyLogType) < 1 {
		prettyLogType = "SIMPLE"
	}
	prettyLoggerConfig = &PrettyLogger{
		writer:  os.Stdout,   // default to stdout
		color:   WhiteFgANSI, // Default color
		logType: prettyLogType,
	}
}

// Gets the type of the log (simple, timestamp etc) and returns it which is then
// used in the actual logging
func getLogType() LogTypes {
	switch prettyLoggerConfig.logType {
	case "SIMPLE":
		return SimpleLog
	case "TIMEBASED":
		return TimestampLog
	// case "TIMECOMPLEX":
	// 	return ComplexTimestampLog
	default:
		return SimpleLog
	}
}

// Main LogEntry struct used throughout the project and for method chaining
type LogEntry struct {
	logFormat string
	logColor  string
	message   string
	timestamp bool
}

// Includes the timestamp in the log.
//
// This should be chained with LogInfo, LogDebug, ... etc. and should have a
// .Print() after it.
//
// Example:
//
// LogDebug("this is a debug log %v", "with a timestamp").Timestamp().Print()
//
// Which would output:
//
// [  DEBUG  2024/10/06 17:08:01 ] this is a debug log with a timestamp
//
// The timestamp will not be logged unless a .Print() is chained after it.
func (le *LogEntry) Timestamp() *LogEntry {
	le.timestamp = true
	return le
}

// Print method to output the log (should be chained after .Timestamp())
func (le *LogEntry) Print() {
	timestamp := ""
	if le.timestamp || prettyLoggerConfig.logType == "TIMEBASED" {
		timestamp = " " + getCurrentTimestamp() + " "
	}
	if prettyLoggerConfig != nil {
		fmt.Fprintf(
			prettyLoggerConfig.writer,
			le.logFormat,
			le.logColor,
			timestamp,
			ResetANSI,
			le.message,
		)
	}
}

// Log general data (green)
func Log(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.LogLog,
		logColor:  GreenFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Debug logs (cyan)
func LogDebug(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.DebugLog,
		logColor:  CyanFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Error logs (red) colour
func LogError(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.ErrorLog,
		logColor:  RedFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Info logs (cyan)
func LogInfo(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.InfoLog,
		logColor:  CyanFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// fatal logs (bright red)
func LogFatal(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.FatalLog,
		logColor:  BrightRedFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Success logs (green)
func LogSuccess(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.SuccessLog,
		logColor:  GreenFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Failed logs (yellow)
func LogFailure(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.FailedLog,
		logColor:  YellowFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}

// Okay logs (green)
func LogOK(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.OkayLog,
		logColor:  GreenFgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}


// Log with filled background
func LogErrorBG(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.ErrorLog,
		logColor:  RedBgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}


// Log with filled background
func LogFailureBG(format string, a ...interface{}) *LogEntry {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	entry := &LogEntry{
		logFormat: logFormats.FailedLog,
		logColor:  YellowBgANSI,
		message:   formattedMessage,
		timestamp: false,
	}
	return entry
}


// Using main here for testing
func main() {

	// Init the logger with simple/complex config
	// InitPrettyLogger("SIMPLE")
	InitPrettyLogger("TIMEBASED")
	// LogDebug("this is a DEBUG log").Print()
	// LogError("this is a error log").Print()
	// LogSuccess("this is a success log").Print()
	// LogInfo("this is a INFO log").Print()
	// LogOK("this is ok").Print();
	// LogOK("this is ok").Print();
	// LogOK("this is ok").Print();
	// LogOK("this is ok").Print();
	// LogOK("this is ok").Print();
	// LogDebug("this i sa debeg log over here ").Print()
	println()
	Log("connecting to database...").Print()
	LogOK("database connected").Print()
    LogInfo("preparing to execute query...").Print()
    LogDebug("query: SELECT * FROM users WHERE username = $1").Print()
	LogFailure("failed to execute query").Print()
	LogFailureBG("failed to execute query").Print()
	LogFatal("segmentation fault, core dumped\n\n").Print()
    LogErrorBG("DUMPING CORE..").Print()
    // LogError("DUMPING CORE..").Print()

    // println("testing all log types below to see how they look\n")
    // Log("hello there").Print();
    // LogDebug("hello there").Print();
    // LogError("hello there").Print();
    // LogInfo("hello there").Print();
    // LogFatal("hello there").Print();
    // LogSuccess("hello there").Print();
    // LogFailure("hello there").Print();
    // LogOK("hello there").Print();
}

// 	LogDebug("this is a debug log %v", "which should print something").Timestamp().Print()
// 	// LogDebug("this is a debug log %v", "which should print something").Timestamp().Print()
// 	LogInfo("job info: %v ", "running job ...")
// 	LogInfo("job info: %v ", "job SUCCESS").Timestamp().Print()

// // Testing different data types to see if it works
// var someString string
// var someInt int
// var someFloat float64
// someString = "this is some string"
// someInt = 42069
// someFloat = 2981389.829810

// LogError("some ERROR shit here")
// LogDebug("this is a debug log")
// LogInfo("some info here")
// LogSuccess("success with some int -> %v", someInt)
// LogFailure("this is a failure message")
// LogFatal("failed to 420: %v %v", someString, someFloat)

// Example of logging with timestamp
// LogDebug("this is a debug log with timestamp").Timestamp().Print()
// LogInfo("some info here with timestamp").Timestamp().Print()
// }

// Test with basic raw ascii for comparison
// var errorTest error
// errorTest = errors.New("some error happened")
// errorTest := nil
// if errorTest == nil {
// 	fmt.Printf("\033[31m[ ERROR ] \033[97;40m%v\033[0m\n", errorTest)
// }
