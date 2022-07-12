package hub

import (
	"cafe/service"
	"github.com/sirupsen/logrus"
)

type (
	NotifierChan chan service.Order

	Hub struct {
		Notifier       NotifierChan
		NewClients     chan NotifierChan
		ClosingClients chan NotifierChan
		clients        map[NotifierChan]struct{}
	}
)

func (h *Hub) Initialize() {
	h.Notifier = make(NotifierChan)
	h.NewClients = make(chan NotifierChan)
	h.ClosingClients = make(chan NotifierChan)
	h.clients = make(map[NotifierChan]struct{})
	go h.listen()
	go func() {
		for {
			if msg, ok := <-service.LiveCh; ok {
				h.Notifier <- msg
			}
		}
	}()
}

func (h *Hub) listen() {
	for {
		select {
		case s := <-h.NewClients:
			h.clients[s] = struct{}{}
			logrus.WithField("openConnections", len(h.clients)).Trace("Websocket connection added")
		case s := <-h.ClosingClients:
			delete(h.clients, s)
			logrus.WithField("openConnections", len(h.clients)).Trace("Websocket connection removed")
		case event := <-h.Notifier:
			for client := range h.clients {
				select {
				case client <- event:
				default:
					close(client)
					delete(h.clients, client)
				}
			}
		}
	}
}
