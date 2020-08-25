package editor

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"kami/render/screen"
)

var (
	currentStage screen.Stage = nil
)

func Render(window *glfw.Window, delta float32) {
	currentStage.Draw(window, delta)
}
