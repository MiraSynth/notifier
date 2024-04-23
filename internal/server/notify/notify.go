package notify

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	"mirasynth.stream/notifier/internal/server/notify/services"
	"mirasynth.stream/notifier/internal/server/notify/services/common"
)

type NotifyController struct {
	web.Controller
}

func (ctrl *NotifyController) Post() {
	input := common.Notify{}

	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input)
	if err != nil {
		ctrl.Data["json"] = err.Error()
	}

	services.NotifyServices(&input)

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}

func RegisterController(rootpath string) *NotifyController {
	services.StartServices()

	ctrl := &NotifyController{}
	web.Router(rootpath, ctrl)

	return ctrl
}
