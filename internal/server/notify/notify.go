package notify

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"mirasynth.stream/notifier/internal/server/notify/services"
	"mirasynth.stream/notifier/internal/server/notify/services/common"
)

func post(c *gin.Context) {
	input := common.Notify{}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = json.Unmarshal(jsonData, &input)
	if err != nil {
		c.JSON(500, err)
		return
	}

	services.NotifyServices(&input)
}

func RegisterController(ge *gin.RouterGroup, rootpath string) {
	services.StartServices()
	ge.POST(rootpath, post)
}
