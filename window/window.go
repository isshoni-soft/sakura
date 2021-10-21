package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/threading"
)

var title = ""

var width = 500
var height = 500

var visible = false

var window *glfw.Window

var logger = logging.NewLogger("window", 20)

func Init() {
	threading.RunMain(func() {
		logger.Log("Initializing GLFW...")

		err := glfw.Init()

		if err != nil {
			panic(err)
		}
	}, true)

	logger.Log("Initialized GLFW v", GLFWVersion())
}

func Display() {
	logger.Log("Displaying window...")
	threading.RunMain(func() {
		if window == nil {
			logger.Log("Constructing window...")
			glfw.WindowHint(glfw.Resizable, glfw.False)
			glfw.WindowHint(glfw.ContextVersionMajor, 4)
			glfw.WindowHint(glfw.ContextVersionMinor, 6)
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

			var err error
			window, err = glfw.CreateWindow(Width(), Height(), Title(), nil, nil)

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
	}, true)
}

func SwapBuffers() {
	if window == nil {
		return
	}

	threading.RunMain(func() {
		window.SwapBuffers()
	}, true)
}

func PollEvents() {
	threading.RunMain(func() {
		glfw.PollEvents()
	}, true)
}

func Destroy() {
	threading.RunMain(func() {
		window.Destroy()

		visible = false
	}, true)
}

func Shutdown() {
	logger.Log("Shutting down window...")

	ifVisible(func() { Destroy() })

	threading.RunMain(func() { glfw.Terminate() }, true)
}

func SetTitle(newTitle string) {
	title = newTitle

	ifVisible(func() { threading.RunMain(func() { window.SetTitle(newTitle) }, true) })
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

	return threading.RunMainResult(func() interface {} {
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
	return threading.RunMainResult(func() interface {} {
		return glfw.GetVersionString()
	}).(string)
}

func updateSize() {
	ifVisible(func() { threading.RunMain(func() { window.SetSize(width, height) }, true) })
}

func ifVisible(f func()) {
	if Visible() {
		f()
	}
}
