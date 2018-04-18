package errorlog

import (
	"model"
	"log"
	"os"
	"props"
)

var logger *log.Logger

func init() {

	file, err := os.OpenFile(props.GetLogFileLocation(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		logger = log.New(file, "ServiceQ: ", log.Ldate | log.Ltime)
	}
}

func IncrementErrorCount(sqp *model.ServiceQProperties, service string, errType int, errReason string) {

	(*sqp).REMutex.Lock()
	defer (*sqp).REMutex.Unlock()
	(*sqp).RequestErrorLog[service] += 1
	logServiceError(service, errType, errReason)
}

func ResetErrorCount(sqp *model.ServiceQProperties, service string) {

	(*sqp).REMutex.Lock()
	defer (*sqp).REMutex.Unlock()
	(*sqp).RequestErrorLog[service] = 0
}

func logServiceError(service string, errType int, errReason string) {

	if logger != nil {
		logger.Printf("Error detected on %s [Code: %d, %s]", service, errType, errReason)
	}
}
