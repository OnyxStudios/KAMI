package render

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	MainCamera = Camera{Position: mgl32.Vec3{0, 0, 0}, Rotation: mgl32.AnglesToQuat(0, 0 ,0, mgl32.XYZ)}
)

type Camera struct {
	Position mgl32.Vec3
	Rotation mgl32.Quat
	Projection mgl32.Mat4
}

func (camera Camera) UpdateProjectionMatrix(fov, width, height, nearPlane, farPlane float32) {
	camera.Projection =  mgl32.Perspective(mgl32.DegToRad(fov), width/height, nearPlane, farPlane)
	LoadProjectionMatrix(&DefaultShaderProgram, camera)
	gl.Viewport(0, 0, int32(width), int32(height))
}