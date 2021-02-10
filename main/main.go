package main

import (
	_ "image/jpeg"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	fSize = 4
	w     = 800
	h     = 600

	texturePath    = "assets/container.jpg"
	boxVShaderPath = "shaders/box.vert"
	boxFShaderPath = "shaders/box.frag"

	lampVShaderPath = "shaders/lamp.vert"
	lampFShaderPath = "shaders/lamp.frag"
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

	initGL()

	cubeShader := NewShader(boxVShaderPath, boxFShaderPath)
	lampShader := NewShader(lampVShaderPath, lampFShaderPath)

	cube := NewCube(cubePositions[0], CubeMesh, cubeShader)
	light := DefaultLight(CubeMesh, lampShader)

	scale := mgl32.Scale3D(.5, .5, .5)
	light.DoTransform(scale)

	obs := []drawable{cube, light}

	window.SetKeyCallback(keyHandler)
	window.SetCursorPosCallback(mouseControl)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	proj := mgl32.Perspective(mgl32.DegToRad(75), w/h, 0.1, 100)

	for !window.ShouldClose() {
		tick()
		glfw.PollEvents()
		view := camera.GetViewMatrix()
		render(view, proj, obs)
		window.SwapBuffers()
	}
}

func render(view, proj mgl32.Mat4, obs []drawable) {
	clear()
	// obs[0].Draw(view, proj)
	for i := 0; i < len(obs); i++ {
		obs[i].Draw(view, proj)
	}

}

func tick() {
	currentTime := glfw.GetTime()
	dT = float32(currentTime) - lastTime
	lastTime = float32(currentTime)
}
