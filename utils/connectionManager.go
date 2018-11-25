package utils

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ConnectionManager struct {
	rwMutex *sync.RWMutex
	conns   map[int64]*websocket.Conn
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		rwMutex: &sync.RWMutex{},
		conns:   make(map[int64]*websocket.Conn, 0),
	}
}

func (m *ConnectionManager) AddConnection(id int64, conn *websocket.Conn) error {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()

	m.conns[id] = conn

	return nil
}

func (m *ConnectionManager) ListenAndServe() {

}
