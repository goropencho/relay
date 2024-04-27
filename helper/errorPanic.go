package helper

import "fmt"

func ErrorPanic(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
