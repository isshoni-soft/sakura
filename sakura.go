package sakura

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/kirito"
	"github.com/isshoni-soft/roxxy"
	"github.com/isshoni-soft/sakura/render"
	"github.com/isshoni-soft/sakura/window"
)

var logger = roxxy.NewLogger("sakura")

var version = Version {
	Major: 0,
	Minor: 0,
	Patch: 2,
	Snapshot: debug,
}

type Initializable interface {
	PreInit()
	Init()
	PostInit()

	SetInitialized(initialized bool)
	Initialized() bool
}

type Tickable interface {
	Tick()
}

type Game interface {
	Tickable
	render.Renderer
	Initializable
}

type SimpleGame struct { }

func (sg *SimpleGame) Clear() {
	render.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

var debug = false

func GetLogger() *roxxy.Logger {
	return logger
}

func SetDebug(b bool) {
	debug = b
}

func Debug() bool {
	return debug
}

func Init(game Game) {
	logger.Log("Initializing the Sakura Engine v", version.GetVersion())

	kirito.Run(func() {
		game.PreInit()
		window.Init()
		window.Display()
		render.Init()

		if Debug() {
			window.SetTitle(window.Title() + " (GL v" + render.GLVersion() + ")")
		}

		game.Init()

		logger.Log("launching main ticker...")

		go mainTicker(game)

		logger.Log("Finishing initialization!")

		game.PostInit()

		logger.Log("Finished initialization!")
	})
}

func mainTicker(game Game) {
	defer Shutdown()

	for !window.ShouldClose() {
		game.Tick()
		game.Clear()
		game.Draw()
		window.SwapBuffers()
		window.PollEvents()
	}
}

func Shutdown() {
	logger.Log("Shutting down Sakura...")

	window.Shutdown()
}
