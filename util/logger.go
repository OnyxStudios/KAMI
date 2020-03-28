package util

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "[KAMI] ", 0)
var ErrLog = log.New(os.Stderr, "[KAMI] ", 0)
