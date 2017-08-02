package core

import "github.com/merakiVE/CVDI/core/config"

type ContextController struct {
	Config config.Configuration
}

type Controller interface {
	SetContext(ContextController)
}
