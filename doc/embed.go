package doc

import (
	"embed"
)

//go:embed OpenAPI/*
var OpenAPI embed.FS

//go:embed public/*
var Public embed.FS
