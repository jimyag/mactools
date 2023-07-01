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
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
)

type Log struct {
	level Level
	log   *log.Logger
}

var (
	Logger *Log
)

func init() {
	stdLog := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	Logger = &Log{
		level: DebugLevel,
		log:   stdLog,
	}

}

func SetLevel(level Level) {
	Logger.level = level
}

func Debug(format string, v ...interface{}) {
	if Logger.level > DebugLevel {
		return
	}
	format = "[DEBUG] " + format
	Logger.log.Output(2, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	if Logger.level > InfoLevel {
		return
	}
	format = "[INFO] " + format
	Logger.log.Output(2, fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	if Logger.level > ErrorLevel {
		return
	}
	format = "[ERROR] " + format
	Logger.log.Output(2, fmt.Sprintf(format, v...))
}
