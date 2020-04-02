package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	nowVersion := 0.5

	var h bool
	var v bool
	var projectName string

	flag.BoolVar(&h, "h",false,"this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&projectName, "create", "", "create new project")
	flag.Usage = usage
	flag.Parse()

	if h {
		flag.Usage()
		os.Exit(0)
	}
	if v {
		fmt.Printf("rayframework-cli version: %v \n",nowVersion)
		os.Exit(0)
	}

	if projectName == ""{
		fmt.Println("please input project name")
		os.Exit(0)
	}

	fmt.Printf("create Project %s \n",projectName)
	createFolderTree(projectName)
	createConfigCode(projectName)
	createRouterCode(projectName)
	createDatabaseCode(projectName)
	createMainCode(projectName)
	fmt.Printf("create Project %s done \n",projectName)
}

func usage()  {
	flag.PrintDefaults()
}
