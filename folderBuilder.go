package main

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"os"
	"os/exec"
	"path"
)

func createFolderTree(projectName string) {
	nowPath, _ := os.Getwd()
	workingDir := path.Join(nowPath, projectName)

	packages :=[]string{
		"github.com/gin-gonic/gin",
		"github.com/gin-contrib/sessions",
		"github.com/gorilla/sessions",
		"github.com/jinzhu/gorm",
		"github.com/sirupsen/logrus",
		"gopkg.in/ini.v1",
	}

	err := os.Mkdir(projectName, os.ModePerm)
	err = os.MkdirAll(path.Join(projectName, "config"), os.ModePerm)
	err = os.MkdirAll(path.Join(projectName, "database"), os.ModePerm)
	err = os.MkdirAll(path.Join(projectName, "pkg", "util"), os.ModePerm)
	err = os.MkdirAll(path.Join(projectName, "router"), os.ModePerm)
	err = os.MkdirAll(path.Join(projectName, "templates"), os.ModePerm)

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = workingDir
	err = cmd.Run()

	fmt.Println("start get go mod package")
	fmt.Println("download dependency packages need some time")
	bar := pb.StartNew(len(packages))
	//建立Project 依賴的套件
	for i:=0;i<len(packages);i++{
		process := exec.Command("go", "get", packages[i])
		process.Dir = workingDir
		process.Run()
		bar.Increment()
	}
	bar.Finish()
	fmt.Println("get go mod package finish")

	if err != nil {
		panic("can't create project")
	}
}

