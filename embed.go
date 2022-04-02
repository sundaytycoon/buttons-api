package buttonsapi

import "embed"

// StaticSwaggerUI is a collection of pre-built static files for swagger web ui.
//go:embed static/swagger-ui
var StaticSwaggerUI embed.FS

// OAPISpecYAML is the Open API Specifications Manifest document that defines golive HTTP API.
//go:embed api/oapi
var OAPISpecYAML embed.FS

// PublicUI is the Open API Specifications Manifest document that defines golive HTTP API.
//go:embed static/public
var PublicUI embed.FS
