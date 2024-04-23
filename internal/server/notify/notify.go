package notify

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type NotifyController struct {
	web.Controller
}

type notify struct {
	Context string `json:"context"`
	Content string `json:"content"`
}

func (ctrl *NotifyController) Post() {
	input := notify{}

	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input)
	if err != nil {
		ctrl.Data["json"] = err.Error()
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}

func RegisterController(rootpath string) *NotifyController {
	ctrl := &NotifyController{}
	web.Router(rootpath, ctrl)

	return ctrl
}
