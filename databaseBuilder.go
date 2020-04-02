package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const DatabaseCodeTemplate = `
package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"%s/config"
	"time"
)

var Context *gorm.DB

func NewDatabase(conf config.DataBaseConfig) {
	var err error

	Context, err = gorm.Open("mysql", conf.GetConnString())
	if err != nil {
		log.Panic("[database] connect to database failed", err)
		return
	}

	err = Context.DB().Ping()
	if err != nil {
		log.Panic("[database] connect to database failed", err)
		return
	}

	Context.DB().SetMaxOpenConns(10)
	Context.DB().SetMaxIdleConns(5)
	Context.DB().SetConnMaxLifetime(time.Second * 60)
}

func CloseDataBase() {
	_ = Context.Close()
}`

func createDatabaseCode(projectName string) {
	fileName := path.Join(projectName, "database", "database.go")
	err := ioutil.WriteFile(fileName, []byte(fmt.Sprintf(DatabaseCodeTemplate, projectName)), os.ModePerm)
	if err != nil {
		log.Panic("can't write database.go")
	}
}
