package core

import (
	"github.com/merakiVE/CVDI/core/config"
	"github.com/kataras/iris"
)

type ContextController struct {
	Config config.Configuration
	App    *iris.Application
}

type Controller interface {
	SetContext(ContextController)
}
