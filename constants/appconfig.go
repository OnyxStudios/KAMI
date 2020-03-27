package constants

import (
	"encoding/json"
	"kami/util"
)

type AppConfig struct {
}

var Config = AppConfig{}

func init() {
	util.CheckErr(json.Unmarshal(util.ReadAsset("app.json"), &Config))
}
