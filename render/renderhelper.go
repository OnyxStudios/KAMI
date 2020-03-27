package render

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"kami/util"
	"strings"
)

const logSize = 512

var isGlInit = false

func LoadVBO(vbo *VertexBufferObject) {
	InitGL()
	gl.GenBuffers(1, &vbo.Handle)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.Handle)
	gl.BufferData(gl.ARRAY_BUFFER, len(vbo.Vertices)*4, gl.Ptr(vbo.Vertices), gl.STATIC_DRAW)

	//cleanup
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func InitGL() {
	if !isGlInit {
		util.CheckErr(gl.Init())
		fmt.Printf("OpenGL Version %v\n", gl.GoStr(gl.GetString(gl.VERSION)))
		isGlInit = true
	}
}

func LoadVAO(vao *VertexArrayObject) {
	InitGL()
	if vao.VBO.Handle == 0 {
		LoadVBO(&vao.VBO)
	}
	gl.GenVertexArrays(1, &vao.Handle)
	gl.BindVertexArray(vao.Handle)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vao.VBO.Handle)
	//tell opengl how to interpret the vbo data
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	//cleanup
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func LoadProgram(program *ShaderProgram) {
	InitGL()
	hVSH := LoadShader(fmt.Sprintf("%v.vsh", program.Location), gl.VERTEX_SHADER)
	hFSH := LoadShader(fmt.Sprintf("%v.fsh", program.Location), gl.FRAGMENT_SHADER)
	program.Handle = gl.CreateProgram()
	gl.AttachShader(program.Handle, hVSH)
	gl.AttachShader(program.Handle, hFSH)
	gl.LinkProgram(program.Handle)
	var success int32
	gl.GetProgramiv(program.Handle, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		log := gl.Str(strings.Repeat("\x00", logSize))
		gl.GetProgramInfoLog(program.Handle, logSize, nil, log)
		panic(fmt.Errorf("shader program link error: %v", gl.GoStr(log)))
	}
	gl.UseProgram(program.Handle)

	//cleanup
	gl.DeleteShader(hVSH)
	gl.DeleteShader(hFSH)
	gl.UseProgram(0)
}

func LoadShader(location string, xtype uint32) uint32 {
	InitGL()
	handle := gl.CreateShader(xtype)
	shaderSrc, freeFn := gl.Strs(util.SReadAsset(location) + "\x00")
	defer freeFn()
	gl.ShaderSource(handle, 1, shaderSrc, nil)
	gl.CompileShader(handle)
	var success int32
	gl.GetShaderiv(handle, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		log := gl.Str(strings.Repeat("\x00", logSize))
		gl.GetShaderInfoLog(handle, logSize, nil, log)
		panic(fmt.Errorf("shader compile error: %v", gl.GoStr(log)))
	}
	return handle
}
