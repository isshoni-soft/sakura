package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/threading"
	"unsafe"
)

var logger = logging.NewLogger("renderer", 20)

type Renderer interface {
	Draw()
	Clear()
}

type Renderable interface {
	Render()
}

func Init() {
	threading.RunMain(func() {
		logger.Log("Initializing renderer...")

		if err := gl.Init(); err != nil {
			panic(err)
		}
	}, true)

	logger.Log("Initialized OpenGL v", GLVersion())
}

func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	threading.RunMain(func() {
		gl.VertexAttribPointer(index, size, xtype, normalized, stride, pointer)
	}, true)
}

func EnableVertexAttribArray(index uint32) {
	threading.RunMain(func() {
		gl.EnableVertexAttribArray(index)
	}, true)
}

func BindVertexArray(arrays uint32) {
	threading.RunMain(func() {
		gl.BindVertexArray(arrays)
	}, true)
}

func GenVertexArrays(n int32, arrays *uint32) {
	threading.RunMain(func() {
		gl.GenVertexArrays(n, arrays)
	}, true)
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	threading.RunMain(func() {
		gl.BufferData(target, size, data, usage)
	}, true)
}

func BindBuffer(target uint32, buffer uint32) {
	threading.RunMain(func() {
		gl.BindBuffer(target, buffer)
	}, true)
}

func GenBuffers(n int32, buffers *uint32) {
	threading.RunMain(func() {
		gl.GenBuffers(n, buffers)
	}, true)
}

func DepthFunc(cap uint32) {
	threading.RunMain(func() {
		gl.DepthFunc(cap)
	}, true)
}

func Enable(cap uint32) {
	threading.RunMain(func() {
		gl.Enable(cap)
	}, true)
}

func ClearColor(red float32, green float32, blue float32, alpha float32) {
	threading.RunMain(func() {
		gl.ClearColor(red, green, blue, alpha)
	}, true)
}

func Clear(mask uint32) {
	threading.RunMain(func() {
		gl.Clear(mask)
	}, true)
}

func QueueRender(renderable Renderable) {
	threading.RunMain(renderable.Render)
}

func Render(renderable Renderable) {
	threading.RunMain(renderable.Render, true)
}

func GLVersion() string {
	return threading.RunMainResult(func () interface {} {
		return gl.GoStr(gl.GetString(gl.VERSION))
	}).(string)
}
