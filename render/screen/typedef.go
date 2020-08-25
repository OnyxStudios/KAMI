package screen

import "github.com/go-gl/glfw/v3.3/glfw"

type Renderable interface {
	Draw()
	Dispose()
}

type Stage interface {
	Draw(window *glfw.Window, delta float32)
	Dispose()
	Load()
	Name() string
}