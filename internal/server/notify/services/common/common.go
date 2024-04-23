package common

type Notify struct {
	Context string `json:"context"`
	Content string `json:"content"`
}

type NotifyService interface {
	Notify(notify *Notify) (string, error)
}
