package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func GetEnvFlag(key string) bool {
	value, isPresent := os.LookupEnv(fmt.Sprintf("kami.%v", strings.ToLower(key)))
	if isPresent {
		b, _ := strconv.ParseBool(value)
		return b
	}
	return false
}