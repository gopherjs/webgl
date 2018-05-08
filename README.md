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
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

func main() {
	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	document.Get("body").Call("appendChild", canvas)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	gl.ClearColor(0.8, 0.3, 0.01, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
```

webgl_example.html:

```html
<html><body><script src="webgl_example.js"></script></body></html>
```

To produce `webgl_example.js` file, run `gopherjs build webgl_example.go`.
