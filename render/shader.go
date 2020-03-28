package render

import (
	"github.com/go-gl/gl/all-core/gl"
	"strings"
)

var DefaultShaderProgram = ShaderProgram{Location:"shaders/viewport"}

func LoadShaders() {
	LoadProgram(&DefaultShaderProgram)
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

func (shader *ShaderProgram) BindAttribLocation(index uint32, location string) {
	if !strings.HasSuffix(location, "\x00") {
		location += "\x00"
	}
	shader.attributes = append(shader.attributes, location)
	gl.BindAttribLocation(shader.Handle, index, gl.Str(location))
}

func (shader *ShaderProgram) CreateUniformLocation(uniform string) int32 {
	if !strings.HasSuffix(uniform, "\x00") {
		uniform += "\x00"
	}
	shader.uniforms = append(shader.uniforms, uniform)
	return gl.GetUniformLocation(shader.Handle, gl.Str(uniform))
}