package render

import "github.com/go-gl/mathgl/mgl64"

func CreateTransformMatrix(translation mgl64.Vec3, rotation mgl64.Quat, scale float64) mgl64.Mat4 {
	matrix := mgl64.Ident4()
	matrix.Mul4(mgl64.Translate3D(translation.X(), translation.Y(), translation.Z()))
	matrix.Mul4(mgl64.HomogRotate3DX(rotation.X()))
	matrix.Mul4(mgl64.HomogRotate3DY(rotation.Y()))
	matrix.Mul4(mgl64.HomogRotate3DZ(rotation.Z()))
	matrix.Mul4(mgl64.Scale3D(scale, scale, scale))

	return matrix
}

func CreateViewMatrix(position mgl64.Vec3, rotation mgl64.Quat) mgl64.Mat4 {
	matrix := mgl64.LookAtV(position.Mul(-1), mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0, 1, 0})
	matrix.Mul4(mgl64.HomogRotate3DX(rotation.X()))
	matrix.Mul4(mgl64.HomogRotate3DY(rotation.Y()))

	return matrix
}