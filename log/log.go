package log

import (
	"fmt"
	"log"
	"time"
	//"os"
)

const (
	LevelAny     = 0
	LevelDebug   = 1
	LevelInfo    = 2
	LevelWarning = 3
	LevelError   = 4
	LevelFatal   = 5
	LevelNone    = 6
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

var mapLevelString2Int = map[string]int{
	"all": LevelAny,

	"any":     LevelAny,
	"debug":   LevelDebug,
	"info":    LevelInfo,
	"warning": LevelWarning,
	"error":   LevelError,
	"fatal":   LevelFatal,
	"none":    LevelNone,

	"disabled": LevelNone,
}

func LevelString2Int(str string) int {
	level, ok := mapLevelString2Int[str]
	if ok {
		return level
	}

	return -1
}

//===========================================

type Logger struct {
	Level int
	Name  string
}

func NewLogger(name string) *Logger {
	return NewLoggerWithLevel(name, LevelWarning)
}

func NewLoggerWithLevel(name string, level int) *Logger {
	return &Logger{Name: buildLoggerName(name), Level: level}
}

func buildLoggerName(name string) string {
	return fmt.Sprintf("[%s] ", name)
}

var defaultlLogger = NewLogger("")

func SetDefaultLoggerName(name string) {
	defaultlLogger.Name = buildLoggerName(name)
}

func SetDefaultLoggerLevel(level int) {
	defaultlLogger.Level = level
}

//===========================================

func (logger *Logger) Debug(v ...interface{}) {
	if logger.Level > LevelDebug {
		return
	}
	log.Print(now("DEBUG"), logger.Name, fmt.Sprint(v...))
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	if logger.Level > LevelDebug {
		return
	}
	log.Print(now("DEBUG"), logger.Name, fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(v ...interface{}) {
	if logger.Level > LevelInfo {
		return
	}
	log.Print(now("INFO"), logger.Name, fmt.Sprint(v...))
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger.Level > LevelInfo {
		return
	}
	log.Print(now("INFO"), logger.Name, fmt.Sprintf(format, v...))
}

func (logger *Logger) Warning(v ...interface{}) {
	if logger.Level > LevelWarning {
		return
	}
	log.Print(now("WARNING"), logger.Name, fmt.Sprint(v...))
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	if logger.Level > LevelWarning {
		return
	}
	log.Print(now("WARNING"), logger.Name, " [WARNING] ", fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(v ...interface{}) {
	if logger.Level > LevelError {
		return
	}
	log.Print(now("ERROR"), logger.Name, fmt.Sprint(v...))
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger.Level > LevelError {
		return
	}
	log.Print(now("ERROR"), logger.Name, fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(v ...interface{}) {
	if logger.Level > LevelFatal {
		return
	}
	log.Fatal(now("FATAL"), logger.Name, fmt.Sprint(v...))
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	if logger.Level > LevelFatal {
		return
	}
	log.Fatal(now("FATAL"), logger.Name, fmt.Sprintf(format, v...))
}

//======================================

func Debug(v ...interface{}) {
	defaultlLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultlLogger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	defaultlLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	defaultlLogger.Infof(format, v...)
}

func Warning(v ...interface{}) {
	defaultlLogger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	defaultlLogger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	defaultlLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultlLogger.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	defaultlLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultlLogger.Fatalf(format, v...)
}

//=====================================

func now(level string) string {
	return fmt.Sprintf("[%s %s]", time.Now().Format("2006-01-02 15:04:05"), level)
}
