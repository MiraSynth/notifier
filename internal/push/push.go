package push

type PushData struct {
	Context string
	Title   string
	Text    string
	Icon    string
}

type Push interface {
	Push(data *PushData) (string, error)
	PushCritical(data *PushData) (string, error)
}

type pushService struct {
}

func NewPushService() Push {
	var p Push = &pushService{}
	return p
}
