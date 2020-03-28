package render

import (
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