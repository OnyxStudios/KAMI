package editor

import (
	"kami/render/screen"
	"kami/util"
	"kami/window/editor/stage"
)

func Display(stage screen.Stage) {
	if currentStage != nil {
		defer currentStage.Dispose()
	}
	currentStage = stage
	if currentStage != nil {
		util.Log.Printf("Loading Stage %v", currentStage.Name())
		currentStage.Load()
	}
}

func Refresh() {
	if currentStage != nil {
		currentStage.Dispose()
		currentStage.Load()
	}
}

var (
	welcome = stage.Welcome{}
	test = stage.TestModel{}
)

func DisplayWelcomeScreen() {
	Display(welcome)
}

func DisplayTestStage() {
	Display(test)
}