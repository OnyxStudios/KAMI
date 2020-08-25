package util

import "github.com/damien-lloyd/gltext"

//global debug switch
var debug = GetEnvFlag("debug")

func DebugEnabled() bool {
	return debug
}


func init()  {
	if DebugEnabled() {
		Log.Println("Debugging enabled")
	}
	gltext.IsDebug = GetEnvFlag("debug.gltext")
}