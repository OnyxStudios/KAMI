package mouse

import "github.com/go-gl/glfw/v3.3/glfw"

var (
	mouseX, mouseY float32
	deltaX, deltaY float32
)

func GetMousePos() (float32, float32) {
	return mouseX, mouseY
}

func GetMouseDelta() (float32, float32) {
	return deltaX, deltaY
}

var GlfwCallback = func(window *glfw.Window, xPos, yPos float64) {
	width, height := window.GetSize()
	deltaX = float32(width)/2 - float32(xPos)
	deltaY = float32(height)/2 - float32(yPos)
	mouseX = float32(xPos)
	mouseY = float32(yPos)
}