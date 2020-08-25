package kami

import (
	"github.com/go-gl/gl/all-core/gl"
	"kami/render"
)

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