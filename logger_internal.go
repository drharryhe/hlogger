package hlogger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var colors = map[string]*color.Color{
	"DEBUG": color.New(color.FgHiBlack),
	"INFO ": color.New(color.FgHiGreen),
	"WARN ": color.New(color.FgHiYellow),
	"ERROR": color.New(color.FgHiRed),
	"FATAL": color.New(color.FgHiRed),
}

var mutex sync.Mutex

func init() {
	TimeZone, _ = time.LoadLocation(DefaultTimeZone)
	DebugMode = os.Getenv("DEBUG") == "1"
	PrintTimestamp = os.Getenv("PRINT_TIMESTAMP") == "1"

	color.NoColor = false

}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}

func formatSeparator(message string, separator string, length int) string {
	if message == "" {
		return strings.Repeat(separator, length)
	}
	prefix := strings.Repeat(separator, 4)
	suffixLength := length - len(message) - len(prefix) - 4
	suffix := ""
	if suffixLength > 0 {
		suffix = strings.Repeat(separator, suffixLength)
	}
	return prefix + "<" + message + ">" + suffix
}

func printMessage(level string, message string) {
	level = strings.ToUpper(level)

	if PrintTimestamp {
		message = addTimestampToMessage(level, message)
	}

	if PrintColors && runtime.GOOS != "windows" && FileOutputName == "" {
		lines := splitInLines(message)
		for _, line := range lines {
			printColoredMessage(level, line)
		}
	} else {
		printNonColoredMessage(level, message)
	}

}

func printNonColoredMessage(level string, message string) {
	w := writerForLevel(level)
	w.Write([]byte(message + "\n"))
}

func printColoredMessage(level string, message string) {
	w := writerForLevel(level)
	if c, ok := colors[level]; ok {
		mutex.Lock()
		defer mutex.Unlock()
		c.EnableColor()
		c.Fprint(w, message)
		w.Write([]byte("\n"))
	}
}

func addTimestampToMessage(level string, message string) string {

	tstamp := time.Now()
	if TimeZone != nil {
		tstamp = tstamp.In(TimeZone)
	}

	formattedTime := tstamp.Format(TimeFormat)
	message = formattedTime + " [" + level + "] " + message

	return message

}

func writerForLevel(level string) io.Writer {
	if FileOutputName != "" {
		if RotateFile {
			if isNewDay() {
				Stdout = newFileLogWriter()
			}
		}
		return Stdout
	} else {
		if level == "ERROR" || level == "FATAL" {
			return Stderr
		}
		return Stdout
	}
}

func splitInLines(text string) []string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\r", "\n", -1)
	lines := strings.Split(text, "\n")
	return lines
}

func newFileLogWriter() io.Writer {
	currentDate := time.Now().Format("2006-01-02")

	var fileName string
	if RotateFile {
		ss := strings.Split(FileOutputName, ".")
		if l := len(ss); l == 1 {
			fileName = fmt.Sprintf("%s_%s", FileOutputName, currentDate)
		} else {
			fileName = fmt.Sprintf("%s_%s.%s", strings.Join(ss[0:l-1], "."), currentDate, ss[l-1])
		}
	} else {
		fileName = FileOutputName
	}

	// 创建文件并写入内容
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("failed to create log file:%v", err)
		os.Exit(-1)
	}

	return file
}

func isNewDay() bool {
	if day := time.Now().Format("2006-01-02"); day == CurrentDay {
		return false
	} else {
		CurrentDay = day
		return true
	}
}
