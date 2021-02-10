package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//LightPos ...
var LightPos mgl32.Vec3 = mgl32.Vec3{3, 3, 3}

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
func NewLight(pos, color mgl32.Vec3, mesh []float32, shader Shader) Light {
	return Light{
		Position:  pos,
		Color:     color,
		Transform: mgl32.Ident4().Mul4(mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())),
		Vao:       makeLampVao(mesh),
		Shader:    shader,
	}
}

//DefaultLight ...
func DefaultLight(mesh []float32, shader Shader) Light {
	return NewLight(LightPos,
		mgl32.Vec3{1, 1, 1}, mesh, shader)
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

func makeLampVao(verts []float32) uint32 {
	var VAO uint32

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	//makeVBO(VBO, verts)

	//POS
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, fSize*6, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	return VAO
}
