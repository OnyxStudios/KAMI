package obj

import (
	"bytes"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"io"
	"kami/render/models"
	"kami/util"
)

func LoadModel(file string) models.Model {
	objData := util.ReadAsset(file)
	objReader := bytes.NewReader(objData)
	modelPart := models.ModelPart{}

	var x, y, z float32
	var textureCoords []mgl32.Vec2
	var normals []mgl32.Vec3

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
			fmt.Fscanf(objReader, "%f %f %f\n", &x, &y, &z)
			modelPart.Vertices = append(modelPart.Vertices, x, y, z)

		// NORMALS.
		case "vn":
			fmt.Fscanf(objReader, "%f %f %f\n", &x, &y, &z)
			normals = append(normals, mgl32.Vec3{x, y, z})

		// TEXTURE VERTICES.
		case "vt":
			fmt.Fscanf(objReader, "%f %f\n", &x, &y)
			textureCoords = append(textureCoords, mgl32.Vec2{x, y})

		// INDICES.
		case "f":
			norm := make([]float32, 4)
			indices := make([]float32, 4)
			uv := make([]float32, 4)
			matches, _ := fmt.Fscanf(objReader, "%f/%f/%f %f/%f/%f %f/%f/%f %f/%f/%f\n", &indices[0], &uv[0], &norm[0], &indices[1], &uv[1], &norm[1], &indices[2], &uv[2], &norm[2], &indices[3], &uv[3], &norm[3])

			if (matches != 9 && matches != 12) || textureCoords == nil {
				panic("Cannot read OBJ file")
			}

			modelPart.Indices = append(modelPart.Indices, uint32(indices[0]-1))
			modelPart.Indices = append(modelPart.Indices, uint32(indices[1]-1))
			modelPart.Indices = append(modelPart.Indices, uint32(indices[2]-1))

			modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[0]) -1].X(), 1 - textureCoords[int(uv[0]) -1].Y())
			modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[1]) -1].X(), 1 - textureCoords[int(uv[1]) -1].Y())
			modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[2]) -1].X(), 1 - textureCoords[int(uv[2]) -1].Y())

			modelPart.Normals = append(modelPart.Normals, normals[int(norm[0]) - 1].X(), normals[int(norm[0]) - 1].Y(), normals[int(norm[0]) - 1].Z())
			modelPart.Normals = append(modelPart.Normals, normals[int(norm[1]) - 1].X(), normals[int(norm[1]) - 1].Y(), normals[int(norm[1]) - 1].Z())
			modelPart.Normals = append(modelPart.Normals, normals[int(norm[2]) - 1].X(), normals[int(norm[2]) - 1].Y(), normals[int(norm[2]) - 1].Z())

			//Triangulate if face is a Quad
			if matches == 12 {
				modelPart.Indices = append(modelPart.Indices, uint32(indices[0]-1), uint32(indices[2]-1), uint32(indices[3]-1))

				modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[0]) -1].X(), 1 - textureCoords[int(uv[0]) -1].Y())
				modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[2]) -1].X(), 1 - textureCoords[int(uv[2]) -1].Y())
				modelPart.TextureCoords = append(modelPart.TextureCoords, textureCoords[int(uv[3]) -1].X(), 1 - textureCoords[int(uv[3]) -1].Y())

				modelPart.Normals = append(modelPart.Normals, normals[int(norm[0]) - 1].X(), normals[int(norm[0]) - 1].Y(), normals[int(norm[0]) - 1].Z())
				modelPart.Normals = append(modelPart.Normals, normals[int(norm[2]) - 1].X(), normals[int(norm[2]) - 1].Y(), normals[int(norm[2]) - 1].Z())
				modelPart.Normals = append(modelPart.Normals, normals[int(norm[3]) - 1].X(), normals[int(norm[3]) - 1].Y(), normals[int(norm[3]) - 1].Z())
			}
		}
	}

	modelPart.GenerateModelVAO()
	return models.Model{Parts:[]models.ModelPart{modelPart}}
}