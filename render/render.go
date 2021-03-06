package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/kirito"
	"github.com/isshoni-soft/roxxy/v1"
	"unsafe"
)

var logger = roxxy_v1.NewLogger("sakura|renderer>")

type Renderer interface {
	Draw()
	Clear()
}

type Renderable interface {
	Render()
}

func Init() {
	kirito.QueueBlocking(func() {
		logger.Log("Initializing renderer...")

		if err := gl.Init(); err != nil {
			panic(err)
		}
	})

	logger.Log("Initialized OpenGL v", GLVersion())
}

func DrawArrays(mode uint32, first int32, second int32) {
	kirito.QueueBlocking(func() {
		gl.DrawArrays(mode, first, second)
	})
}

func UseShaderProgram(program *ShaderProgram) {
	kirito.QueueBlocking(func() {
		gl.UseProgram(program.Id())
	})
}

func LinkShaderProgram(program *ShaderProgram) {
	kirito.QueueBlocking(func() {
		gl.AttachShader(program.Id(), program.FragmentShader().Id())
		gl.AttachShader(program.Id(), program.VertexShader().Id())
		gl.LinkProgram(program.Id())
	})
}

func CompileShader(shader *Shader) {
	kirito.QueueBlocking(func() {
		strs, free := gl.Strs(shader.code...)

		gl.ShaderSource(shader.Id(), 1, strs, nil)
		gl.CompileShader(shader.Id())

		free()

		if !isShaderCompiled(shader) {
			logger.Log("Shader failed to compile!")
			logger.Log(getShaderLogs(shader))
			panic("Shader failed to compile")
		}
	})
}

func getShaderLogs(shader *Shader) string {
	var maxLength int32

	gl.GetShaderiv(shader.Id(), gl.INFO_LOG_LENGTH, &maxLength)

	var errorLog uint8

	gl.GetShaderInfoLog(shader.Id(), maxLength, &maxLength, &errorLog)

	return gl.GoStr(&errorLog)
}

func GetShaderLogs(shader *Shader) string {
	return kirito.Get(func() string { return getShaderLogs(shader) })
}

func isShaderCompiled(shader *Shader) bool {
	var result int32

	gl.GetShaderiv(shader.Id(), gl.COMPILE_STATUS, &result)

	return result == gl.TRUE
}

func IsShaderCompiled(shader *Shader) bool {
	return kirito.Get(func() bool { return isShaderCompiled(shader) })
}

func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	kirito.QueueBlocking(func() {
		gl.VertexAttribPointer(index, size, xtype, normalized, stride, pointer)
	})
}

func EnableVertexAttribArray(index uint32) {
	kirito.QueueBlocking(func() {
		gl.EnableVertexAttribArray(index)
	})
}

func BindVertexArray(arrays uint32) {
	kirito.QueueBlocking(func() {
		gl.BindVertexArray(arrays)
	})
}

func GenVertexArrays(n int32, arrays *uint32) {
	kirito.QueueBlocking(func() {
		gl.GenVertexArrays(n, arrays)
	})
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	kirito.QueueBlocking(func() {
		gl.BufferData(target, size, data, usage)
	})
}

func BindBuffer(target uint32, buffer uint32) {
	kirito.QueueBlocking(func() {
		gl.BindBuffer(target, buffer)
	})
}

func GenBuffers(n int32, buffers *uint32) {
	kirito.QueueBlocking(func() {
		gl.GenBuffers(n, buffers)
	})
}

func DepthFunc(cap uint32) {
	kirito.QueueBlocking(func() {
		gl.DepthFunc(cap)
	})
}

func Enable(cap uint32) {
	kirito.QueueBlocking(func() {
		gl.Enable(cap)
	})
}

func ClearColor(red float32, green float32, blue float32, alpha float32) {
	kirito.QueueBlocking(func() {
		gl.ClearColor(red, green, blue, alpha)
	})
}

func Clear(mask uint32) {
	kirito.QueueBlocking(func() {
		gl.Clear(mask)
	})
}

func QueueRender(renderable Renderable) {
	kirito.Queue(renderable.Render)
}

func Render(renderable Renderable) {
	kirito.QueueBlocking(renderable.Render)
}

func GLVersion() string {
	return kirito.Get(func() string {
		return gl.GoStr(gl.GetString(gl.VERSION))
	})
}
