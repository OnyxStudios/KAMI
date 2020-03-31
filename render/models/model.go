package models

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"kami/render"
)

type Model struct {
	Parts []ModelPart
}

type ModelPart struct {
	Name string
	Position mgl32.Vec3
	Rotation mgl32.Quat

	Vao render.VertexArrayObject
	Vertices, TextureCoords, Normals []float32
	Indices []uint32
}

func (part *ModelPart) GenerateModelVAO() {
	vao := render.VertexArrayObject{BufferCount:3}
	render.LoadVAO(&vao)
	vao.Bind()
	vao.AddAttribData(0, 3, part.Vertices, 0, 0)
	vao.AddAttribData(1, 2, part.TextureCoords, 0, 0)
	vao.AddAttribData(2, 3, part.Normals, 0, 0)
	gl.BindVertexArray(0)

	part.Vao = vao
}