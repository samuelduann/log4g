package log4g

import (
	"fmt"
	"io"
	"os"
	"sync"
    "strings"
	"time"
    "path/filepath"
)

const (
	FilenameSuffixInDay = "20060102"
	FilenameSuffixInHour = "2006010215"
	FilenameSuffixInMinute = "200601021504"
	FilenameSuffixInSecond = "20060102150405"
	StandardLogPrefix = "2006/01/02 15:04:05"
)

const (
	debugLevel = iota
	infoLevel
	noticeLevel
	warningLevel
	errorLevel
)

var (
	logPrefix = map[int]string{
		debugLevel:   "DEBUG",
		infoLevel:    "INFO",
		noticeLevel:  "NOTICE",
		warningLevel: "WARN",
		errorLevel:   "ERROR",
	}
)

type Logger struct {
	sync.Mutex
	filenamePrefix        string
	filenameSuffixFormat  string
	currentFilenameSuffix string
	fileWriter            io.WriteCloser
}

func NewLogger(filenamePrefix string, filenameSuffixFormat string) *Logger {
	l := &Logger{
		filenamePrefix:        filenamePrefix,
		filenameSuffixFormat:  filenameSuffixFormat,
		currentFilenameSuffix: "",
	}
    return l
}

func (l *Logger) updateInnerLogger(now time.Time) {
	filenameSuffix := now.Format(l.filenameSuffixFormat)
	if filenameSuffix != l.currentFilenameSuffix {
        if l.fileWriter != nil && l.fileWriter != os.Stdout {
		    l.fileWriter.Close()
        }
        os.Rename(l.filenamePrefix, fmt.Sprintf("%s.%s", l.filenamePrefix, l.currentFilenameSuffix))
		if l.fileWriter = l.getFileWriter(l.filenamePrefix); l.fileWriter == os.Stdout {
            fmt.Fprintf(os.Stderr, "log4g init with prefix:%s failed, log to stdout\n", l.filenamePrefix)
        }
		l.currentFilenameSuffix = filenameSuffix
	}
}

func (l *Logger) write(level int, format string, content ...interface{}) {
	now := time.Now()
	l.Lock()
	defer l.Unlock()
	l.updateInnerLogger(now)

	var s string
	if format == "" {
		s = fmt.Sprintf("%s [%s] %s\n", now.Format(StandardLogPrefix), logPrefix[level], fmt.Sprint(content...))
	} else {
		s = fmt.Sprintf("%s [%s] %s\n", now.Format(StandardLogPrefix), logPrefix[level], fmt.Sprintf(format, content...))
	}

	l.fileWriter.Write([]byte(s))
}

func (l *Logger) Info(content ...interface{}) {
	l.write(infoLevel, "", content...)
}

func (l *Logger) Infof(format string, content ...interface{}) {
	l.write(infoLevel, format, content...)
}

func (l *Logger) Warn(content ...interface{}) {
	l.write(warningLevel, "", content...)
}

func (l *Logger) Warnf(format string, content ...interface{}) {
	l.write(warningLevel, format, content...)
}

func (l *Logger) Notice(content ...interface{}) {
	l.write(noticeLevel, "", content...)
}

func (l *Logger) Noticef(format string, content ...interface{}) {
	l.write(noticeLevel, format, content...)
}

func (l *Logger) Debug(content ...interface{}) {
	l.write(debugLevel, "", content...)
}

func (l *Logger) Debugf(format string, content ...interface{}) {
	l.write(debugLevel, format, content...)
}

func (l *Logger) checkAndMkdir(filenamePrefix string) error {
    sep := string(filepath.Separator)
    if strings.Contains(filenamePrefix, sep) {
        lastIdx := strings.LastIndex(filenamePrefix, sep)
        logPath := filenamePrefix[:lastIdx + 1]
        return os.MkdirAll(logPath, 0755)
    }
    return nil
}

func (l *Logger) getFileWriter(filenamePrefix string) io.WriteCloser {
    if err := l.checkAndMkdir(filenamePrefix); err != nil {
        return os.Stdout
    }
	fileWriter, err := os.OpenFile(fmt.Sprintf("%s", filenamePrefix), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return fileWriter
}
