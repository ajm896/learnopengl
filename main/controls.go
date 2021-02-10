package main

import "github.com/go-gl/glfw/v3.3/glfw"

func mouseControl(w *glfw.Window, xpos, ypos float64) {
	if firstMouse {
		lastX = float32(xpos)
		lastY = float32(ypos)
		firstMouse = false
	}
	xOffset := float32(xpos) - lastX
	yOffset := lastY - float32(ypos)

	lastX = float32(xpos)
	lastY = float32(ypos)

	camera.ProcessMouseMovement(xOffset, yOffset)

}

func keyHandler(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {

	switch key {
	case glfw.KeyEscape:
		w.SetShouldClose(true)
	case glfw.KeyW:
		camera.ProcessKeyboard(Up, dT)
	case glfw.KeyA:
		camera.ProcessKeyboard(Left, dT)
	case glfw.KeyS:
		camera.ProcessKeyboard(Down, dT)
	case glfw.KeyD:
		camera.ProcessKeyboard(Right, dT)
	default:
		return
	}
}
