package main

import (
	"image"
	"image/draw"
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	cubePositions = []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{-1.5, -2.2, -2.5},
		{-3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1.0, -1.5},
	}
)

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
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.MakeContextCurrent()

	return window
}

func initGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.DEPTH_TEST)
}

func makeObject(verts []float32) uint32 {
	var VBO, VAO uint32

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	makeVBO(VBO, verts)
	// makeEBO(EBO, indices)

	//POS
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, fSize*5, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// //Color
	// gl.VertexAttribPointer(1, 3, gl.FLOAT, false, fSize*5, gl.PtrOffset(12))
	// gl.EnableVertexAttribArray(1)

	// //Text Coords
	// gl.VertexAttribPointer(2, 2, gl.FLOAT, false, fSize*5, gl.PtrOffset(12))
	// gl.EnableVertexAttribArray(2)

	// m := loadTextImg(texturePath)

	// gl.GenTextures(1, &texture)
	// gl.BindTexture(gl.TEXTURE_2D, texture)

	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	// gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, 512, 512, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(m.Pix))
	// gl.GenerateMipmap(gl.TEXTURE_2D)

	return VAO
}
func makeVBO(vbo uint32, verts []float32) {
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, fSize*len(verts), gl.Ptr(verts), gl.STATIC_DRAW)
}

func makeEBO(ebo uint32, indices []uint32) {
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, fSize*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)
}
func loadTextImg(texturePath string) *image.RGBA {
	file, err := os.Open(texturePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	m := image.NewRGBA(data.Bounds())
	draw.Draw(m, m.Bounds(), data, image.Pt(0, 0), draw.Src)
	return m
}

func clear() {
	gl.ClearColor(0.2, 0.3, 0.3, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
