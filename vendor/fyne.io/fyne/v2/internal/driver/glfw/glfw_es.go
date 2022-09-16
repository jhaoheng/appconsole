//go:build (gles || arm || arm64) && !darwin && !js && !wasm && !test_web_driver
// +build gles arm arm64
// +build !darwin
// +build !js
// +build !wasm
// +build !test_web_driver

package glfw

import "github.com/go-gl/glfw/v3.3/glfw"

func initWindowHints() {
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
}
