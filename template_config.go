package main

const TemplateConfig  =`
package config

import (
	"os"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	Database DataBaseConfig
	Runtime  RuntimeConfig
}

type DataBaseConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
	Port     string
}

type RuntimeConfig struct {
	Mode            string
	Port            string
}

func NewConfig(path string) Config {
	data, err := ini.Load(path)
	if err != nil {
		log.Panic("[config] load int data failed reason", err)
	}

	var config Config

	db := data.Section("database")
	config.Database = DataBaseConfig{
		User:     db.Key("user").String(),
		Password: db.Key("password").String(),
		Host:     db.Key("host").String(),
		DBName:   db.Key("name").String(),
		Port:     db.Key("port").String(),
	}

	runtime := data.Section("runtime")
	config.Runtime = RuntimeConfig{
		Mode:            runtime.Key("mode").String(),
		Port:            runtime.Key("port").String(),
	}

	log.Info("[config] load int data success")

	return config
}

func (d *DataBaseConfig) GetConnString() string {
	
	DBConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if DBConnectionString != "" {
		return DBConnectionString
	}

	if d.User == "" || d.Password == "" {
		return ""
	}

	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.Port, d.DBName)
}`

const DefaultIni  =`
[database]
user = ""
password = ""
host = 127.0.0.1
port = 3306
name = ""

[runtime]
mode = debug
port = :8080`
