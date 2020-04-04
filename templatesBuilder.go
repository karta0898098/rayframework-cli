package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

const IndexTemplate  = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>RayFramework</title>
</head>
<body>
    <p>Welcome to RayFramework</p>
</body>
</html>
`

func createIndexTemplates(projectName string)  {
	fileName := path.Join(projectName, "templates", "index.html")
	err := ioutil.WriteFile(fileName, []byte(IndexTemplate), os.ModePerm)
	if err != nil{
		log.Panic("can't write index.html")
	}
}