package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func getLevel(prefix string) int {
	for i, flag := range levelFlags {
		if flag == prefix {
			return i
		}
	}
	return 0
}

func Debug(v ...interface{}) {
	if levelFlags[DEBUG] == DefaultPrefix {
		setPrefix(DEBUG)
		logger.Println(v...)
	}
}

func Info(v ...interface{}) {
	if INFO >= getLevel(DefaultPrefix) {
		setPrefix(INFO)
		logger.Println(v...)
	}
}

func Warn(v ...interface{}) {
	if WARN >= getLevel(DefaultPrefix) {
		setPrefix(WARN)
		logger.Println(v...)
	}
}

func Error(v ...interface{}) {
	if ERROR >= getLevel(DefaultPrefix) {
		setPrefix(ERROR)
		logger.Println(v...)
	}
}

func Fatal(v ...interface{}) {
	if FATAL >= getLevel(DefaultPrefix) {
		setPrefix(FATAL)
		logger.Println(v...)
	}
}

func setPrefix(level int) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
