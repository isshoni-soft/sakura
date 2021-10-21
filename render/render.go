package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/threading"
)

var logger = logging.NewLogger("renderer", 20)

func Init() {
	threading.RunMain(func() {
		logger.Log("Initializing renderer...")

		if err := gl.Init(); err != nil {
			panic(err)
		}
	}, true)

	logger.Log("Initialized OpenGL v", GLVersion())
}

func GLVersion() string {
	return threading.RunMainResult(func () interface {} {
		return gl.GoStr(gl.GetString(gl.VERSION))
	}).(string)
}