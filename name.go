package buttonsapi

const (
	ValueEnvLocal = "local"
	ValueEnvAlpha = "alpha"
	ValueEnvProd  = "prod"
)

type Service string

const (
	ButtonsAPI   Service = "buttons-api"
	ButtonsAdmin Service = "buttons-admin"
	ButtonsWeb   Service = "buttons-web"
)

const (
	Google   = "google"
	Facebook = "facebook"
)
