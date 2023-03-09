package hlogger

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDebugEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = true
	PrintColors = false

	Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | debug\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugfEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = true
	PrintColors = false

	Debugf("hello %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | hello 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = false

	Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSeparatorDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = false

	DebugSeparator("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSeparatorEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = true
	PrintColors = false

	DebugSeparator("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | ====[ debug ]===================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = true
	PrintColors = false

	data := map[string]string{"hello": "world"}

	DebugDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	DebugMode = true
	PrintColors = false

	data := map[string]string{"hello": "world"}

	DebugDump(data, "dprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | dprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfo(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Info("info 100%")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 100%\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfof(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Infof("info %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	InfoSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | ====[ info ]====================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	InfoDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	InfoDump(data, "iprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | iprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarn(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Warn("warn")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Warnf("warn %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	WarnSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | ====[ info ]====================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	WarnDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	WarnDump(data, "wprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | wprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestError(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Error("error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | error\n", actualStdErr)

}

func TestErrorf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	Errorf("error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | error 2\n", actualStdErr)

}

func TestErrorSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	ErrorSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | ====[ info ]====================================================================\n", actualStdErr)

}

func TestErrorDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	ErrorDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr)

}

func TestErrorDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	data := map[string]string{"hello": "world"}

	ErrorDump(data, "eprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | eprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr)

}

func TestStackTrace(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	StackTrace(errors.New("my error"))

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.True(t, strings.HasPrefix(actualStdErr, "test | ERROR | my error\n"))

}

type CustomError struct{}

func (m *CustomError) Error() string {
	return "boom"
}

func Test_StackTraceCustom(t *testing.T) {

	resetLogConfig()
	stdout, _ := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	StackTrace(&CustomError{})

	actualStdOut := stdout.String()

	assert.Equal(t, "", actualStdOut, "stdout")
}

func TestFormattedStackTrace(t *testing.T) {

	actual := FormattedStackTrace(errors.New("my error"))
	assert.NotEmpty(t, actual)

}

func TestFatal(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	oldOsExit := OsExit
	defer func() {
		OsExit = oldOsExit
	}()

	var got int
	OsExit = func(code int) {
		got = code
	}

	Fatal("fatal error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | FATAL | fatal error\n", actualStdErr)
	assert.Equal(t, 1, got)

}

func TestFatalf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	PrintColors = false

	oldOsExit := OsExit
	defer func() {
		OsExit = oldOsExit
	}()

	var got int
	OsExit = func(code int) {
		got = code
	}

	Fatalf("fatal error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | FATAL | fatal error 2\n", actualStdErr)
	assert.Equal(t, 1, got)

}

func TestCheckError(t *testing.T) {

	type test struct {
		name             string
		err              error
		debug            bool
		expectedStdout   string
		expectedStderr   string
		expectedExitCode int
	}

	var tests = []test{
		{"nil-debug-nocolor", nil, true, "", "", -1},
		{"nil-debug-color", nil, true, "", "", -1},

		{"err-release-nocolor", errors.New("test"), false, "", "test | FATAL | test\n", 1},
		{"err-debug-nocolor", errors.New("test"), true, "", "test | FATAL | test\n", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resetLogConfig()
			stdout, stderr := redirectOutput()
			defer resetLogOutput()

			oldOsExit := OsExit
			defer func() {
				OsExit = oldOsExit
			}()

			var got int
			OsExit = func(code int) {
				got = code
			}

			DebugMode = tc.debug
			PrintColors = false

			CheckError(tc.err)

			actualStdOut := stdout.String()
			actualStdErr := stderr.String()

			assert.Equal(t, tc.expectedStdout, actualStdOut)
			if tc.debug {
				assert.True(t, strings.HasPrefix(actualStdErr, tc.expectedStderr), actualStdErr)
			} else {
				assert.Equal(t, tc.expectedStderr, actualStdErr)
			}

			if tc.expectedExitCode > 0 {
				assert.Equal(t, 1, got)
			}

		})
	}
}

func TestLog(t *testing.T) {
	DebugMode = true
	PrintTimestamp = true
	PrintColors = true

	myVar := map[string]string{"hello": "world"}

	Debug("arg1", "arg2")
	Debugf("arg1 %d", 1)
	DebugDump(myVar, "prefix")
	DebugSeparator("title")

	Info("arg1", "arg2")
	Infof("arg1 %d", 1)
	InfoDump(myVar, "prefix")
	InfoSeparator("title")

	Warn("arg1", "arg2")
	Warnf("arg1 %d", 1)
	WarnDump(myVar, "prefix")
	WarnSeparator("title")

	Error("arg1", "arg2")
	Errorf("arg1 %d", 1)
	ErrorDump(myVar, "prefix")
	ErrorSeparator("title")

	//Fatal("arg1", "arg2")
	//Fatalf("arg1 %d", 1)

	err := funcWithError()
	CheckError(err)
}

func TestSimpleFileLog(t *testing.T) {
	SetFileOutput("./test.log", false)

	DebugMode = true
	PrintTimestamp = true
	PrintColors = true
	TimeFormat = "2006-01-02 15:04:05.000"

	myVar := map[string]string{"hello": "world"}

	Debug("arg1", "arg2")
	Debugf("arg1 %d", 1)
	DebugDump(myVar, "prefix")
	DebugSeparator("title")

	Info("arg1", "arg2")
	Infof("arg1 %d", 1)
	InfoDump(myVar, "prefix")
	InfoSeparator("title")

	Warn("arg1", "arg2")
	Warnf("arg1 %d", 1)
	WarnDump(myVar, "prefix")
	WarnSeparator("title")

	Error("arg1", "arg2")
	Errorf("arg1 %d", 1)
	ErrorDump(myVar, "prefix")
	ErrorSeparator("title")

	//Fatal("arg1", "arg2")
	//Fatalf("arg1 %d", 1)

	err := funcWithError()
	CheckError(err)
}

func TestRotateFileLog(t *testing.T) {
	SetFileOutput("./test.log", true)

	DebugMode = true
	PrintTimestamp = true
	PrintColors = true
	TimeFormat = "2006-01-02 15:04:05.000"

	myVar := map[string]string{"hello": "world"}

	Debug("arg1", "arg2")
	Debugf("arg1 %d", 1)
	DebugDump(myVar, "prefix")
	DebugSeparator("title")

	Info("arg1", "arg2")
	Infof("arg1 %d", 1)
	InfoDump(myVar, "prefix")
	InfoSeparator("title")

	Warn("arg1", "arg2")
	Warnf("arg1 %d", 1)
	WarnDump(myVar, "prefix")
	WarnSeparator("title")

	Error("arg1", "arg2")
	Errorf("arg1 %d", 1)
	ErrorDump(myVar, "prefix")
	ErrorSeparator("title")

	//Fatal("arg1", "arg2")
	//Fatalf("arg1 %d", 1)

	err := funcWithError()
	CheckError(err)
}

func funcWithError() error {
	err := errors.New("this is a error")
	StackTrace(err)
	return err
}

func resetLogConfig() {
	PrintTimestamp = true
	PrintColors = true
	DebugMode = false
	TimeZone, _ = time.LoadLocation("Europe/Brussels")
	TimeFormat = TestingTimeFormat
}

func redirectOutput() (*bytes.Buffer, *bytes.Buffer) {
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	Stdout = stdout
	Stderr = stderr
	return stdout, stderr
}

func resetLogOutput() {
	Stdout = os.Stdout
	Stderr = os.Stderr
	TimeFormat = DefaultTimeFormat
}
