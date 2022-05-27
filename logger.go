package logger

import (
	"fmt"
	"os"
)

type LoggerType struct {
	LOG_DESTENATION string
}

func (L *LoggerType) SetLoggerFile(path string) {
	L.LOG_DESTENATION = path
}

func (L *LoggerType) Add(level string, message string, errorMessage error) {
	if level == "" {
		level = "INFO"
	}
	var logLine string
	if errorMessage != nil {
		logLine = "ERROR | " + message + " : " + errorMessage.Error()
		fmt.Println(logLine)
	} else {
		logLine = level + " | " + message
		fmt.Println(logLine)
	}
	f, err := os.OpenFile(L.LOG_DESTENATION, os.O_CREATE|os.O_APPEND, 0600)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Ошибка открытия файла: ", err)
		}
	}(f)

	if err != nil {
		fmt.Println("Ошибка открытия файла лога")
	}

	_, err = f.WriteString(logLine + "\n")
	if err != nil {
		fmt.Println("Ошибка записи лога")
	}
}
