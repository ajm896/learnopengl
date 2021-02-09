package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	w = 800
	h = 600

	vertexShaderSource = `
	#version 330 core
	layout (location = 0) in vec3 vp;
	layout (location = 1) in vec3 aColor;
	out vec4 color;
	void main() {
		gl_Position = vec4(vp, 1.0);
		color = vec4(aColor, 1.0);
	}
` + "\x00"

	fragmentShaderSource = `
	#version 330 core
	in vec4 color;
	out vec4 frag_colour;
	void main() {
		frag_colour = color;
	}
` + "\x00"
)

var (
	triangle = []float32{
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}

	indices = []uint32{
		0, 1, 3,
		1, 2, 3,
	}
)

func main() {
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	prog := initGL()
	vao := makeVAO(triangle)
	window.SetKeyCallback(keyHandler)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	for !window.ShouldClose() {
		render(prog, vao)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func initGLFW() *glfw.Window {

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w, h, "Hello World", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func keyHandler(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyEscape:
		w.SetShouldClose(true)
	default:
		return
	}
}

func render(p uint32, vao uint32) {
	gl.ClearColor(0.2, 0.3, 0.3, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(p)
	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(indices))
	// gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func makeVAO(verts []float32) uint32 {
	var VBO, VAO uint32
	var EBO uint32
	//gl.BindVertexArray(VAO)

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STATIC_DRAW)

	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1,
		3,
		gl.FLOAT,
		false,
		6*4,
		gl.PtrOffset(12))
	gl.EnableVertexAttribArray(1)
	return VAO
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
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
