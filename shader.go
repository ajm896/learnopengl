package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id uint32
}

// NewShader creates new prog with vShader and fShader
func NewShader(vPath string, fPath string) Shader {

	vShader, err := compileShader(vPath, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fShader, err := compileShader(fPath, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vShader)
	gl.AttachShader(prog, fShader)
	gl.LinkProgram(prog)

	gl.DeleteShader(vShader)
	gl.DeleteShader(fShader)

	return Shader{id: prog}
}

// Use Shader
func (s Shader) Use() {
	gl.UseProgram(s.id)
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	source, err := os.Open(path)
	if err != nil {
		panic("Couldn't read frag shader")
	}

	data := make([]byte, 512)
	source.Read(data)
	sourceString := string(data)

	csources, free := gl.Strs(sourceString)
	gl.ShaderSource(shader, 1, csources, nil)

	free()

	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
