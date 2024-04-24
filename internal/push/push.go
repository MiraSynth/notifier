package push

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

type PushData struct {
	Context  string
	Title    string
	Text     string
	Icon     string
	Critical bool
}

type Push interface {
	Push(data *PushData) error
	PushCritical(data *PushData) error
}

type pushService struct {
}

func NewPushService() Push {
	var p Push = &pushService{}
	return p
}

func logError(data *PushData, urgency string, exitcode int, err error) {
	log.WithFields(log.Fields{
		"platformOS":   runtime.GOOS,
		"platformArch": runtime.GOARCH,
		"data":         data,
		"urgency":      urgency,
		"exitCode":     exitcode,
		"error":        err.Error(),
	}).Error(err)
}
