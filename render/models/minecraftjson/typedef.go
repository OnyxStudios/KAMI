package minecraftjson

type Serialized struct {
	Parent string              `json:"parent"`
	Textures map[string]string `json:"textures"`
	Elements []Element         `json:"elements"`
}

type Element struct {
	Name string           `json:"name"`
	From []float32        `json:"from"`
	To []float32          `json:"to"`
	Faces map[string]Face `json:"faces"`

	Vertices, TextureCoords, Normals []float32
	Indices []uint32
}

type Face struct {
	Texture string `json:"texture"`
	Uv []float32 `json:"uv"`
	Rotation float32 `json:"rotation"`
}
