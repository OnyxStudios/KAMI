package models

import (
	"github.com/go-gl/gl/all-core/gl"
	"kami/render"
)

type Model struct {
	Name string
	Vao render.VertexArrayObject
	Vertices, TextureCoords, Normals []float32
	Indices []uint32
}

func (model *Model) GenerateModelVAO() {
	vao := render.VertexArrayObject{BufferCount:3}
	render.LoadVAO(&vao)
	vao.Bind()
	vao.AddAttribData(0, 3, model.Vertices, 0, 0)
	vao.AddAttribData(1, 2, model.TextureCoords, 0, 0)
	vao.AddAttribData(2, 3, model.Normals, 0, 0)
	gl.BindVertexArray(0)

	model.Vao = vao
}