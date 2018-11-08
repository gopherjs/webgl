# webgl

[![GoDoc](https://godoc.org/github.com/gopherjs/webgl?status.svg)](https://godoc.org/github.com/gopherjs/webgl)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopherjs/webgl)](https://goreportcard.com/report/github.com/gopherjs/webgl)

[GopherJS](https://github.com/gopherjs/gopherjs) bindings for [WebGL 1.0](https://www.khronos.org/registry/webgl/specs/latest/1.0/) context.

## Example

![Screenshot](https://cloud.githubusercontent.com/assets/1924134/3566022/5d81f2d0-0ae0-11e4-82e4-3cb33b83d8d3.png)

webgl_example.go:

```Go
package main

import (
	"fmt"

	"github.com/gopherjs/gopherwasm/js"
	"github.com/gopherjs/webgl"
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("createElement", "canvas")
	if canvas == js.Null() {
		fmt.Println("CANVAS IS NULL")
		return
	}

	canvas.Set("width", 800)
	canvas.Set("height", 600)
	canvas.Get("style").Set("outline", "none")
	document.Get("body").Call("appendChild", canvas)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global().Call("alert", "Error: "+err.Error())
	}

	gl.Call("clearColor", 0.8, 0.3, 0.01, 1)
	gl.Call("clear", gl.COLOR_BUFFER_BIT)
}
```

js_example.html:

```html
<html><body><script src="webgl_example.js"></script></body></html>
```

To produce `webgl_example.js` file, run `gopherjs build webgl_example.go`.

wasm_example:
```html
<html>
	<head>
		<meta charset="utf-8">
		<script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
				go.run(result.instance);
			});
		</script>
	</head>
	<body></body>
</html>
```

1. `GOOS=js GOARCH=wasm go build -o main.wasm`
2. `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"`
3. Serve the files. E.g. with `goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'`

