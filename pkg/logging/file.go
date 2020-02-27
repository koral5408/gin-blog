package logging

import (

	"time"
	"fmt"

	"gin-blog/pkg/setting"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s",setting.AppSetting.RuntimeRootPath,setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

/*
var (
	LogSavePath = "runtime/log/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName,time.Now().Format(TimeFormat),LogFileExt)

	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch{
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Fail to OpenFile :%v",err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v",err)
	}

	return handle
}

func mkDir () {
	dir, _ := os.Getwd() //获取的是哪的地址？本文件所在的文件夹的地址么
	err := os.MkdirAll(dir + "/" +getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
 */