package features

import (
	"github.com/gin-contrib/pprof"

	"github.com/simiancreative/services/server"
)

func AddPprof() {
	router := server.GetRouter()
	pprof.Register(router)
}
