package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

//Cube Object
type Cube struct {
	Point     mgl32.Vec3
	Transform mgl32.Mat4
	Vao       uint32
}

//NewCube ...
func NewCube(point mgl32.Vec3, vao uint32) Cube {
	return Cube{
		Point:     point,
		Transform: mgl32.Ident4(),
		Vao:       vao,
	}
}

//DoTransform ...
func (c *Cube) DoTransform(m mgl32.Mat4) {
	c.Transform = c.Transform.Mul4(m)
	// c.Transform = m.Mul4(c.Transform)
}
