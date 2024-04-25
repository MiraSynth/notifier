package live

import (
	"github.com/beego/beego/v2/server/web"
)

type HealthLiveController struct {
	web.Controller
}

func (ctrl *HealthLiveController) Get() {
	ctrl.Ctx.WriteString("live")
	ctrl.Ctx.ResponseWriter.WriteHeader(200)
}

func RegisterController(rootpath string) *HealthLiveController {
	ctrl := &HealthLiveController{}
	web.Router(rootpath, ctrl)

	return ctrl
}
