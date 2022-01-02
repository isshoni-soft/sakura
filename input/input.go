package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/isshoni-soft/sakura/event"
	"github.com/isshoni-soft/sakura/event/events"
)

var held = make(map[glfw.Key]bool)

type KeyEventData struct {
	Key       glfw.Key
	KeyName   string
	Scancode  int
	Action    glfw.Action
	Modifiers glfw.ModifierKey
}

func IsKeyDown(key glfw.Key) bool {
	if v, ok := held[key]; ok {
		return v
	}

	return false
}

func GLFWCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		held[key] = true
	} else if action == glfw.Release {
		held[key] = false
	}

	event.FireEvent(&event.Event{
		Name: events.INPUT,
		Data: KeyEventData{
			Key:       key,
			KeyName:   glfw.GetKeyName(key, glfw.GetKeyScancode(key)),
			Scancode:  scancode,
			Action:    action,
			Modifiers: mods,
		},
	})
}
