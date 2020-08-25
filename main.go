package main

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"kami/constants"
	"kami/render"
	"kami/render/text"
	"kami/util"
	"kami/window/editor"
	"kami/window/mouse"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

var (
	fov            float32 = 45.0
	nearPlane      float32 = 0.1
	farPlane       float32 = 1000
)

func main() {
	err := glfw.Init()
	util.FCheckErr(err, "could not initialize glfw: %v")
	defer glfw.Terminate()
	glfw.DefaultWindowHints()
	glfw.WindowHint(glfw.AutoIconify, glfw.False)
	glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.FocusOnShow, glfw.True)
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 3) //opengl 3.3
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.Samples, 4)                            //4x FSAA
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) //We don't want the old OpenGL
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    //required for mac

	monitor := glfw.GetPrimaryMonitor()
	videoMode := monitor.GetVideoMode()
	glfw.WindowHint(glfw.RedBits, videoMode.RedBits)
	glfw.WindowHint(glfw.GreenBits, videoMode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, videoMode.BlueBits)
	glfw.WindowHint(glfw.RefreshRate, videoMode.RefreshRate)
	window, err := glfw.CreateWindow(constants.Config.WindowWidth, constants.Config.WindowHeight, fmt.Sprintf("%v %v", constants.Config.Title, constants.Config.Version), nil, nil)
	util.FCheckErr(err, "could not create OpenGL window: %v")
	window.MakeContextCurrent() //create openGL context
	glfw.SwapInterval(1)
	render.InitGL()
	sizeCallback := func(window *glfw.Window, width int, height int) {
		fWidth := float32(width)
		fHeight := float32(height)
		render.MainCamera.UpdateProjectionMatrix(fov, fWidth, fHeight, nearPlane, farPlane)
		text.SetWindowSize(fWidth, fHeight)
	}
	window.SetSizeCallback(sizeCallback)
	window.SetMaximizeCallback(func(window *glfw.Window, iconified bool) {
		frameWidth, frameHeight := window.GetFramebufferSize()

		if !iconified {
			vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
			window.SetPos((vidMode.Width-frameWidth)/2, (vidMode.Height-frameHeight)/2)
		}
		sizeCallback(window, frameWidth, frameHeight)
	})
	window.SetCursorPosCallback(mouse.GlfwCallback)

	text.LoadFonts()
	defer text.ReleaseFonts()
	w, h := window.GetSize()
	text.SetWindowSize(float32(w), float32(h))

	//TODO load resources here

	frameWidth, frameHeight := window.GetFramebufferSize()
	render.LoadShaders()
	render.MainCamera.UpdateProjectionMatrix(fov, float32(frameWidth), float32(frameHeight), nearPlane, farPlane)


	//gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.ClearColor(0.4, 0.4, 0.4, 0.0)


	editor.DisplayWelcomeScreen()
	//editor.DisplayTestStage() //TODO debug

	gl.Enable(gl.DEPTH_TEST)
	lastTime := glfw.GetTime()
	for !window.ShouldClose() {
		render.CheckGlError()
		time := glfw.GetTime()
		elapsedTime := float32(time - lastTime)
		lastTime = time

		//TODO process keybinds
		if window.GetAttrib(glfw.Focused) == glfw.True {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			if window.GetKey(glfw.KeyEscape) == glfw.Press {
				window.SetShouldClose(true)
			}

			editor.Render(window, elapsedTime)

			//w, h := window.GetSize()

			// Do OpenGL stuff.

			window.SwapBuffers()
		}
		glfw.PollEvents()
	}
}
