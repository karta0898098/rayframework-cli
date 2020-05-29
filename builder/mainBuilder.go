package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type MainBuilder struct {
	Template    string
	ProjectName string
	Path        string
	File        string
	SessionCode string
}

func (c *MainBuilder) Action() {
	fileName := path.Join(c.ProjectName, c.Path, c.File)
	err := ioutil.WriteFile(fileName, []byte(fmt.Sprintf(c.Template, c.ProjectName, c.ProjectName, c.ProjectName, c.SessionCode, c.ProjectName)), os.ModePerm)
	if err != nil {
		fmt.Println("can't write:", c.File)
	}

	fmt.Printf("create file %s success \n", c.File)
}
