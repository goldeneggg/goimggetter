package imggetter

import (
	"log"
	"os"
)

type igLogger interface {
	err(v ...interface{})
	debug(v ...interface{})
}

type igLog struct {
	errorLogger *log.Logger
	debugLogger *log.Logger
	isDebug     bool
}

func (il *igLog) err(v ...interface{}) {
	il.errorLogger.Println(v)
}

func (il *igLog) debug(v ...interface{}) {
	if il.isDebug {
		il.debugLogger.Println(v)
	}
}

var l igLogger

func prepareLogger(debug bool) {
	l = &igLog{errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags), debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags), isDebug: debug}
}
