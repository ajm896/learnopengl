package main

import "github.com/go-gl/mathgl/mgl32"

type drawable interface {
	Draw(view, proj mgl32.Mat4)
}
