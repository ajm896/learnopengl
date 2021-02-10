package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	//CubeMesh ...
	CubeMesh = []float32{
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
)

//Cube Object
type Cube struct {
	Point     mgl32.Vec3
	Transform mgl32.Mat4
	Vao       uint32
	Shader    Shader
}

//NewCube ...
func NewCube(point mgl32.Vec3, vao uint32, shader Shader) Cube {
	return Cube{
		Point: point,
		Transform: mgl32.Ident4().Mul4(mgl32.Translate3D(point.X(),
			point.Y(), point.Z())),
		Vao:    vao,
		Shader: shader,
	}
}

//DoTransform ...
func (c *Cube) DoTransform(m mgl32.Mat4) {
	c.Transform = c.Transform.Mul4(m)
	// c.Transform = m.Mul4(c.Transform)
}

//Draw ...
func (c Cube) Draw(view, proj mgl32.Mat4) {
	c.Shader.Use()
	c.Shader.SetVec3("lightColor\x00", mgl32.Vec3{0.5, 1, 1})
	c.Shader.SetMat4("view\x00", view)
	c.Shader.SetMat4("proj\x00", proj)
	c.Shader.SetMat4("model\x00", c.Transform)
	gl.BindVertexArray(c.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}
