package main

import (
	"fmt"
	"os"
	"path"

)

type AppBuilder struct {
	projectName string
	workingDir  string
	actions     []ExecActions
}

type ExecActions interface {
	Action()
}

func (b *AppBuilder) Name(name string) *AppBuilder {
	b.projectName = name
	return b
}

func (b *AppBuilder) WorkingDir() *AppBuilder {
	nowPath, _ := os.Getwd()
	b.workingDir = path.Join(nowPath, b.projectName)
	return b
}

func (b *AppBuilder) Folder() *AppBuilder {
	folderBuilder := &FolderBuilder{}
	folderBuilder = folderBuilder.ProjectName(b.projectName)
	folderBuilder = folderBuilder.WorkingDir(b.workingDir)
	folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "config"))
	folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "database"))
	folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "pkg", "util"))
	folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "router"))
	folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "templates"))

	if config.UseDockerCompose {
		folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "data"))
		folderBuilder = folderBuilder.Folder(path.Join(b.projectName, "sql"))
	}

	folderBuilder = folderBuilder.Packages(&[]string{
		"github.com/gin-gonic/gin",
		"github.com/gin-contrib/sessions",
		"github.com/gorilla/sessions",
		"github.com/jinzhu/gorm",
		"github.com/sirupsen/logrus",
		"gopkg.in/ini.v1",
		"github.com/go-redis/redis/v7",
	})
	b.actions = append(b.actions, folderBuilder)
	return b
}

func (b *AppBuilder) Router() *AppBuilder {
	Builder := &CodeBuilder{
		Template:    TemplateRouter,
		ProjectName: b.projectName,
		Path:        "router",
		File:        "router.go",
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) Util() *AppBuilder {
	Builder := &CodeBuilder{
		Template:    TemplateUtility,
		ProjectName: b.projectName,
		Path:        "pkg/util",
		File:        "utility.go",
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) Database() *AppBuilder {
	Builder := &CodeBuilder{
		Template:    TemplateDatabase,
		ProjectName: b.projectName,
		Path:        "database",
		File:        "database.go",
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) Templates() *AppBuilder {
	Builder := &CodeBuilder{
		Template:    TemplateIndex,
		ProjectName: b.projectName,
		Path:        "templates",
		File:        "index.html",
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) Docker() *AppBuilder {
	if config.UseDockerCompose {
		waitForItBuilder := &CodeBuilder{
			Template:    TemplateWaitForIt,
			ProjectName: b.projectName,
			File:        "wait-for-it.sh",
		}
		dockerBuilder := &CodeBuilder{
			Template:    TemplateDockerfileWait,
			ProjectName: b.projectName,
			File:        "dockerfile",
		}
		sqlBuilder := &CodeBuilder{
			Template:    fmt.Sprintf(TemplateInitdb, b.projectName),
			ProjectName: b.projectName,
			Path:        "sql",
			File:        "initdb.sql",
		}

		mysqlPassword := GeneratePassword(16)
		redisPassword := GeneratePassword(16)
		dockerComposeBuilder := &CodeBuilder{
			Template:    fmt.Sprintf(TemplateDockerCompose, mysqlPassword, redisPassword, redisPassword, mysqlPassword, b.projectName),
			ProjectName: b.projectName,
			File:        "docker-compose.yml",
		}

		b.actions = append(b.actions, dockerBuilder, waitForItBuilder, sqlBuilder, dockerComposeBuilder)
	} else {
		Builder := &CodeBuilder{
			Template:    TemplateDockerfile,
			ProjectName: b.projectName,
			File:        "dockerfile",
		}
		b.actions = append(b.actions, Builder)
	}
	return b
}

func (b *AppBuilder) DockerDB() *AppBuilder {
	Builder := &CodeBuilder{
		Template:    TemplateDockerfile,
		ProjectName: b.projectName,
		File:        "dockerfile",
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) Config() *AppBuilder {
	codeBuilder := &CodeBuilder{
		Template:    TemplateConfig,
		ProjectName: b.projectName,
		Path:        "config",
		File:        "config.go",
	}
	iniBuilder := &CodeBuilder{
		Template:    DefaultIni,
		ProjectName: b.projectName,
		File:        "config.ini",
	}
	b.actions = append(b.actions, codeBuilder, iniBuilder)
	return b
}

func (b *AppBuilder) Main() *AppBuilder {
	Builder := &MainBuilder{
		Template:    TemplateMain,
		ProjectName: b.projectName,
		File:        "main.go",
		SessionCode: GeneratePassword(8),
	}
	b.actions = append(b.actions, Builder)
	return b
}

func (b *AppBuilder) DockerMain() *AppBuilder {
	return b
}

func (b *AppBuilder) Build() {
	for i := 0; i < len(b.actions); i++ {
		b.actions[i].Action()
	}
}
