package server

import (
	"github.com/dlyycious/gopos/api"
)

func Run(addr string) {
	app := fiberConfig()
	api.Register(app)
	panic(app.Listen(addr))
}
