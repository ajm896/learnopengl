package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//Light ...
type Light struct {
	Position  mgl32.Vec3
	Mesh      []float32
	Color     mgl32.Vec3
	Vao       uint32
	Shader    Shader
	Transform mgl32.Mat4
}

//NewLight ...
func NewLight(pos, color mgl32.Vec3, vao uint32, shader Shader) Light {
	return Light{
		Position:  pos,
		Color:     color,
		Transform: mgl32.Ident4().Mul4(mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())),
		Vao:       vao,
		Shader:    shader,
	}
}

//DefaultLight ...
func DefaultLight(vao uint32, shader Shader) Light {
	return NewLight(mgl32.Vec3{1, 1, 1},
		mgl32.Vec3{1, 1, 1}, vao, shader)
}

//DoTransform ...
func (l *Light) DoTransform(m mgl32.Mat4) {
	l.Transform = l.Transform.Mul4(m)
	// c.Transform = m.Mul4(c.Transform)
}

//Draw ...
func (l Light) Draw(view, proj mgl32.Mat4) {
	l.Shader.Use()
	l.Shader.SetMat4("view\x00", view)
	l.Shader.SetMat4("proj\x00", proj)
	l.Shader.SetMat4("model\x00", l.Transform)
	gl.BindVertexArray(l.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}
