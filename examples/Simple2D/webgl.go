package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

func main() {
	// Step1: Prepare the canvas and get WebGL context

	document := js.Global.Get("document")
	canvas := document.Call("getElementById", "my_Canvas")

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	// Step2: Define the geometry and store it in buffer objects

	vertices := [6]float32{-0.5, 0.5, -0.5, -0.5, 0.0, -0.5}

	// Create a new buffer object
	vertex_buffer := gl.CreateBuffer()

	// Bind an empty array buffer to it
	gl.BindBuffer(gl.ARRAY_BUFFER, vertex_buffer)

	// Pass the vertices data to the buffer
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	// Unbind the buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, nil)

	/* Step3: Create and compile Shader programs */

	// Vertex shader source code
	vertCode := `attribute vec2 coordinates;` +
		`void main(void) {` + ` gl_Position = vec4(coordinates,0.0, 1.0);` + `}`

	//Create a vertex shader object
	vertShader := gl.CreateShader(gl.VERTEX_SHADER)

	//Attach vertex shader source code
	gl.ShaderSource(vertShader, vertCode)

	//Compile the vertex shader
	gl.CompileShader(vertShader)

	//Fragment shader source code
	fragCode := `void main(void) {` + `gl_FragColor = vec4(0.0, 0.0, 0.0, 0.1);` + `}`

	// Create fragment shader object
	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	// Attach fragment shader source code
	gl.ShaderSource(fragShader, fragCode)

	// Compile the fragment shader
	gl.CompileShader(fragShader)

	// Create a shader program object to store combined shader program
	shaderProgram := gl.CreateProgram()

	// Attach a vertex shader
	gl.AttachShader(shaderProgram, vertShader)

	// Attach a fragment shader
	gl.AttachShader(shaderProgram, fragShader)

	// Link both programs
	gl.LinkProgram(shaderProgram)

	// Use the combined shader program object
	gl.UseProgram(shaderProgram)

	/* Step 4: Associate the shader programs to buffer objects */

	//Bind vertex buffer object
	gl.BindBuffer(gl.ARRAY_BUFFER, vertex_buffer)

	//Get the attribute location
	coord := gl.GetAttribLocation(shaderProgram, "coordinates")

	//point an attribute to the currently bound VBO
	gl.VertexAttribPointer(coord, 2, gl.FLOAT, false, 0, 0)

	//Enable the attribute
	gl.EnableVertexAttribArray(coord)

	/* Step5: Drawing the required object (triangle) */

	// Clear the canvas
	gl.ClearColor(0.5, 0.5, 0.5, 0.9)

	// Enable the depth test
	gl.Enable(gl.DEPTH_TEST)

	// Clear the color buffer bit
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set the view port
	gl.Viewport(0, 0, 300, 300)

	// Draw the triangle
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
