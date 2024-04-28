package ready

import (
	"github.com/gin-gonic/gin"
)

func RegisterController(r *gin.Engine, rootpath string) {
	r.GET(rootpath, func(c *gin.Context) {
		c.String(200, "ready")
	})
}
