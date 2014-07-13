# webgl

[GopherJS](https://github.com/neelance/gopherjs) binding to the webgl context.

```Go
import "github.com/gopherjs/webgl"
import "github.com/gopherjs/gopherjs/js"

document := js.Global("document")
canvas := document.Call("createElement", "canvas")
document.Get("body").Call("appendChild", canvas)

attrs := webgl.DefaultAttributes()
attrs.Alpha = false

gl, err := webgl.NewContext(canvas, attrs)
if err != nil {
	panic(err)
}

gl.ClearColor(1, 0, 0, 1)
gl.Clear(webgl.COLOR_BUFFER_BIT)
```
