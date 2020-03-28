package render

import "github.com/go-gl/mathgl/mgl64"

var (
	MainCamera = Camera{Position: mgl64.Vec3{}, Rotation: mgl64.QuatIdent()}
)

type Camera struct {
	Position mgl64.Vec3
	Rotation mgl64.Quat
	Projection mgl64.Mat4
}