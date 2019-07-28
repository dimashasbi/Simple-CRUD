package mylog

import (
	"M-GateDBConfig/model"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"os"
)

// ILog for Interfacing code using Log
type ILog interface {
	InitLog()
	Info()
	Debug()
	Message()
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

// Error For Standard Error
func Error(e error) {

}

// FatalError if Problem is so hard.
func FatalError(e error) {
	log.Error(e)
	// log.LOGGER("Error").Error(e)

	// matiin koneksi DB

	// matiin koneksi Websocket

	// matiin listen server
	panic(e)
}
