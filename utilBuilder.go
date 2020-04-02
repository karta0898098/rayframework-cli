package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

const UtilCode = `
package util

import (
	"math/rand"
	"time"
)

func GeneratePassword(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

`

func createUtilCode(projectName string) {
	fileName := path.Join(projectName, "pkg", "util", "utiliy.go")
	err := ioutil.WriteFile(fileName, []byte(UtilCode), os.ModePerm)
	if err != nil {
		log.Panic("can't write router.go")
	}
}
