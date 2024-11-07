package helper

import "fmt"

func ErrorPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("Panic Error: %d", err))
	}
}
