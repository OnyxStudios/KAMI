package main

import (
	"bytes"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl64"
	"image"
	"image/draw"
	"kami/constants"
	"kami/render"
	"kami/util"
	"runtime"
)

var (
	projectionMatrix mgl64.Mat4
	camera render.Camera
)


var cubeVertices = []float32 {
	-1, -1, -1,
	1, -1, -1,
	1, 1, -1,
	-1, -1, 1,
	1, -1, 1,
	1, 1, 1,
	-1, 1, 1,
}

var cubeTextureCoords = []float32 {
	0, 0,
	1, 0,
	1, 1,
	0, 1,
}

var cubeNormals = []float32 {
	0, 0, 1,
	1, 0, 0,
	0, 0, -1,
	-1, 0, 0,
	0, 1, 0,
	0, -1, 0,
}

var cubeIndices = []int32 {
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

	//TODO debug prints?

	glfw.DefaultWindowHints()
	//glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.AutoIconify, glfw.False)
	//glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.FocusOnShow, glfw.True)
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 3) //opengl 3.3
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.Samples, 4) //4x FSAA
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) //We don't want the old OpenGL
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) //required for mac

	monitor := glfw.GetPrimaryMonitor()
	videoMode := monitor.GetVideoMode()
	glfw.WindowHint(glfw.RedBits, videoMode.RedBits)
	glfw.WindowHint(glfw.GreenBits, videoMode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, videoMode.BlueBits)
	glfw.WindowHint(glfw.RefreshRate, videoMode.RefreshRate)
	window, err := glfw.CreateWindow(constants.WindowWidth, constants.WindowHeight, fmt.Sprintf("%v %v", constants.Title, constants.Version), nil, nil)
	util.FCheckErr(err, "could not create OpenGL window: %v")
	window.SetMaximizeCallback(func(window *glfw.Window, iconified bool) {
		width, height := window.GetSize()
		projectionMatrix = mgl64.Perspective(100, float64(width/height), 0.1, 1000)
		loadProjectionMatrix()
	})
	window.MakeContextCurrent() //create openGL context
	glfw.SwapInterval(1)
	render.InitGL()

	//TODO load resources here
	width, height := window.GetSize()
	projectionMatrix = mgl64.Perspective(100, float64(width/height), 0.1, 1000)
	loadProjectionMatrix()
	camera = render.Camera{Position: mgl64.Vec3{0, 0, 0}, Rotation: mgl64.AnglesToQuat(0, 0, 0, mgl64.XYZ)}

	gl.BindAttribLocation(render.DefaultShaderProgram.Handle, 0, gl.Str("position\x00"))
	gl.BindAttribLocation(render.DefaultShaderProgram.Handle, 1, gl.Str("textureCoords\x00"))
	gl.BindAttribLocation(render.DefaultShaderProgram.Handle, 2, gl.Str("normal\x00"))
	transformationMatrixUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("transformationMatrix\x00"))
	viewMatrixUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("viewMatrix\x00"))
	lightPositionUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("lightPosition\x00"))
	lightColorUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("lightColor\x00"))
	environmentColorUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("environmentColor\x00"))

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	render.LoadShaders()

	//backend.Load()
	for !window.ShouldClose() {
		//TODO process keybinds
		if window.GetAttrib(glfw.Focused) == glfw.True {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			if window.GetKey(glfw.KeyEscape) == glfw.Press {
				window.SetShouldClose(true)
			}

			// Do OpenGL stuff.
			gl.UseProgram(render.DefaultShaderProgram.Handle)
			gl.Uniform3f(environmentColorUniform, 0.078, 0.078, 0.078)
			gl.Uniform3f(lightPositionUniform, 0, 10, 0)
			gl.Uniform3f(lightColorUniform, 1, 1, 1)
			viewMatrix := render.CreateViewMatrix(camera.Position, camera.Rotation)
			gl.UniformMatrix4dv(viewMatrixUniform, 1, false, &viewMatrix[0])

			var vao uint32
			gl.GenVertexArrays(1, &vao)
			gl.BindVertexArray(vao)

			bindIndices(36, cubeIndices)
			storeDataInAttribs(0, 3, len(cubeVertices), cubeVertices, 0)
			storeDataInAttribs(1, 2, len(cubeTextureCoords), cubeTextureCoords, 0)
			storeDataInAttribs(2, 3, len(cubeNormals), cubeNormals, 0)
			gl.EnableVertexAttribArray(0)
			gl.EnableVertexAttribArray(1)
			gl.EnableVertexAttribArray(2)

			texture := loadTexture("textures/planks.png")
			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, texture)

			transformMatrix := render.CreateTransformMatrix(mgl64.Vec3{0, 0, -1}, mgl64.AnglesToQuat(0, 0, 0, mgl64.XYZ), 1)
			gl.UniformMatrix4dv(transformationMatrixUniform, 1, false, &transformMatrix[0])
			gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, gl.Ptr(cubeIndices))

			window.SwapBuffers()
		}
		glfw.PollEvents()
	}
}

//TODO move to designated file
func bindIndices(size int, data []int32) {
	var vbo uint32
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
}

//TODO move to designated file
func storeDataInAttribs(attribute uint32, coordSize int32, size int, data []float32, offset int) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attribute, coordSize, gl.FLOAT, false, 0, gl.PtrOffset(offset))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

//TODO move to designated file
func loadTexture(fileName string) uint32 {
	data := util.ReadAsset(fileName)
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return 0
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X * 4 {
		return 0
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	return texture
}

//TODO move to designated file
func loadProjectionMatrix() {
	projectionMatrixUniform := gl.GetUniformLocation(render.DefaultShaderProgram.Handle, gl.Str("projectionMatrix\x00"))
	gl.UniformMatrix4dv(projectionMatrixUniform, 1, false, &projectionMatrix[0])
}
