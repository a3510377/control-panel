package node

import (
	"encoding/json"
)

type EventName string

const (
	StateEvent   EventName = "STATE"
	MessageEvent EventName = "MESSAGE"
)

type HandleEvent struct {
	Name EventName `json:"name"`
	Data any       `json:"data"` // for `MessageEvent` is `string`, for `StateEvent` is `int`
}

func (ev *HandleEvent) JsonString() (string, error) {
	buf, err := json.Marshal(ev)
	return string(buf), err
}

func (i *Instance) AddEvent(handler func(HandleEvent)) {
	i.handles = append(i.handles, handler)
}

func (i *Instance) Dispatch(ev HandleEvent) {
	for _, handler := range i.handles {
		handler(ev)
	}
}
