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
		util.Log.Println(fmt.Sprintf("OpenGL Version %v", gl.GoStr(gl.GetString(gl.VERSION))))
		isGlInit = true
	}
}

func LoadVAO(vao *VertexArrayObject) {
	InitGL()
	if vao.BufferCount <= 0 {
		util.ErrLog.Panicln("tried to create empty VAO!")
	}
	vao.VertexBuffers = make([]VertexBufferObject, vao.BufferCount)
	gl.GenVertexArrays(vao.BufferCount, &vao.Handle)
}

func LoadProgram(program *ShaderProgram) {
	InitGL()
	hVSH := LoadShader(fmt.Sprintf("%v.vsh", program.Location), gl.VERTEX_SHADER)
	hGSH := LoadShader(fmt.Sprintf("%v.gsh", program.Location), gl.GEOMETRY_SHADER)
	hFSH := LoadShader(fmt.Sprintf("%v.fsh", program.Location), gl.FRAGMENT_SHADER)
	program.Handle = gl.CreateProgram()
	gl.AttachShader(program.Handle, hVSH)
	if hGSH != 0 {
		gl.AttachShader(program.Handle, hGSH)
	}
	gl.AttachShader(program.Handle, hFSH)
	gl.LinkProgram(program.Handle)
	var success int32
	gl.GetProgramiv(program.Handle, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		log := gl.Str(strings.Repeat("\x00", logSize))
		gl.GetProgramInfoLog(program.Handle, logSize, nil, log)
		util.ErrLog.Panicln(fmt.Errorf("shader program link error: %v", gl.GoStr(log)))
	}
	program.UseShader()

	//cleanup
	gl.DeleteShader(hVSH)
	gl.DeleteShader(hFSH)
	gl.UseProgram(0)
}

func LoadShader(location string, xtype uint32) uint32 {
	InitGL()
	if !util.AssetExists(location) {
		return 0
	}
	shaderSrc, freeFn := gl.Strs(util.SReadAsset(location) + "\x00")
	defer freeFn()
	handle := gl.CreateShader(xtype)
	gl.ShaderSource(handle, 1, shaderSrc, nil)
	gl.CompileShader(handle)
	var success int32
	gl.GetShaderiv(handle, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		log := gl.Str(strings.Repeat("\x00", logSize))
		gl.GetShaderInfoLog(handle, logSize, nil, log)
		util.ErrLog.Panicln(fmt.Errorf("shader compile error: %v", gl.GoStr(log)))
	}
	return handle
}

func CheckGlError() {
	for err := gl.GetError(); err != gl.NO_ERROR; err = gl.GetError() {
		util.ErrLog.Println(fmt.Sprintf("OpenGL ERROR: %v", err))
	}
}
