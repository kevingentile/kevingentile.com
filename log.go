package main

import (
	"log"
	"os"
)

var stdOutLogger = log.New(os.Stdout, "", log.Ltime)
var debugLogger = log.New(os.Stderr, "kg DEBUG: ", log.Ltime)
var errorLogger = log.New(os.Stderr, "kg ERROR: ", log.Ltime)

func Print(v ...interface{}) {
	stdOutLogger.Println(v...)
}

func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

func Fatal(v ...interface{}) {
	stdOutLogger.Println(v...)
	os.Exit(1)
}
