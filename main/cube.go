package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	//CubeMesh ...
	CubeMesh = []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0,
		0.5, -0.5, -0.5, 0.0, 0.0, -1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0,

		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0,

		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, -1.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0,
		-0.5, -0.5, 0.5, -1.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0,
		0.5, -0.5, -0.5, 0.0, -1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0,

		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
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
func NewCube(point mgl32.Vec3, mesh []float32, shader Shader) Cube {
	return Cube{
		Point: point,
		Transform: mgl32.Ident4().Mul4(mgl32.Translate3D(point.X(),
			point.Y(), point.Z())),
		Vao:    makeCubeVao(mesh),
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
	c.Shader.SetVec3("lightPos\x00", LightPos)
	c.Shader.SetVec3("lightColor\x00", mgl32.Vec3{1, 1, 1})

	//Material
	c.Shader.SetVec3("material.ambient\x00", mgl32.Vec3{0.0215, 0.1745, 0.0215})
	c.Shader.SetVec3("material.diffuse\x00", mgl32.Vec3{0.07568, 0.61424, 0.07568})
	c.Shader.SetVec3("material.specular\x00", mgl32.Vec3{0.633, 0.727811, 0.633})
	c.Shader.SetFloat("material.shininess\x00", .6*128)

	c.Shader.SetVec3("light.ambinet\x00", mgl32.Vec3{0.2, 0.2, 0.2})
	c.Shader.SetVec3("light.diffuse\x00", mgl32.Vec3{0.5, 0.5, 0.5})
	c.Shader.SetVec3("light.specular\x00", mgl32.Vec3{1.0, 1.0, 1.0})

	c.Shader.SetMat4("view\x00", view)
	c.Shader.SetMat4("proj\x00", proj)
	c.Shader.SetMat4("model\x00", c.Transform)
	c.Shader.SetVec3("viewPos\x00", camera.Position)
	gl.BindVertexArray(c.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}

func makeCubeVao(verts []float32) uint32 {
	var VBO, VAO uint32

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	makeVBO(VBO, verts)
	// makeEBO(EBO, indices)

	//POS
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, fSize*6, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	//Normals
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, fSize*6, gl.PtrOffset(12))
	gl.EnableVertexAttribArray(1)

	return VAO
}
