package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type CodeBuilder struct {
	Template    string
	ProjectName string
	Path        string
	File        string
}

func (c *CodeBuilder) Action() {
	fileName := path.Join(c.ProjectName, c.Path, c.File)
	err := ioutil.WriteFile(fileName, []byte(c.Template), os.ModePerm)
	if err != nil {
		fmt.Println("can't write:", c.File)
	}

	fmt.Printf("create file %s success \n", c.File)
}
