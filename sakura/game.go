package sakura

import (
	"github.com/isshoni-soft/kirito"
	"github.com/isshoni-soft/roxxy"
	"github.com/isshoni-soft/sakura/render"
	"github.com/isshoni-soft/sakura/window"
	"math"
	"time"
)

var gameSingleton *Game

var debug = false

var logger = roxxy.NewLogger("sakura>")

var shutdownSignal = make(chan bool)

var version = Version{
	Major:    0,
	Minor:    0,
	Patch:    7,
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

type Game struct {
	Tickable
	render.Renderer
	Initializable

	GameData *interface{}
	Version  Version
	Logger   roxxy.Logger

	renderDelta *DeltaTicker
	logicDelta  *DeltaTicker
}

func GetLogger() *roxxy.Logger {
	return logger
}

func GetGame() *Game {
	return gameSingleton
}

func SetDebug(b bool) {
	debug = b
}

func Debug() bool {
	return debug
}

func Init(game *Game) {
	gameSingleton = game
	gameSingleton.renderDelta = &DeltaTicker{
		Function: renderTick,
	}
	gameSingleton.logicDelta = &DeltaTicker{
		Function: logicTick,
	}

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

		// go mainTicker(game)

		logger.Log("Finishing initialization!")

		game.PostInit()

		logger.Log("Finished initialization!")
		logger.Log("Awaiting shutdown signal...")
		<-shutdownSignal
	})
}

func renderTick(game Game, ticker DeltaTicker) {
	defer Shutdown()

	for !window.ShouldClose() {
		game.Clear()
		game.Draw()
		window.SwapBuffers()
		window.PollEvents()
	}
}

func logicTick(game Game, ticker DeltaTicker) {
	rate := 20
	done := 0
	goal := 1 * rate
	start := time.Now()

	for !window.ShouldClose() {
		target := math.Floor(float64((time.Now().UnixMilli() - start.UnixMilli()) * int64(rate)))
		game.Tick()
	}
}

func Shutdown() {
	logger.Log("Shutting down Sakura...")

	window.Shutdown()
	shutdownSignal <- true
}
