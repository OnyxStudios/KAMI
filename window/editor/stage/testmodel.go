package stage

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"kami/render"
	"kami/render/models/kami"
	"kami/render/models/obj"
	"kami/util"
	"kami/window/mouse"
)

type TestModel struct {

}

func (t TestModel) Name() string {
	return "Test Stage"
}

var (
	cubeModel kami.Model
	texture uint32
)

func (t TestModel) Draw(window *glfw.Window, delta float32) {
	render.DefaultShaderProgram.Shader.UseShader()
	viewMatrix := render.CreateViewMatrix(render.MainCamera.Position, render.MainCamera.Rotation)
	gl.UniformMatrix4fv(render.DefaultShaderProgram.ViewMatrix, 1, false, &viewMatrix[0])
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	deltaX, deltaY := mouse.GetMouseDelta()

	for index, element := range cubeModel.Parts {
		if window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
			cubeModel.Parts[index].Rotation.V[0] += -deltaY * delta / 50
			cubeModel.Parts[index].Rotation.V[1] += -deltaX * delta / 50
		}
		cubeModel.Parts[index].Position[2] = -10
		transformMatrix := render.CreateTransformMatrix(element.Position, element.Rotation, 6)

		element.Vao.Bind()
		gl.UniformMatrix4fv(render.DefaultShaderProgram.TransformMatrix, 1, false, &transformMatrix[0])

		gl.DrawElements(gl.TRIANGLES, int32(len(element.Indices)), gl.UNSIGNED_INT, gl.Ptr(element.Indices))
	}
}

func (t TestModel) Dispose() {
	var amount int32 = 1
	gl.DeleteTextures(amount, &texture)
}

func (t TestModel) Load() {
	cubeModel = obj.LoadModel("models/monkey.obj")
	texture = util.LoadTexture("textures/untextured.png")
}


