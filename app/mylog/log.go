package mylog

import (
	"M-GateDBConfig/model"
	"fmt"
	"os"

	log "github.com/jeanphorn/log4go"
)

// ILog for Interfacing code using Log
type ILog interface {
	InitLog()
	Info()
	Debug()
	CAMessage(requestID, RC, TrxType, msg string, flownum int)
	MPAYMessage(requestID, RC, MTI, msg string, flownum int)
	Warning()
	Error(error)
	FatalError(error)
}

var (
	path = "./logs" // default log path
)

// InitLog to Initialization Log
func InitLog(mod model.BaseApplicationConfig) error {

	path = mod.LogFolder // configurable logs path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
	}
	log.LoadConfiguration("./logconfig.json")
	return nil
}

// MnDebug for Debugging in Simple Format
func MnDebug(msg string) {
	log.LOGGER("MiniDebug").Debug(msg)
	fmt.Println(msg)
}

// FlDebug for Debugggin in Full Format
func FlDebug(msg string, data interface{}) {
	log.LOGGER("FullDebug").Debug(msg, data)
}

// Error For Standard Error, Messaging Error, Fail Query, etc
func Error(e error, data interface{}) {
	log.LOGGER("Error").Error(e, data)
}

// FatalError for (DB connection lose, Load Configuration Failed, Listening Port Failed, Dependency failed)
func FatalError(e error) {
	log.LOGGER("FtlError").Error(e)
	panic(e)
}
