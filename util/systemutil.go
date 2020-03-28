package util

import (
	"fmt"
)

func CheckErr(e error) {
	if e != nil {
		ErrLog.Fatalln(e)
	}
}

func FCheckErr(e error, format string) {
	if e != nil {
		ErrLog.Fatalln(fmt.Errorf(format, e))
	}
}
