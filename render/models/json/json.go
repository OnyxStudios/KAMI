package json

import (
	"encoding/json"
	"github.com/go-gl/mathgl/mgl32"
	"kami/render/models"
	"kami/util"
)

type JsonFormat struct {
	Parent string `json:"parent"`
	Textures map[string]string `json:"textures"`
	Elements []Element `json:"elements"`
}

type Element struct {
	Name string `json:"name"`
	From []float32 `json:"from"`
	To []float32 `json:"to"`
	Faces map[string]Face `json:"faces"`

	Vertices, TextureCoords, Normals []float32
	Indices []uint32
}

type Face struct {
	Texture string `json:"texture"`
	Uv []float32 `json:"uv"`
}

func LoadModel(path string) models.Model {
	modelString, _ := util.CheckReadFile(path)

	if len(modelString) <= 0 {
		modelString = util.SReadAsset(path)
	}

	var jsonFormat JsonFormat
	json.Unmarshal([]byte(modelString), &jsonFormat)

	GenerateModelData(&jsonFormat)
	model := models.Model{}

	for _, element := range jsonFormat.Elements {
		part := models.ModelPart{
			Name:          element.Name,
			Vertices:      element.Vertices,
			TextureCoords: element.TextureCoords,
			Normals:       element.Normals,
			Indices:       element.Indices,
		}

		part.GenerateModelVAO()
		model.Parts = append(model.Parts, part)
	}

	return model
}

func GenerateModelData(model *JsonFormat) {
	var vertices []float32
	var textureCoords []float32
	var normals []float32
	var indices []uint32

	for index, element := range model.Elements  {
		startVertex := mgl32.Vec3{element.From[0], element.From[1], element.From[2]}
		endVertex := mgl32.Vec3{element.To[0], element.To[1], element.To[2]}

		GenerateFace(mgl32.Vec3{startVertex.X(), startVertex.Y(), startVertex.Z()}, mgl32.Vec3{startVertex.X(), endVertex.Y(), endVertex.Z()}, &vertices, &normals, &indices)
		GenerateFace(mgl32.Vec3{endVertex.X(), startVertex.Y(), startVertex.Z()}, mgl32.Vec3{endVertex.X(), endVertex.Y(), endVertex.Z()}, &vertices, &normals, &indices)

		GenerateFace(mgl32.Vec3{startVertex.X(), startVertex.Y(), startVertex.Z()}, mgl32.Vec3{endVertex.X(), endVertex.Y(), startVertex.Z()}, &vertices, &normals, &indices)
		GenerateFace(mgl32.Vec3{startVertex.X(), startVertex.Y(), endVertex.Z()}, mgl32.Vec3{endVertex.X(), endVertex.Y(), endVertex.Z()}, &vertices, &normals, &indices)

		GenerateFace(mgl32.Vec3{startVertex.X(), startVertex.Y(), startVertex.Z()}, mgl32.Vec3{endVertex.X(), startVertex.Y(), endVertex.Z()}, &vertices, &normals, &indices)
		GenerateFace(mgl32.Vec3{startVertex.X(), endVertex.Y(), startVertex.Z()}, mgl32.Vec3{endVertex.X(), endVertex.Y(), endVertex.Z()}, &vertices, &normals, &indices)

		for _, face := range element.Faces {
			textureCoords = append(textureCoords, face.Uv...)
		}

		model.Elements[index].Vertices = vertices
		model.Elements[index].TextureCoords = textureCoords
		model.Elements[index].Normals = normals
		model.Elements[index].Indices = indices
		vertices = []float32{}
		textureCoords = []float32{}
		normals = []float32{}
		indices = []uint32{}
	}
}

func GenerateFace(startPoint, endPoint mgl32.Vec3, vertices, normals *[]float32, indices *[]uint32) {
	var faceVerts []float32
	var faceNormals []float32
	var faceIndices []uint32

	//Triangle 1
	faceVerts = append(faceVerts, startPoint.X())
	faceVerts = append(faceVerts, startPoint.Y())
	faceVerts = append(faceVerts, startPoint.Z())

	faceVerts = append(faceVerts, endPoint.X())
	faceVerts = append(faceVerts, endPoint.Y())
	faceVerts = append(faceVerts, endPoint.Z())

	if startPoint.Y() != endPoint.Y() {
		faceVerts = append(faceVerts, startPoint.X())
		faceVerts = append(faceVerts, endPoint.Y())
		faceVerts = append(faceVerts, startPoint.Z())
	} else {
		faceVerts = append(faceVerts, endPoint.X())
		faceVerts = append(faceVerts, startPoint.Y())
		faceVerts = append(faceVerts, startPoint.Z())
	}

	//Triangle 2
	faceVerts = append(faceVerts, startPoint.X())
	faceVerts = append(faceVerts, startPoint.Y())
	faceVerts = append(faceVerts, startPoint.Z())

	if startPoint.Y() != endPoint.Y() {
		faceVerts = append(faceVerts, endPoint.X())
		faceVerts = append(faceVerts, endPoint.Y())
		faceVerts = append(faceVerts, endPoint.Z())
	}else {
		faceVerts = append(faceVerts, startPoint.X())
		faceVerts = append(faceVerts, startPoint.Y())
		faceVerts = append(faceVerts, endPoint.Z())
	}

	faceVerts = append(faceVerts, endPoint.X())
	faceVerts = append(faceVerts, startPoint.Y())
	faceVerts = append(faceVerts, endPoint.Z())

	normal := mgl32.Vec3{faceVerts[6] - faceVerts[0], faceVerts[7] - faceVerts[1], faceVerts[8] - faceVerts[2]}.
		Cross(mgl32.Vec3{faceVerts[3] - faceVerts[0], faceVerts[4] - faceVerts[1], faceVerts[5] - faceVerts[2]})

	normal2 := mgl32.Vec3{faceVerts[15] - faceVerts[9], faceVerts[16] - faceVerts[10], faceVerts[17] - faceVerts[11]}.
		Cross(mgl32.Vec3{faceVerts[12] - faceVerts[9], faceVerts[13] - faceVerts[10], faceVerts[14] - faceVerts[11]})

	faceNormals = append(faceNormals, normal.X())
	faceNormals = append(faceNormals, normal.Y())
	faceNormals = append(faceNormals, normal.Z())

	faceNormals = append(faceNormals, normal2.X())
	faceNormals = append(faceNormals, normal2.Y())
	faceNormals = append(faceNormals, normal2.Z())

	*vertices = append(*vertices, faceVerts...)
	*normals = append(*normals, faceNormals...)

	if len(*vertices) >= 18 {
		for i := 18; i > 0; i-- {
			faceIndices = append(faceIndices, uint32(len(*vertices)-i))
		}
	}

	*indices = append(*indices, faceIndices...)
}