package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

const RouterCodeTemplate = `
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


func createRouterCode(projectName string)  {
	fileName := path.Join(projectName, "router", "router.go")
	err := ioutil.WriteFile(fileName, []byte(RouterCodeTemplate), os.ModePerm)
	if err != nil{
		log.Panic("can't write router.go")
	}
}