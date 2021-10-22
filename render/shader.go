package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/sakura/threading"
)

type Shader struct {
	id uint32
	code []string
	compiled bool
}

func (s *Shader) Id() uint32 {
	return s.id
}

func (s *Shader) Code() []string {
	return s.code
}

func (s *Shader) Compiled() bool {
	return s.compiled
}

func ShaderFromStrings(xtype uint32, code ...string) *Shader {
	id := threading.RunMainResult(func () interface {} {
		return gl.CreateShader(xtype)
	}).(uint32)

	return &Shader {
		id: id,
		code: code,
		compiled: false,
	}
}

type ShaderProgram struct {
	id uint32
	vertex *Shader
	fragment *Shader
}

func (sp *ShaderProgram) VertexShader() *Shader {
	return sp.vertex
}

func (sp *ShaderProgram) FragmentShader() *Shader {
	return sp.fragment
}

func (sp *ShaderProgram) Id() uint32 {
	return sp.id
}

func NewShaderProgram(vertex *Shader, fragment *Shader) *ShaderProgram {
	id := threading.RunMainResult(func() interface {} {
		return gl.CreateProgram()
	}).(uint32)

	return &ShaderProgram {
		id: id,
		vertex: vertex,
		fragment: fragment,
	}
}
