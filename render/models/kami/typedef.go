package kami

import (
	"github.com/go-gl/mathgl/mgl32"
	"kami/render"
)

type Model struct {
	Parts []ModelPart
}

type ModelPart struct {
	Name string
	Position mgl32.Vec3
	Rotation mgl32.Quat

	Vao render.VertexArrayObject
	Vertices, TextureCoords, Normals []float32
	Indices []uint32
}
