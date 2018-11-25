package rest

import (
	"encoding/json"
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

	log.Println("upgraded")

	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	log.Println("auth received")

	var authID AuthID
	if err := json.Unmarshal(p, &authID); err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	log.Println("auth parsed")

	err = srv.components.ConnManager.AddConnection(authID.AuthID, conn)
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	log.Println("connection added to manager")

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

		switch cmdRequest.Type {
		case "make_private_chat":
			data, ok := cmdRequest.Data.(MakePrivateChatRequestData)
			if !ok {
				log.Println(cmdRequest)
				continue
			}

			id, err := srv.components.ChatRepository.StartPrivateChat(authID.AuthID, data.ProfileID)
			if err != nil {
				log.Println(err)
				continue
			}

			cmdResponse := CommandResponse{
				Type: "make_private_chat",
				Data: MakePrivateChatResponseData{
					ChatID: id,
				},
			}

			rsp, _ := json.Marshal(cmdResponse)

			if err := conn.WriteMessage(messageType, rsp); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
