package models

import (
	"bytes"
	"fmt"
	"io"
	"kami/util"
)

func NewModel(file string) Model {
	objData := util.ReadAsset(file)
	objReader := bytes.NewReader(objData)
	model := Model{}

	var x, y, z float32
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
			model.Vertices = append(model.Vertices, x, y, z)

		// NORMALS.
		case "vn":
			fmt.Fscanf(objReader, "%f %f %f\n", &x, &y, &z)
			model.Normals = append(model.Normals, x, y, z)

		// TEXTURE VERTICES.
		case "vt":
			fmt.Fscanf(objReader, "%f %f\n", &x, &y)
			model.TextureCoords = append(model.TextureCoords, x, y)

		// INDICES.
		case "f":
			norm := make([]float32, 4)
			vec := make([]uint32, 4)
			uv := make([]float32, 4)
			matches, _ := fmt.Fscanf(objReader, "%f/%f/%f %f/%f/%f %f/%f/%f %f/%f/%f\n", &vec[0], &uv[0], &norm[0], &vec[1], &uv[1], &norm[1], &vec[2], &uv[2], &norm[2], &vec[3], &uv[3], &norm[3])

			if matches != 9 && matches != 12 {
				panic("Cannot read OBJ file")
			}

			model.Indices = append(model.Indices, vec[0] - 1, vec[1] - 1, vec[2] - 1)

			//Triangulate if face is a Quad
			if matches == 12 {
				model.Indices = append(model.Indices, vec[0] - 1, vec[2] - 1, vec[3] - 1)
			}
		}
	}

	model.GenerateModelVAO()
	return model
}