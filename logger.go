package logger

import (
	"fmt"
	"os"
)

type logger struct {
	LOG_DESTENATION string
}

func (l *logger) SetLoggerFile(path string) {
	l.LOG_DESTENATION = path
}

func (l *logger) Add(level string, message string, err error) {
	var error_message string
	if err != nil {
		error_message = err.Error()
		fmt.Println(level+" | "+message+" : ", error_message)
	} else {
		error_message = ""
		fmt.Println(level + " | " + message)
	}
	f, err := os.Open(l.LOG_DESTENATION)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

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
