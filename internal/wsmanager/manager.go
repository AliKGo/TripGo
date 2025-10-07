package wsmanager

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	conns map[string]*websocket.Conn // map[userID] = connection
	mu    sync.RWMutex               // protects concurrent access
}

func (m *Manager) Add(userID string, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.conns[userID] = conn
}

func (m *Manager) Remove(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.conns, userID)
}

func (m *Manager) Send(userID string, message []byte) error {
	m.mu.RLock()
	conn, ok := m.conns[userID]
	m.mu.RUnlock()
	if !ok {
		return errors.New("connection not found")
	}
	return conn.WriteMessage(websocket.TextMessage, message)
}
