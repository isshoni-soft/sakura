package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/isshoni-soft/sakura/event"
	"github.com/isshoni-soft/sakura/event/events"
)

type KeyEventData struct {
	Key glfw.Key
}

func GLFWCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	event.FireEvent(&event.Event{
		Name: events.INPUT,
		Data: KeyEventData{
			Key: key,
		},
	})
}
