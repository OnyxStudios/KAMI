package main

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"kami/constants"
	"kami/render"
	"kami/util"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

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
	window.SetMaximizeCallback(func(window *glfw.Window, iconified bool) {
		width, height := window.GetSize()
		screenWidth, screenHeight := glfw.GetPrimaryMonitor().GetPhysicalSize()

		render.MainCamera.Projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width/height), 0.1, 1000)
		render.LoadProjectionMatrix(&render.DefaultShaderProgram, render.MainCamera)
		gl.Viewport(0, 0, int32(width), int32(height))

		if !iconified {
			window.SetPos((width-screenWidth)/2, (height-screenHeight)/2)
		}
	})

	//TODO load resources here
	cubeModel := render.NewModel("models/cube.obj")
	width, height := window.GetSize()
	render.LoadShaders()
	render.MainCamera.Projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width/height), 0.1, 1000)
	gl.Viewport(0, 0, int32(width), int32(height))
	render.LoadProjectionMatrix(&render.DefaultShaderProgram, render.MainCamera)
	render.DefaultShaderProgram.SetAttribLocation(0, "position")
	render.DefaultShaderProgram.SetAttribLocation(1, "textureCoords")
	render.DefaultShaderProgram.SetAttribLocation(2, "normal")
	transformationMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("transformationMatrix")
	viewMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("viewMatrix")
	lightPositionUniform := render.DefaultShaderProgram.CreateUniformLocation("lightPosition")
	lightColorUniform := render.DefaultShaderProgram.CreateUniformLocation("lightColor")
	environmentColorUniform := render.DefaultShaderProgram.CreateUniformLocation("environmentColor")

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	cubeVAO := render.VertexArrayObject{BufferCount:3}
	render.LoadVAO(&cubeVAO)
	cubeVAO.Bind()
	cubeVAO.AddAttribData(0, 3, cubeModel.GetVertices(), 0, 0)
	cubeVAO.AddAttribData(1, 2, cubeModel.GetTextureCoords(), 0, 0)
	cubeVAO.AddAttribData(2, 3, cubeModel.GetNormals(), 0, 0)
	gl.BindVertexArray(0)

	lastTime := glfw.GetTime()
	angle := 0.0
	texture := util.LoadTexture("textures/planks.png")
	gl.Enable(gl.DEPTH_TEST)
	for !window.ShouldClose() {
		render.CheckGlError()
		time := glfw.GetTime()
		elapsedTime := time - lastTime
		lastTime = time
		angle += elapsedTime

		//TODO process keybinds
		if window.GetAttrib(glfw.Focused) == glfw.True {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			if window.GetKey(glfw.KeyEscape) == glfw.Press {
				window.SetShouldClose(true)
			}

			// Do OpenGL stuff.
			render.DefaultShaderProgram.UseShader()
			gl.Uniform3f(environmentColorUniform, 0.078, 0.078, 0.078)
			gl.Uniform3f(lightPositionUniform, 0, 10, 0)
			gl.Uniform3f(lightColorUniform, 0.69, 0.90, 1)
			viewMatrix := render.CreateViewMatrix(render.MainCamera.Position, render.MainCamera.Rotation)
			gl.UniformMatrix4fv(viewMatrixUniform, 1, false, &viewMatrix[0])
			cubeVAO.Bind()

			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, texture)
			transformMatrix := render.CreateTransformMatrix(mgl32.Vec3{0, 0, -10}, mgl32.AnglesToQuat(0, float32(angle), 0, mgl32.XYZ), 1)
			gl.UniformMatrix4fv(transformationMatrixUniform, 1, false, &transformMatrix[0])
			gl.DrawElements(gl.TRIANGLES, int32(len(cubeModel.VecIndices)), gl.UNSIGNED_INT, gl.Ptr(cubeModel.GetIndicesAsIntArr()))

			window.SwapBuffers()
		}
		glfw.PollEvents()
	}
}
