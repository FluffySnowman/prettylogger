/*
# PL: A simple Pretty Logger made in go

All the code for the pretty logger is contained within this one file for
simplicity.

# Usage

Initialise the pretty logger with a config (can be SIMPLE, SIMPLE2 [newly added
log format] or TIMEBASED) which is automatically used everywhere once set.

Example:

	pl.InitPrettyLogger("SIMPLE")       // basic
	pl.InitPrettyLogger("SIMPLE2")      // new concise format
	pl.InitPrettyLogger("TIMEBASED")    // shows timestamps
	pl.LogInfo("Hello World")

Multiple arguments:

	pl.LogDebug("this is a debug log %v", "foo bar")

# All available log types

	pl.Log("log")
	pl.LogDebug("debug")
	pl.LogError("error")
	pl.LogInfo("info")
	pl.LogFatal("fatal")
	pl.LogSuccess("success")
	pl.LogFailure("failure")
	pl.LogOK("ok")
	pl.LogErrorBG("errorbg")
	pl.LogFailureBG("failerbg")

Author: @FluffySnowman (GitHub)

Source: https://github.com/FluffySnowman/prettylogger
*/
package prettylogger

// package main // main here for testing

import (
	"fmt"
	"io"
	"os"
	"strings"
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

	// Bright foreground colours
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
const (
	ResetANSI = "\033[0m"
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

// Format for normal SIMPLE logs (original)
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

// Format for timestamp TIMEBASED logs (original)
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

// format for SIMPLE2 logs
var Simple2Log = LogTypes{
	LogLog:     "INFO [%s|%s] %v\n", // log is info in simple2
	DebugLog:   "DEBUG[%s|%s] %v\n",
	ErrorLog:   "ERROR[%s|%s] %v\n",
	FatalLog:   "FATAL[%s|%s] %v\n",
	InfoLog:    "INFO [%s|%s] %v\n",
	SuccessLog: "INFO [%s|%s] %v\n", // success is info
	FailedLog:  "ERROR[%s|%s] %v\n", // failure is error
	OkayLog:    "INFO [%s|%s] %v\n", // ok is info
}

// Configuration for the logger
type PrettyLogger struct {
	writer   io.Writer
	logLevel string
	color    string
	logType  string
}

func getCurrentTimestamp() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func getSimple2Timestamp() string {
	return time.Now().Format("01-02|15:04:05.000")
	// return time.Now().Format("01-02-2006|15:04:05.000")
}

// Global pretty logger instance (used to r/w config from)
var prettyLoggerConfig *PrettyLogger

// sets the log format to SIMPLE, SIMPLE2 (newly added format) or TIMEBASED.
//
// Default is set to SIMPLE.
//
// Do not set the log format multiple times as it will cause problems
func InitPrettyLogger(prettyLogType string) {
	if len(prettyLogType) < 1 {
		prettyLogType = "SIMPLE"
	}
	prettyLoggerConfig = &PrettyLogger{
		writer:  os.Stdout,
		color:   WhiteFgANSI,
		logType: prettyLogType,
	}
}

func getLogType() LogTypes {
	switch prettyLoggerConfig.logType {
	case "SIMPLE":
		return SimpleLog
	case "TIMEBASED":
		return TimestampLog
	case "SIMPLE2":
		return Simple2Log
	default:
		return SimpleLog
	}
}

func printLog(logFormat string, logColor string, message string, timestamp bool) {
	if prettyLoggerConfig == nil {
		return
	}

	if prettyLoggerConfig.logType == "SIMPLE2" {
		timestampStr := getSimple2Timestamp()
		parts := strings.SplitN(timestampStr, "|", 2)
		datePart := parts[0]
		timePart := parts[1]

		levelWord := ""
		restFormat := ""
		{
			idxBracket := strings.Index(logFormat, "[")
			idxSpace := strings.Index(logFormat, " ")
			if idxBracket == -1 {
				idxBracket = len(logFormat)
			}
			if idxSpace == -1 {
				idxSpace = len(logFormat)
			}

			endIdx := idxBracket
			if idxSpace != -1 && idxSpace < idxBracket {
				endIdx = idxSpace
			}
			levelWord = logFormat[:endIdx]
			restFormat = logFormat[len(levelWord):]
		}

		coloredDate := YellowFgANSI + datePart + ResetANSI
		coloredTime := CyanFgANSI + timePart + ResetANSI
		coloredMessage := WhiteFgANSI + message + ResetANSI

		fmt.Fprintf(prettyLoggerConfig.writer,
			"%s%s%s"+restFormat,
			logColor, levelWord, ResetANSI,
			coloredDate, coloredTime, coloredMessage,
		)
		return
	}

	timestampStr := ""
	if timestamp || prettyLoggerConfig.logType == "TIMEBASED" {
		timestampStr = " " + getCurrentTimestamp() + " "
	}

	fmt.Fprintf(
		prettyLoggerConfig.writer,
		logFormat,
		logColor,
		timestampStr,
		ResetANSI,
		message,
	)
}

func Log(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.LogLog, GreenFgANSI, formattedMessage, false)
}

func LogDebug(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.DebugLog, CyanFgANSI, formattedMessage, false)
}

func LogError(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.ErrorLog, RedFgANSI, formattedMessage, false)
}

func LogInfo(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.InfoLog, GreenFgANSI, formattedMessage, false)
}

func LogFatal(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.FatalLog, BrightRedFgANSI, formattedMessage, false)
}

func LogSuccess(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.SuccessLog, GreenFgANSI, formattedMessage, false)
}

func LogFailure(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.FailedLog, RedFgANSI, formattedMessage, false)
}

func LogOK(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.OkayLog, GreenFgANSI, formattedMessage, false)
}

func LogErrorBG(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.ErrorLog, RedBgANSI, formattedMessage, false)
}

func LogFailureBG(format string, a ...interface{}) {
	logFormats := getLogType()
	formattedMessage := fmt.Sprintf(format, a...)
	printLog(logFormats.FailedLog, YellowBgANSI, formattedMessage, false)
}

// // Using main here for testing
// func main() {

// 	// Init the logger with simple/complex config
// 	// InitPrettyLogger("SIMPLE")
// 	// InitPrettyLogger("TIMEBASED")
// 	InitPrettyLogger("SIMPLE2")
// 	// LogDebug("this is a DEBUG log").Print()
// 	// LogError("this is a error log").Print()
// 	// LogSuccess("this is a success log").Print()
// 	// LogInfo("this is a INFO log").Print()
// 	// LogOK("this is ok").Print();
// 	// LogOK("this is ok").Print();
// 	// LogOK("this is ok").Print();
// 	// LogOK("this is ok").Print();
// 	// LogOK("this is ok").Print();
// 	// LogDebug("this i sa debeg log over here ").Print()
// 	// println()
// 	Log("log")
// 	LogDebug("debug")
// 	LogError("error")
// 	LogInfo("info")
// 	LogFatal("fatal")
// 	LogSuccess("success")
// 	LogFailure("failure")
// 	LogOK("ok")
// 	LogErrorBG("errorbg")
// 	LogFailureBG("failerbg")
// 	// Log("connecting to database...")
// 	// LogOK("database connected")
// 	// LogSuccess("database connected")
// 	// LogInfo("this should be an info log")
// 	// LogInfo("this should be an info log")
// 	// LogDebug("query: SELECT * FROM users WHERE username = $1")
// 	// LogFailure("failed to execute query")
// 	// LogFailureBG("failed to execute query")
// 	// LogFatal("segmentation fault, core dumped")
// 	// LogErrorBG("DUMPING CORE..")
// 	//     LogInfo("preparing to execute query...").Print()
// 	//     // LogError("DUMPING CORE..").Print()

// 	// // println("testing all log types below to see how they look\n")
// 	// // Log("hello there").Print();
// 	// // LogDebug("hello there").Print();
// 	// // LogError("hello there").Print();
// 	// // LogInfo("hello there").Print();
// 	// // LogFatal("hello there").Print();
// 	// // LogSuccess("hello there").Print();
// 	// // LogFailure("hello there").Print();
// 	// // LogOK("hello there").Print();
// }

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
