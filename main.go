package main

import (
	"github.com/dlyycious/gopos/config"
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/internal/server"
)

func main() {
	database.Run()
	server.Run(config.APP_ADDRESS)
}
