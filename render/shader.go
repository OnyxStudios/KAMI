package render

import (
	"github.com/go-gl/gl/all-core/gl"
	"strings"
)

var DefaultShaderProgram = DefaultShader{
	Shader:          ShaderProgram{Location:"shaders/viewport"},
}

func LoadShaders() {
	LoadProgram(&DefaultShaderProgram.Shader)
	DefaultShaderProgram.Shader.SetAttribLocation(0, "position")
	DefaultShaderProgram.Shader.SetAttribLocation(1, "textureCoords")
	DefaultShaderProgram.Shader.SetAttribLocation(2, "normal")
	DefaultShaderProgram.TransformMatrix = DefaultShaderProgram.Shader.CreateUniformLocation("transformationMatrix")
	DefaultShaderProgram.ViewMatrix = DefaultShaderProgram.Shader.CreateUniformLocation("viewMatrix")
}

type DefaultShader struct {
	Shader ShaderProgram
	TransformMatrix int32
	ViewMatrix int32
}

//a shader program that connects a vertex and fragment shader
type ShaderProgram struct {
	Handle   uint32
	Location string
	uniforms []string
	attributes []string
}

func (shader *ShaderProgram) UseShader() {
	gl.UseProgram(shader.Handle)
}

func (shader *ShaderProgram) SetAttribLocation(index uint32, location string) {
	if !strings.HasSuffix(location, "\x00") {
		location += "\x00"
	}
	shader.attributes = append(shader.attributes, location)
	gl.BindAttribLocation(shader.Handle, index, gl.Str(location))
	CheckGlError()
}

func (shader *ShaderProgram) CreateUniformLocation(uniform string) int32 {
	if !strings.HasSuffix(uniform, "\x00") {
		uniform += "\x00"
	}
	shader.uniforms = append(shader.uniforms, uniform)
	location := gl.GetUniformLocation(shader.Handle, gl.Str(uniform))
	CheckGlError()
	return location
}