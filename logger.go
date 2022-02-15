package logger

import "fmt"

func Add(level string, message string, err error) {
	fmt.Println(level+" | "+message+" : ", err)
}
