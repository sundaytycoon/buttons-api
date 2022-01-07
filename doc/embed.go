package doc

import (
	"embed"
)

//go:embed OpenAPI/*
var OpenAPI embed.FS
