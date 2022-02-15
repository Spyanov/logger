package logger

import "fmt"

func Add(level string, message string, err error) {
	var error_message string
	if err != nil {
		error_message = err.Error()
		fmt.Println(level+" | "+message+" : ", error_message)
	} else {
		error_message = ""
		fmt.Println(level + " | " + message)
	}

}
