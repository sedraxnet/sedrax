package prefixmanager

import (
	"github.com/sedraxnet/sedrax/infrastructure/logger"
	"github.com/sedraxnet/sedrax/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
