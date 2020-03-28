package render

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	MainCamera = Camera{Position: mgl32.Vec3{}, Rotation: mgl32.QuatIdent()}
)

type Camera struct {
	Position mgl32.Vec3
	Rotation mgl32.Quat
	Projection mgl32.Mat4
}