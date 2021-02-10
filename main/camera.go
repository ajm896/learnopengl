package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	//YAW ...
	YAW float32 = -90
	//PITCH ...
	PITCH float32 = 0.0
	//SPEED ...
	SPEED float32 = 10
	//SENSITIVITY ...
	SENSITIVITY float32 = 0.1
)

//CamMove ...
type CamMove int

const (
	//Up Fore Cam
	Up CamMove = iota
	//Down back cam
	Down
	//Left left
	Left
	//Right right
	Right
)

//Camera ...
type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WorldUp  mgl32.Vec3

	Yaw   float32
	Pitch float32

	Speed       float32
	Sensitivity float32
}

//DefaultCamera ...
func DefaultCamera() Camera {
	c := Camera{
		Position:    mgl32.Vec3{-1, 1, 3.0},
		Front:       mgl32.Vec3{0, 0, -1.0},
		Up:          mgl32.Vec3{0, 1.0, 0},
		WorldUp:     mgl32.Vec3{0, 1.0, 0},
		Yaw:         YAW,
		Pitch:       PITCH,
		Speed:       SPEED,
		Sensitivity: SENSITIVITY,
	}

	c.updateCameraVectors()
	return c
}

//GetViewMatrix ...
func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	view := mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
	return view
}

//ProcessKeyboard ...
func (c *Camera) ProcessKeyboard(direction CamMove, deltaTime float32) {
	vel := c.Speed * deltaTime

	switch direction {
	case Up:
		c.Position = c.Position.Add(c.Front.Mul(vel))
	case Down:
		c.Position = c.Position.Sub(c.Front.Mul(vel))
	case Left:
		c.Position = c.Position.Sub(c.Right.Mul(vel))
	case Right:
		c.Position = c.Position.Add(c.Right.Mul(vel))
	}

	c.updateCameraVectors()
}

//ProcessMouseMovement ...
func (c *Camera) ProcessMouseMovement(xOffset, yOffset float32) {
	xOffset *= c.Sensitivity
	yOffset *= c.Sensitivity

	c.Yaw += xOffset
	c.Pitch += yOffset

	if c.Pitch > 89.0 {
		c.Pitch = 89.0
	}
	if c.Pitch < -89.0 {
		c.Pitch = -89.0
	}

	c.updateCameraVectors()
}

func (c *Camera) updateCameraVectors() {
	// fmt.Println("Start: ", c.Front)
	newX := math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))
	newY := math.Sin(float64(mgl32.DegToRad(c.Pitch)))
	newZ := math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))
	// fmt.Println(c.Front)
	c.Front[0] = float32(newX)
	c.Front[1] = float32(newY)
	c.Front[2] = float32(newZ)
	// fmt.Println(c.Front)
	c.Front = c.Front.Normalize()

	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
	// fmt.Println("End: ", c.Front)

	// c.Position = c.Front.Mul(-1)
}
