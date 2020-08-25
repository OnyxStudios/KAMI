package constants

import (
	"encoding/json"
	"kami/util"
)

//see assets/app.json
var Config = AppConfig{}

type AppConfig struct {
	Title string `json:"title"`
	Version string `json:"version"`
	WindowWidth int `json:"window_width"`
	WindowHeight int `json:"window_height"`
}

func init() {
	util.CheckErr(json.Unmarshal(util.ReadAsset("app.json"), &Config))
}
