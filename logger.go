package logger

import (
	"fmt"
	"os"
)

var LOG_DESTENATION string = ""

func Add(level string, message string, err error) {
	var error_message string
	if err != nil {
		error_message = err.Error()
		fmt.Println(level+" | "+message+" : ", error_message)
	} else {
		error_message = ""
		fmt.Println(level + " | " + message)
	}
	f, err := os.Open(LOG_DESTENATION)
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
