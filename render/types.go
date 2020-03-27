package render

//a shader program that connects a vertex and fragment shader
type ShaderProgram struct {
	Handle   uint32
	Location string
}

//simple VBO to keep our data
type VertexBufferObject struct {
	Handle   uint32
	Vertices []float32
}

// VAO stores references to positions inside a VBO
type VertexArrayObject struct {
	Handle uint32
	VBO    VertexBufferObject
}

func MakeVAO(vertices []float32) VertexArrayObject {
	vbo := VertexBufferObject{Vertices: vertices}
	LoadVBO(&vbo)
	vao := VertexArrayObject{VBO: vbo}
	LoadVAO(&vao)
	return vao
}
