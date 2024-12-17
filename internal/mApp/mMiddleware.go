package mApp

import (
	"net/http"

	"MollyBlog/config"
	
	"github.com/gin-gonic/gin"
)

func (ma *MApp) AuthMiddleware(ctx *gin.Context) {
	if ctx.Request.URL.Path == config.MConfigInstance.UpdateEndpoint {
		secret := ctx.GetHeader("molly-secret")
		if secret != config.MConfigInstance.UpdateSecret {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}

	ctx.Next()
}
