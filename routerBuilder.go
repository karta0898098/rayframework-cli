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
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	//TODO Register App Router or Register Api Router
}`

func createRouterCode(projectName string)  {
	fileName := path.Join(projectName, "router", "router.go")
	err := ioutil.WriteFile(fileName, []byte(RouterCodeTemplate), os.ModePerm)
	if err != nil{
		log.Panic("can't write router.go")
	}
}