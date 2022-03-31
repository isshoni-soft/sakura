package sakura

import (
	"github.com/isshoni-soft/kirito"
	"github.com/isshoni-soft/roxxy/v1"
	"github.com/isshoni-soft/sakura/event"
	"github.com/isshoni-soft/sakura/event/events"
	"github.com/isshoni-soft/sakura/render"
	"github.com/isshoni-soft/sakura/window"
)

var game *Game

var debug = false
var running = true

var logger = roxxy_v1.NewLogger("sakura>")

var shutdownSignal = make(chan bool)

var version = Version{
	Major:    0,
	Minor:    0,
	Patch:    8,
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

type GameLogic interface {
	Tickable
	render.Renderer
	Initializable
}

type Game struct {
	Logic   GameLogic
	Version Version
	Logger  *roxxy_v1.Logger

	renderDelta *DeltaTicker
	logicDelta  *DeltaTicker
}

func (g *Game) RenderDelta() *DeltaTicker {
	return g.renderDelta
}

func (g *Game) LogicDelta() *DeltaTicker {
	return g.logicDelta
}

func GetLogger() *roxxy_v1.Logger {
	return logger
}

func GetGame() *Game {
	return game
}

func SetDebug(b bool) {
	debug = b
}

func Debug() bool {
	return debug
}

func Running() bool {
	return !window.ShouldClose() && running
}

func Init(g Game) {
	game = &g

	game.renderDelta = &DeltaTicker{
		Function: renderTick,
		Defer:    Shutdown,
	}
	game.logicDelta = &DeltaTicker{
		Function: logicTick,
		Defer:    func() {},
	}

	logger.Log("Initializing the Sakura Engine v", version.GetVersion())

	kirito.Run(func() {
		game.Logic.PreInit()
		window.Init()
		window.Display()
		render.Init()

		if Debug() {
			window.SetTitle(window.Title() + " (GL v" + render.GLVersion() + ")")
		}

		game.Logic.Init()

		logger.Log("launching main ticker...")

		game.logicDelta.Run(game)
		game.renderDelta.Run(game)

		logger.Log("Finishing initialization!")

		game.Logic.PostInit()

		logger.Log("Finished initialization!")
		logger.Log("Awaiting shutdown signal...")
		<-shutdownSignal
	})
}

func renderTick(game *Game, ticker *DeltaTicker) {
	ticker.RecalculateDelta()
	event.FireEvent(&event.Event{
		Name: events.ENGINE_RENDER_TICK,
		Data: nil,
	})
	game.Logic.Clear()
	game.Logic.Draw()
	window.SwapBuffers()
	window.PollEvents()
}

func logicTick(game *Game, ticker *DeltaTicker) {
	ticker.RecalculateDelta()
	event.FireEvent(&event.Event{
		Name: events.ENGINE_LOGIC_TICK,
		Data: nil,
	})
	game.Logic.Tick()
}

func Shutdown() {
	running = false
	logger.Log("Shutting down Sakura...")

	window.Shutdown()
	shutdownSignal <- true
}
