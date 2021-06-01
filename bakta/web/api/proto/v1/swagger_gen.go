package api

import (
	"embed"
)

//go:embed jobs.swagger.json
var f embed.FS

func GetSwaggerEmbedded() *embed.FS {
	return &f
}
