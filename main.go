package main

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"kami/constants"
	"kami/render"
	"kami/render/models/obj"
	"kami/util"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

var (
	fov float32 = 45.0
	nearPlane float32 = 0.1
	farPlane float32 = 1000
	mouseX, mouseY float32
	deltaX, deltaY float32
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
	window, err := glfw.CreateWindow(constants.WindowWidth, constants.WindowHeight, fmt.Sprintf("%v %v", constants.Title, constants.Version), nil, nil)
	util.FCheckErr(err, "could not create OpenGL window: %v")
	window.MakeContextCurrent() //create openGL context
	glfw.SwapInterval(1)
	render.InitGL()
	window.SetSizeCallback(func(window *glfw.Window, width int, height int) {
		frameWidth, frameHeight := window.GetFramebufferSize()
		render.MainCamera.UpdateProjectionMatrix(fov, float32(frameWidth), float32(frameHeight), nearPlane, farPlane)
	})
	window.SetMaximizeCallback(func(window *glfw.Window, iconified bool) {
		frameWidth, frameHeight := window.GetFramebufferSize()

		if !iconified {
			vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
			window.SetPos((vidMode.Width - frameWidth) / 2, (vidMode.Height - frameHeight) / 2)
		}

		render.MainCamera.UpdateProjectionMatrix(fov, float32(frameWidth), float32(frameHeight), nearPlane, farPlane)
	})

	window.SetCursorPosCallback(func(window *glfw.Window, xPos, yPos float64) {
		width, height := window.GetSize()
		deltaX = float32(width) / 2 - float32(xPos)
		deltaY = float32(height) / 2 - float32(yPos)
		mouseX = float32(xPos)
		mouseY = float32(yPos)
	})

	//TODO load resources here
	cubeModel := obj.LoadModel("models/triangulated_cube.obj")
	frameWidth, frameHeight := window.GetFramebufferSize()
	render.LoadShaders()
	render.MainCamera.UpdateProjectionMatrix(fov, float32(frameWidth), float32(frameHeight), nearPlane, farPlane)
	render.DefaultShaderProgram.SetAttribLocation(0, "position")
	render.DefaultShaderProgram.SetAttribLocation(1, "textureCoords")
	render.DefaultShaderProgram.SetAttribLocation(2, "normal")
	transformationMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("transformationMatrix")
	viewMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("viewMatrix")

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	lastTime := glfw.GetTime()
	texture := util.LoadTexture("textures/test.png")
	gl.Enable(gl.DEPTH_TEST)

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

			// Do OpenGL stuff.
			render.DefaultShaderProgram.UseShader()
			viewMatrix := render.CreateViewMatrix(render.MainCamera.Position, render.MainCamera.Rotation)
			gl.UniformMatrix4fv(viewMatrixUniform, 1, false, &viewMatrix[0])

			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, texture)

			for index, element := range cubeModel.Parts {
				if window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
					cubeModel.Parts[index].Rotation.V[0] += -deltaY * elapsedTime / 16
					cubeModel.Parts[index].Rotation.V[1] += -deltaX * elapsedTime / 16
				}
				cubeModel.Parts[index].Position[2] = -10
				transformMatrix := render.CreateTransformMatrix(element.Position, element.Rotation, 1)

				element.Vao.Bind()
				gl.UniformMatrix4fv(transformationMatrixUniform, 1, false, &transformMatrix[0])
				gl.DrawElements(gl.TRIANGLES, int32(len(element.Indices)), gl.UNSIGNED_INT, gl.Ptr(element.Indices))
			}
			window.SwapBuffers()
		}
		glfw.PollEvents()
	}
}
