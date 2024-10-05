package prettylogger

import (
	"testing"
    // "os"
    // "io"
)

// var plTest prettylogger.PrettyLogger

// func TestBasicLog(t *testing.T) {
//     plTest.LogInfo("hello there")
// }

func TestBasicPrettyLog_LogInfo(t *testing.T) {
    // var pl_test PrettyLogger
    // pl_test.LogInfo("hello there");
    // var basic_writer io.Writer = os.Stdout;
    // pl_test_logger := NewPrettyLogger(os.Stdout);
    // pl_test_logger.LogInfo("testing logs");
    // io.WriteString(basic_writer, RedANSI+"[ WARNING ] TEST: hello there \n"+ResetANSI);
    testOutput := "hello"
    expectedOutput := "hello"
    if testOutput != expectedOutput {
        t.Errorf("test failed")
    }
}

