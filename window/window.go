package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/isshoni-soft/kirito"
	"github.com/isshoni-soft/roxxy"
	"github.com/isshoni-soft/sakura/input"
)

var title = ""

var width = 500
var height = 500

var visible = false

var window *glfw.Window

var logger = roxxy.NewLogger("sakura|window>")

func Init() {
	kirito.QueueBlocking(func() {
		logger.Log("Initializing GLFW...")

		err := glfw.Init()

		if err != nil {
			panic(err)
		}
	})

	logger.Log("Initialized GLFW v", GLFWVersion())
}

func GetLogger() *roxxy.Logger {
	return logger
}

func Display() {
	logger.Log("Displaying window...")

	kirito.QueueBlocking(func() {
		if window == nil {
			logger.Log("Constructing window...")
			glfw.WindowHint(glfw.Resizable, glfw.False)
			glfw.WindowHint(glfw.ContextVersionMajor, 4)
			glfw.WindowHint(glfw.ContextVersionMinor, 6)
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

			var err error
			window, err = glfw.CreateWindow(Width(), Height(), Title(), nil, nil)
			window.SetKeyCallback(input.GLFWCallback)

			if err != nil {
				panic(err)
			}
			logger.Log("Successfully constructed window")
		} else {
			logger.Log("Reusing previous window")
		}

		window.MakeContextCurrent()
		window.Show()

		visible = true
	})
}

func SwapBuffers() {
	if window == nil {
		return
	}

	kirito.QueueBlocking(func() {
		window.SwapBuffers()
	})
}

func PollEvents() {
	kirito.QueueBlocking(func() {
		glfw.PollEvents()
	})
}

func Destroy() {
	kirito.QueueBlocking(func() {
		window.Destroy()

		visible = false
	})
}

func Shutdown() {
	logger.Log("Shutting down window...")

	ifVisible(func() { Destroy() })

	kirito.QueueBlocking(func() { glfw.Terminate() })
}

func SetTitle(newTitle string) {
	title = newTitle

	ifVisible(func() { kirito.QueueBlocking(func() { window.SetTitle(newTitle) }) })
}

func SetWidth(newWidth int) {
	width = newWidth

	updateSize()
}

func SetHeight(newHeight int) {
	height = newHeight

	updateSize()
}

func ShouldClose() bool {
	if window == nil {
		return false
	}

	return kirito.Get(func() interface{} {
		return window.ShouldClose()
	}).(bool)
}

func Title() string {
	return title
}

func Width() int {
	return width
}

func Height() int {
	return height
}

func Visible() bool {
	return visible
}

func GLFWVersion() string {
	return kirito.Get(func() interface{} {
		return glfw.GetVersionString()
	}).(string)
}

func updateSize() {
	ifVisible(func() { kirito.QueueBlocking(func() { window.SetSize(width, height) }) })
}

func ifVisible(f func()) {
	if Visible() {
		f()
	}
}
