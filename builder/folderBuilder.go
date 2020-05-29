package builder

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"os"
	"os/exec"
)

type FolderBuilder struct {
	workingDir  string
	projectName string
	packages    []string
	folders     []string
}

func (f *FolderBuilder) ProjectName(name string) *FolderBuilder {
	f.projectName = name
	return f
}

func (f *FolderBuilder) WorkingDir(dir string) *FolderBuilder {
	f.workingDir = dir
	return f
}

func (f *FolderBuilder) Package(name string) *FolderBuilder {
	f.packages = append(f.packages, name)
	return f
}

func (f *FolderBuilder) Packages(names *[]string) *FolderBuilder {
	f.packages = append(f.packages, *names...)
	return f
}

func (f *FolderBuilder) Folder(name string) *FolderBuilder {
	f.folders = append(f.folders, name)
	return f
}

func (f *FolderBuilder) Folders(names []string) *FolderBuilder {
	f.folders = append(f.folders, names...)
	return f
}

func (f *FolderBuilder) Action() {

	err := os.Mkdir(f.projectName, os.ModePerm)
	//建立子資料夾
	for i := 0; i < len(f.folders); i++ {
		err := os.MkdirAll(f.folders[i], os.ModePerm)
		if err != nil {
			fmt.Println("can't create folder reason: ", f.folders[i])
		}
	}

	//執行go mod
	cmd := exec.Command("go", "mod", "init", f.projectName)
	cmd.Dir = f.workingDir
	err = cmd.Run()

	if err != nil {
		fmt.Println("can't exec go mod reason: ", err)
	}

	bar := pb.StartNew(len(f.packages))
	for i := 0; i < len(f.packages); i++ {
		process := exec.Command("go", "get", f.packages[i])
		process.Dir = f.workingDir
		process.Run()
		bar.Increment()
	}
	bar.Finish()
	fmt.Println("get go mod package finish")
}
