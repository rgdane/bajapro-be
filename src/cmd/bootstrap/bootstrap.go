package bootstrap

import (
	"jk-api/internal/container"
)

var (
	Services *container.AppContainer
)

func Setup() {
	InitEnv()
	Services = InitContainer()
}
