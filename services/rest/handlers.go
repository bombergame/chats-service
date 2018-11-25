package rest

import (
	"encoding/json"
	"github.com/satori/go.uuid"
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
	var authID int64 = 1

	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//err = srv.components.connManager.AddConnection(authID, conn)
	//if err != nil {
	//	log.Println(err)
	//}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}

		var cmdRequest CommandRequest
		if err := json.Unmarshal(p, &cmdRequest); err != nil {
			log.Println(err)
			continue
		}

		log.Println(cmdRequest)

		if cmdRequest.Type != "make_private_chat" {
			log.Println(" not make private chat")
			continue
		}

		cmdResponse := CommandResponse{
			Type: "make_private_chat",
			Data: MakePrivateChatResponseData{
				ChatID: uuid.NewV4(),
			},
		}

		rsp, _ := json.Marshal(cmdResponse)

		if err := conn.WriteMessage(messageType, rsp); err != nil {
			log.Println(err)
			return
		}
	}
}
