package util

import "fmt"

func CheckErr(e error)  {
	if e != nil {
		panic(e)
	}
}

func FCheckErr(e error, format string) {
	if e != nil {
		panic(fmt.Errorf(format, e))
	}
}
