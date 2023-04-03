package wendyevents

import (
	"encoding/json"

	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
	"github.com/Meduzz/wendy"
)

type (
	Event struct {
		Topic string
		Body  *wendy.Body
	}

	EventListener struct {
		name      string
		listeners map[string]Listener
	}

	Listener func(*Event)
)

func NewEventListener(name string) *EventListener {
	listeners := make(map[string]Listener)

	return &EventListener{name, listeners}
}

func (e *EventListener) WithListener(topic string, listener Listener) {
	e.listeners[topic] = listener
}

func (e *EventListener) Start(srv *rpc.RPC) error {
	for topic, listener := range e.listeners {
		err := srv.Handler(topic, e.name, e.wrapper(listener))

		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EventListener) wrapper(listener Listener) api.Handler {
	return func(ctx api.Context) {
		event := &Event{}
		event.Topic = ctx.Msg().Subject
		ct := ctx.Msg().Header.Get("Content-Type")

		switch ct {
		case "application/json":
			event.Body = wendy.Json(json.RawMessage(ctx.Raw()))
		case "text/html":
		case "text/javascript":
		case "text/css":
			event.Body = wendy.Static(ct, ctx.Raw())
		default:
			event.Body = wendy.Text(string(ctx.Msg().Data))
		}

		listener(event)
	}
}
