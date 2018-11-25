package rest

import (
	"fmt"
	"github.com/bombergame/chats-service/config"
	"github.com/bombergame/chats-service/repositories"
	"github.com/bombergame/chats-service/utils"
	"github.com/bombergame/common/rest"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

const (
	BasePath = "/chats"
)

type Service struct {
	rest.Service
	config     Config
	components Components
	upgrader   websocket.Upgrader
}

type Config struct {
	rest.Config
}

type Components struct {
	rest.Components
	ConnManager    *utils.ConnectionManager
	ChatRepository repositories.ChatRepository
}

func NewService(cf Config, cpn Components) *Service {
	cf.Host, cf.Port = "", config.HttpPort

	srv := &Service{
		Service: *rest.NewService(
			cf.Config,
			cpn.Components,
		),
		config:     cf,
		components: cpn,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	mx := mux.NewRouter()
	mx.Handle(
		fmt.Sprintf("%s", BasePath),
		handlers.MethodHandler{
			http.MethodPost: http.HandlerFunc(srv.createChat),
		},
	)
	mx.Handle(
		fmt.Sprintf("%s/%s", BasePath, ChatIDPath),
		handlers.MethodHandler{
			http.MethodGet:    http.HandlerFunc(srv.getChat),
			http.MethodPatch:  http.HandlerFunc(srv.updateChat),
			http.MethodDelete: http.HandlerFunc(srv.deleteChat),
		},
	)
	mx.Handle(
		fmt.Sprintf("%s/%s/websocket", BasePath, ChatIDPath),
		http.HandlerFunc(srv.joinChat),
	)

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
