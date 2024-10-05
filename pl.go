// PL: A simple Pretty Logger made in go
//
// All the code for the pretty logger is contained within this one file for
// simplicity.
//
// Most of everything here uses receivers.
// 
package prettylogger
// package main

import (
	"io"
	"log"
    "fmt"
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
	ResetANSI             = "\033[0m"       // Reset all (most used here)
	BoldFormatANSI        = "\033[1m"       // Bold
	DimFormatANSI         = "\033[2m"       // Dim
	ItalicFormatANSI      = "\033[3m"       // Italic
	UnderlineFormatANSI   = "\033[4m"       // Underline
	BlinkFormatANSI       = "\033[5m"       // Blink
	ReverseFormatANSI     = "\033[7m"       // Reverse (swaps bg and fg )
	HiddenFormatANSI      = "\033[8m"       // John Cena
	StrikethroughFormatANSI = "\033[9m"     // Strikethrough

	// Reset all formats (these are only here for specific use cases. The
    // general ResetANSI will be used everywherec)
	ResetBoldFormatANSI         = "\033[21m"
	ResetDimFormatANSI          = "\033[22m"
	ResetItalicFormatANSI       = "\033[23m"
	ResetUnderlineFormatANSI    = "\033[24m"
	ResetBlinkFormatANSI        = "\033[25m"
	ResetReverseFormatANSI      = "\033[27m"
	ResetHiddenFormatANSI       = "\033[28m"
	ResetStrikethroughFormatANSI = "\033[29m"
)

// All variants of log side (the text put before the thing the user will log).
// These will be printed befor the respective logs.
const (
    LOGLogBasic     =       "[    LOG    ]"
    DEBUGLogBasic   =       "[   DEBUG   ]"
    ERRORLogBasic   =       "[   ERROR   ]"
    FATALLogBasic   =       "[   FATAL   ]"
    INFOLogBasic    =       "[    INF    ]"
)

type PrettyLogger struct {
	LogLogger *log.Logger
	LogWriter io.Writer
}

func NewPrettyLogger(writer io.Writer) *PrettyLogger {
    return &PrettyLogger{
        LogLogger: log.New(writer, "", log.LstdFlags),
        LogWriter: writer,
    }
}

func (pl *PrettyLogger) LogInfo(message string) {
    pl.LogLogger.Printf("[ INFO ] %s", message);
}

// func main() {
//     var basic_writer io.Writer = os.Stdout;
//     pl_test_logger := NewPrettyLogger(os.Stdout);
//     pl_test_logger.LogInfo("hello there");
//     io.WriteString(basic_writer, RedANSI+"[ WARNING ] Something went wrong\n"+ResetANSI);
// }


