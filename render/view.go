package render

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func CreateTransformMatrix(translation mgl32.Vec3, rotation mgl32.Quat, scale float32) mgl32.Mat4 {
	matrix := mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.Translate3D(translation.X(), translation.Y(), translation.Z()))
	//matrix = matrix.Mul4(rotation.Mat4()) //FIXME why doesn't this work
	matrix = matrix.Mul4(mgl32.HomogRotate3DX(rotation.X()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(rotation.Y()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DZ(rotation.Z()))
	matrix = matrix.Mul4(mgl32.Scale3D(scale, -scale, scale)) //scale by negative y to flip everything, because screen coords start top left

	return matrix
}

func CreateViewMatrix(position mgl32.Vec3, rotation mgl32.Quat) mgl32.Mat4 {
	matrix := mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.HomogRotate3DX(rotation.X()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(rotation.Y()))
	matrix = matrix.Mul4(mgl32.Translate3D(-position.X(), -position.Y(), -position.Z()))

	return matrix
}

func LoadProjectionMatrix(shader *ShaderProgram, camera Camera) {
	shader.UseShader()
	matrixID := shader.CreateUniformLocation("projectionMatrix")
	gl.UniformMatrix4fv(matrixID, 1, false, &camera.Projection[0])
	CheckGlError()
	//cleanup
	gl.UseProgram(0)
}