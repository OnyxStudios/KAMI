package render

var DefaultShaderProgram = ShaderProgram{Location:"shaders/viewport"}

func LoadShaders() {
	LoadProgram(&DefaultShaderProgram)
}