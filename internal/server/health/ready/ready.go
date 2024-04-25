package ready

import (
	"github.com/beego/beego/v2/server/web"
)

type HealthReadyController struct {
	web.Controller
}

func (ctrl *HealthReadyController) Get() {
	ctrl.Ctx.WriteString("ready")
	ctrl.Ctx.ResponseWriter.WriteHeader(200)
}

func RegisterController(rootpath string) *HealthReadyController {
	ctrl := &HealthReadyController{}
	web.Router(rootpath, ctrl)

	return ctrl
}
