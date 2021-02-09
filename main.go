package main

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	fSize = 4
	w     = 800
	h     = 600

	texturePath = "assets/container.jpg"
	vPath       = "shaders/box.vert"
	fPath       = "shaders/box.frag"
)

var (
	bModel = []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	bIndices = []uint32{
		0, 1, 3,
		1, 2, 3,
	}
)

func main() {
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	prog := initGL()

	vao, texture := makeObject(bModel, bIndices)

	window.SetKeyCallback(keyHandler)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	for !window.ShouldClose() {
		prog.Use()
		model := mgl32.HomogRotate3DX(mgl32.DegToRad(-55))
		view := mgl32.Translate3D(0, 0, -3)
		proj := mgl32.Perspective(mgl32.DegToRad(45), w/h, 0.1, 100)
		prog.SetMat4("model\x00", model)
		prog.SetMat4("view\x00", view)
		prog.SetMat4("proj\x00", proj)

		glfw.PollEvents()
		render(vao, texture)
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

func initGL() Shader {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := NewShader(vPath, fPath)
	gl.Enable(gl.DEPTH_TEST)

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

func render(vao uint32, texture uint32) {
	clear()

	gl.BindVertexArray(vao)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINES)
	// gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}

func makeObject(verts []float32, indices []uint32) (uint32, uint32) {
	var VBO, VAO, EBO, texture uint32

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	makeVBO(VBO, verts)
	makeEBO(EBO, indices)

	//POS
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, fSize*5, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// //Color
	// gl.VertexAttribPointer(1, 3, gl.FLOAT, false, fSize*5, gl.PtrOffset(12))
	// gl.EnableVertexAttribArray(1)

	//Text Coords
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, fSize*5, gl.PtrOffset(12))
	gl.EnableVertexAttribArray(2)

	m := loadTextImg(texturePath)

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, 512, 512, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(m.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return VAO, texture
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
