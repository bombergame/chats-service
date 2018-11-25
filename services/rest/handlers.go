package rest

import (
	"log"
	"net/http"
)

func (srv *Service) createChat(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) getChat(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) updateChat(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) deleteChat(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) joinChat(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
