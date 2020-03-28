package render

import (
	"github.com/go-gl/gl/all-core/gl"
)

//simple VBO to keep our data
type VertexBufferObject struct {
	Handle   uint32
	Vertices []float32
}

// VAO stores references to positions inside a VBO
type VertexArrayObject struct {
	Handle        uint32
	BufferCount   int32
	VertexBuffers []VertexBufferObject
}

func (vao *VertexArrayObject) Bind() {
	gl.BindVertexArray(vao.Handle)
}

func (vao *VertexArrayObject) AddAttribData(index uint32, attribSize int32, data []float32, stride int32, offset int) {
	vbo := VertexBufferObject{Vertices: data}
	LoadVBO(&vbo)
	vao.VertexBuffers[index] = vbo
	vao.Bind()
	gl.EnableVertexAttribArray(index)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.Handle)
	//tell opengl how to interpret the vbo data
	//gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.VertexAttribPointer(index, attribSize, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	//cleanup
	CheckGlError()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}
