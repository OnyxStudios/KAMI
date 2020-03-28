package main

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl64"
	"kami/constants"
	"kami/render"
	"kami/test"
	"kami/util"
	"runtime"
)

var cubeVertices = []float32{
	-1, -1, -1,
	1, -1, -1,
	1, 1, -1,
	-1, -1, 1,
	1, -1, 1,
	1, 1, 1,
	-1, 1, 1,
}

var cubeTextureCoords = []float32{
	0, 0,
	1, 0,
	1, 1,
	0, 1,
}

var cubeNormals = []float32{
	0, 0, 1,
	1, 0, 0,
	0, 0, -1,
	-1, 0, 0,
	0, 1, 0,
	0, -1, 0,
}

var cubeIndices = []int32{
	0, 1, 3, 3, 1, 2,
	1, 5, 2, 2, 5, 6,
	5, 4, 6, 6, 4, 7,
	4, 0, 7, 7, 0, 3,
	3, 2, 7, 7, 2, 6,
	4, 5, 0, 0, 5, 1,
}

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
		render.MainCamera.Projection = mgl64.Perspective(mgl64.DegToRad(45.0), float64(width/height), 0.1, 1000)
		test.LoadProjectionMatrix(&render.DefaultShaderProgram)
	})

	//TODO load resources here
	width, height := window.GetSize()
	render.LoadShaders()
	render.MainCamera.Projection = mgl64.Perspective(mgl64.DegToRad(45.0), float64(width/height), 0.1, 1000)
	test.LoadProjectionMatrix(&render.DefaultShaderProgram)
	render.DefaultShaderProgram.BindAttribLocation(0, "position")
	render.DefaultShaderProgram.BindAttribLocation(1, "textureCoords")
	render.DefaultShaderProgram.BindAttribLocation(2, "normal")
	transformationMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("transformationMatrix")
	viewMatrixUniform := render.DefaultShaderProgram.CreateUniformLocation("viewMatrix")
	lightPositionUniform := render.DefaultShaderProgram.CreateUniformLocation("lightPosition")
	lightColorUniform := render.DefaultShaderProgram.CreateUniformLocation("lightColor")
	environmentColorUniform := render.DefaultShaderProgram.CreateUniformLocation("environmentColor")

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {
		for err := gl.GetError(); err != gl.NO_ERROR; err = gl.GetError() {
			util.ErrLog.Println(fmt.Sprintf("OpenGL ERROR: %v", err))
		}

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
			gl.Uniform3f(lightColorUniform, 1, 1, 1)
			viewMatrix := render.CreateViewMatrix(render.MainCamera.Position, render.MainCamera.Rotation)
			gl.UniformMatrix4dv(viewMatrixUniform, 1, false, &viewMatrix[0])

			var vao uint32 //TODO replace with render.MakeVAO()
			gl.GenVertexArrays(1, &vao)
			gl.BindVertexArray(vao)

			test.BindIndices(36, cubeIndices)
			test.StoreDataInAttribs(0, 3, len(cubeVertices), cubeVertices, 0)
			test.StoreDataInAttribs(1, 2, len(cubeTextureCoords), cubeTextureCoords, 0)
			test.StoreDataInAttribs(2, 3, len(cubeNormals), cubeNormals, 0)
			gl.EnableVertexAttribArray(0)
			gl.EnableVertexAttribArray(1)
			gl.EnableVertexAttribArray(2)

			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, test.CubeTexture)

			//transformMatrix := render.CreateTransformMatrix(mgl64.Vec3{0, 0, -1}, mgl64.AnglesToQuat(0, 0, 0, mgl64.XYZ), 1)
			modelMatrix := mgl64.Ident4()

			gl.UniformMatrix4dv(transformationMatrixUniform, 1, false, &modelMatrix[0])
			gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, gl.Ptr(cubeIndices))

			window.SwapBuffers()
		}
		glfw.PollEvents()
	}
}
