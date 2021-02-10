package main

import (
	_ "image/jpeg"
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
	dT       float32 = 0.0
	lastTime float32 = 0.0

	lastX float32 = 400
	lastY float32 = 300

	camera = DefaultCamera()

	firstMouse = true
)

func main() {
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	prog := initGL()

	vao, texture := makeObject(mesh)

	window.SetKeyCallback(keyHandler)
	window.SetCursorPosCallback(mouseControl)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	proj := mgl32.Perspective(mgl32.DegToRad(75), w/h, 0.1, 100)

	for !window.ShouldClose() {
		tick()
		prog.Use()

		view := camera.GetViewMatrix()

		prog.SetMat4("view\x00", view)
		prog.SetMat4("proj\x00", proj)

		glfw.PollEvents()
		render(vao, texture, prog)
		window.SwapBuffers()
	}
}

func render(vao, texture uint32, prog Shader) {
	clear()

	gl.BindVertexArray(vao)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	for i := 0; i < 10; i++ {
		cube := NewCube(cubePositions[i], vao)

		trans := mgl32.Translate3D(cube.Point.X(), cube.Point.Y(), cube.Point.Z())
		rot := mgl32.HomogRotate3D(mgl32.DegToRad(20.0*float32(i)), mgl32.Vec3{1.0, .3, .5})

		cube.DoTransform(trans)
		cube.DoTransform(rot)

		prog.SetMat4("model\x00", cube.Transform)

		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}

}

func tick() {
	currentTime := glfw.GetTime()
	dT = float32(currentTime) - lastTime
	lastTime = float32(currentTime)
}
