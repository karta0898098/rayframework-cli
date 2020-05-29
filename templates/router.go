package templates

const Router = `
package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	//TODO Register App Router or Register Api Router
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}`
