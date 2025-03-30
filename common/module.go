package common

import (
	"go.uber.org/fx"

	"auth/common/cache"
	"auth/common/config"
)

var Module = fx.Module("common",
	config.Module,
	cache.Module,
	//store.Module,
)
