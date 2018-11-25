package main

import (
	"github.com/bombergame/chats-service/repositories/postgres"
	"github.com/bombergame/chats-service/services/rest"
	"github.com/bombergame/chats-service/utils"
	"github.com/bombergame/common/logs"
	restful "github.com/bombergame/common/rest"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	conn := postgres.NewConnection()
	defer conn.Close()
	if err := conn.Open(); err != nil {
		logger.Fatal(err)
		return
	}

	restSrv := rest.NewService(
		rest.Config{},
		rest.Components{
			Components: restful.Components{
				Logger: logger,
			},
			ChatRepository: postgres.NewProfileRepository(conn),
			ConnManager:    utils.NewConnectionManager(),
		},
	)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := restSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch

	if err := restSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
