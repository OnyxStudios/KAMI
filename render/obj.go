package render

import (
	"bytes"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"io"
	"kami/util"
)

type Model struct {
	Normals, Vertices []mgl32.Vec3
	TextureCoords []mgl32.Vec2
	VecIndices, NormalIndices, UvIndices []float32
}

func NewModel(file string) Model {
	objData := util.ReadAsset(file)
	objReader := bytes.NewReader(objData)
	model := Model{}

	for {
		var lineType string
		_, err := fmt.Fscanf(objReader, "%s", &lineType)

		if err != nil {
			if err == io.EOF {
				break
			}
		}

		switch lineType {
		// VERTICES.
		case "v":
			vec := mgl32.Vec3{}
			fmt.Fscanf(objReader, "%f %f %f\n", &vec[0], &vec[1], &vec[2])
			model.Vertices = append(model.Vertices, vec)

		// NORMALS.
		case "vn":
			vec := mgl32.Vec3{}
			fmt.Fscanf(objReader, "%f %f %f\n", &vec[0], &vec[1], &vec[2])
			model.Normals = append(model.Normals, vec)

		// TEXTURE VERTICES.
		case "vt":
			vec := mgl32.Vec2{}
			fmt.Fscanf(objReader, "%f %f\n", &vec[0], &vec[1])
			model.TextureCoords = append(model.TextureCoords, vec)

		// INDICES.
		case "f":
			norm := make([]float32, 3)
			vec := make([]float32, 3)
			uv := make([]float32, 3)
			matches, _ := fmt.Fscanf(objReader, "%f/%f/%f %f/%f/%f %f/%f/%f\n", &vec[0], &uv[0], &norm[0], &vec[1], &uv[1], &norm[1], &vec[2], &uv[2], &norm[2])

			if matches != 9 {
				panic("Cannot read OBJ file")
			}

			model.NormalIndices = append(model.NormalIndices, norm[0])
			model.NormalIndices = append(model.NormalIndices, norm[1])
			model.NormalIndices = append(model.NormalIndices, norm[2])

			model.VecIndices = append(model.VecIndices, vec[0] - 1)
			model.VecIndices = append(model.VecIndices, vec[1] - 1)
			model.VecIndices = append(model.VecIndices, vec[2] - 1)

			model.UvIndices = append(model.UvIndices, uv[0])
			model.UvIndices = append(model.UvIndices, uv[1])
			model.UvIndices = append(model.UvIndices, uv[2])
		}
	}

	return model
}

func (model Model) GetVertices() []float32 {
	var vertices []float32

	for _, value := range model.Vertices {
		vertices = append(vertices, value.X())
		vertices = append(vertices, value.Y())
		vertices = append(vertices, value.Z())
	}

	return vertices
}

func (model Model) GetTextureCoords() []float32 {
	var textureCoords []float32

	for _, value := range model.TextureCoords {
		textureCoords = append(textureCoords, value.X())
		textureCoords = append(textureCoords, value.Y())
	}

	return textureCoords
}

func (model Model) GetNormals() []float32 {
	var normals []float32

	for _, value := range model.Normals {
		normals = append(normals, value.X())
		normals = append(normals, value.Y())
		normals = append(normals, value.Z())
	}

	return normals
}

func (model Model) GetIndicesAsIntArr() []uint32 {
	var indices []uint32

	for _, value := range model.VecIndices {
		indices = append(indices, uint32(value))
	}

	return indices
}