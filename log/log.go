package log

import (
	log "github.com/jeanphorn/log4go"
)

// ILog for Interfacing code using Log
type ILog interface {
	InitLog()
	LogMessage()
	LogError()
	LogWarning()
}

// InitLog to Initialization Log
func InitLog() {
	log.LoadConfiguration("./example.json")
}
