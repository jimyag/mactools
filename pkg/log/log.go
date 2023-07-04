/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package log

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l Level) Color(s string) string {
	switch l {
	case LevelTrace:
		return fmt.Sprintf("\033[34m%s\033[0m", s)
	case LevelDebug:
		return fmt.Sprintf("\033[36m%s\033[0m", s)
	case LevelInfo:
		return fmt.Sprintf("\033[32m%s\033[0m", s)
	case LevelWarn:
		return fmt.Sprintf("\033[33m%s\033[0m", s)
	case LevelError:
		return fmt.Sprintf("\033[31m%s\033[0m", s)
	default:
		return s
	}
}

type Log struct {
	level Level
	log   *log.Logger
}

type LogItem struct {
	Level   Level
	Message string
}

func (i *LogItem) String() string {
	return fmt.Sprintf("[%s] %s", i.Level.Color(i.Level.String()), i.Message)
}

func (l *Log) SetLevel(level Level) {
	l.level = level
}

func (l *Log) Trace(format string, v ...interface{}) {
	l.Log(LevelTrace, format, v...)
}

func (l *Log) Debug(format string, v ...interface{}) {
	l.Log(LevelDebug, format, v...)
}

func (l *Log) Info(format string, v ...interface{}) {
	l.Log(LevelInfo, format, v...)
}

func (l *Log) Warn(format string, v ...interface{}) {
	l.Log(LevelWarn, format, v...)
}

func (l *Log) Error(format string, v ...interface{}) {
	l.Log(LevelError, format, v...)
}

func (l *Log) Log(level Level, format string, v ...interface{}) {
	if l.level > level {
		return
	}
	logItem := &LogItem{
		Level:   level,
		Message: fmt.Sprintf(format, v...),
	}
	_ = l.log.Output(4, logItem.String())
}

var (
	DefaultLogger *Log
)

func init() {
	DefaultLogger = &Log{
		level: LevelDebug,
		log:   log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile),
	}
}

func SetLevel(level Level) {
	DefaultLogger.level = level
}

func Trace(format string, v ...interface{}) {
	DefaultLogger.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	DefaultLogger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	DefaultLogger.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	DefaultLogger.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	DefaultLogger.Error(format, v...)
}
