// PL: A simple Pretty Logger made in go
//
// All the code for the pretty logger is contained within this one file for
// simplicity.
//
// package prettylogger
package main // main here for testing

import (
	"fmt"
	"io"
	"time"

	// "log"
	"os"

	// "fmt"
	// "os"
	// "testing"
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
	// general ResetANSI will be used everywherec)
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
// These will be printed befor the respective logs and some can change
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
	LogLog   string
	DebugLog string
	ErrorLog string
	FatalLog string
	InfoLog  string
}

// Simple log format with static text and string format specifiers for the ansi,
// message and reset.
var SimpleLog = LogTypes{
	LogLog:   "%s[   LOG   ] %s %s \n",
	DebugLog: "%s[  DEBUG  ] %s %s \n",
	ErrorLog: "%s[  ERROR  ] %s %s \n",
	FatalLog: "%s[  FATAL  ] %s %s \n",
	InfoLog:  "%s[  INFO   ] %s %s \n",
}

var TimestampLog = LogTypes{
	LogLog:   "[  LOG  %v ] ",
	DebugLog: "[ DEBUG %v ] ",
	ErrorLog: "[ ERROR %v ] ",
	FatalLog: "[ FATAL %v ] ",
	InfoLog:  "[ INFO  %v ] ",
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
	return time.Now().Format(time.RFC3339Nano)[:23] + "Z"
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
        case "SIMPLETIME": 
            return TimestampLog
        default:
            return SimpleLog
    }
}

// Info logs with white foreground
// func LogInfo(message string) {
// 	if prettyLoggerConfig != nil {
// 		writeLog(InfoLogBasic, message)
// 	}
// }

// Debug logs with cyan background
func LogDebug(message string) {
    logFormats := getLogType()
	if prettyLoggerConfig != nil {
        writeLog(logFormats.DebugLog, CyanFgANSI, message)
	}
}

// func LogError(message string) {
// 	if prettyLoggerConfig != nil {
// 		writeLog(ErrorLogBasic, message)
// 	}
// }

// function that actually writes the logs
func writeLog(logFormat string, logColor string, message string) {
	if prettyLoggerConfig != nil {

        fmt.Fprintf(
            prettyLoggerConfig.writer,
            logFormat, 
            logColor,
            message,
            ResetANSI,
        )

        // fmt.Fprintf(prettyLoggerConfig.writer, "%s%s %s%s\n",
        // prettyLoggerConfig.color, prefix, message, ResetANSI)
	}
}

// Using main here for testing
func main() {
	// Init the logger e 
	InitPrettyLogger("SIMPLE")

	LogDebug("this should be printed with the debug log ")
	// LogError("some ERROR shit here")
    // LogInfo("some info here ")
}

// SetColor(WhiteFgANSI)
// // Allows setting the colour for the logger (WhiteFgANSI, GreenFgANSI, etc.)
// func SetColor(color string) {
// 	if prettyLoggerConfig != nil {
// 		prettyLoggerConfig.color = color
// 	}
// }

