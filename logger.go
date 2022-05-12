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

func (L *LoggerType) Add(level string, message string, err error) {
	var error_message string
	if err != nil {
		error_message = err.Error()
		fmt.Println(level+" | "+message+" : ", error_message)
	} else {
		error_message = ""
		fmt.Println(level + " | " + message)
	}
	f, err := os.OpenFile(L.LOG_DESTENATION, os.O_APPEND|os.O_WRONLY, 0600)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Ошибка открытия файла: ", err)
		}
	}(f)

	if err != nil {
		fmt.Println("Ошибка открытия файла лога")
	}
	_, err = f.WriteString(level + " | " + message + " : " + error_message)
	if err != nil {
		fmt.Println("Ошибка записи лога")
	}
}
