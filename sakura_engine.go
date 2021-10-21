package sakura

import (
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/render"
	"github.com/isshoni-soft/sakura/threading"
	"github.com/isshoni-soft/sakura/window"
)

var version = Version {
	Major: 0,
	Minor: 0,
	Patch: 1,
	Snapshot: debug,
}

type Game interface {
	PreInit()
	Init()
	//PostInit()
}

var debug = false

func SetDebug(b bool) {
	debug = b
}

func Debug() bool {
	return debug
}

func Init(game Game) {
	logging.GetLogger().Log("Initializing the Sakura Engine v", version.GetVersion())

	threading.InitMainThread(10, func() {
		game.PreInit()
		window.Init()
		window.Display()
		render.Init()

		if Debug() {
			window.SetTitle(window.Title() + " (GL v" + render.GLVersion() + ")")
		}

		game.Init()
	})
}

func Shutdown() {
	logging.GetLogger().Log("Shutting down Sakura...")

	threading.ShutdownMainThread()
	window.Shutdown()
}