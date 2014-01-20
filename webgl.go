// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webgl

import (
	"errors"
	"github.com/neelance/gopherjs/js"
)

const (
	ARRAY_BUFFER                                 = 34962
	ARRAY_BUFFER_BINDING                         = 34964
	ATTACHED_SHADERS                             = 35717
	BACK                                         = 1029
	BLEND                                        = 3042
	BLEND_COLOR                                  = 32773
	BLEND_DST_ALPHA                              = 32970
	BLEND_DST_RGB                                = 32968
	BLEND_EQUATION                               = 32777
	BLEND_EQUATION_ALPHA                         = 34877
	BLEND_EQUATION_RGB                           = 32777
	BLEND_SRC_ALPHA                              = 32971
	BLEND_SRC_RGB                                = 32969
	BLUE_BITS                                    = 3412
	BOOL                                         = 35670
	BOOL_VEC2                                    = 35671
	BOOL_VEC3                                    = 35672
	BOOL_VEC4                                    = 35673
	BROWSER_DEFAULT_WEBGL                        = 37444
	BUFFER_SIZE                                  = 34660
	BUFFER_USAGE                                 = 34661
	BYTE                                         = 5120
	CCW                                          = 2305
	CLAMP_TO_EDGE                                = 33071
	COLOR_ATTACHMENT0                            = 36064
	COLOR_BUFFER_BIT                             = 16384
	COLOR_CLEAR_VALUE                            = 3106
	COLOR_WRITEMASK                              = 3107
	COMPILE_STATUS                               = 35713
	COMPRESSED_TEXTURE_FORMATS                   = 34467
	CONSTANT_ALPHA                               = 32771
	CONSTANT_COLOR                               = 32769
	CONTEXT_LOST_WEBGL                           = 37442
	CULL_FACE                                    = 2884
	CULL_FACE_MODE                               = 2885
	CURRENT_PROGRAM                              = 35725
	CURRENT_VERTEX_ATTRIB                        = 34342
	CW                                           = 2304
	DECR                                         = 7683
	DECR_WRAP                                    = 34056
	DELETE_STATUS                                = 35712
	DEPTH_ATTACHMENT                             = 36096
	DEPTH_BITS                                   = 3414
	DEPTH_BUFFER_BIT                             = 256
	DEPTH_CLEAR_VALUE                            = 2931
	DEPTH_COMPONENT                              = 6402
	DEPTH_COMPONENT16                            = 33189
	DEPTH_FUNC                                   = 2932
	DEPTH_RANGE                                  = 2928
	DEPTH_STENCIL                                = 34041
	DEPTH_STENCIL_ATTACHMENT                     = 33306
	DEPTH_TEST                                   = 2929
	DEPTH_WRITEMASK                              = 2930
	DITHER                                       = 3024
	DONT_CARE                                    = 4352
	DST_ALPHA                                    = 772
	DST_COLOR                                    = 774
	DYNAMIC_DRAW                                 = 35048
	ELEMENT_ARRAY_BUFFER                         = 34963
	ELEMENT_ARRAY_BUFFER_BINDING                 = 34965
	EQUAL                                        = 514
	FASTEST                                      = 4353
	FLOAT                                        = 5126
	FLOAT_MAT2                                   = 35674
	FLOAT_MAT3                                   = 35675
	FLOAT_MAT4                                   = 35676
	FLOAT_VEC2                                   = 35664
	FLOAT_VEC3                                   = 35665
	FLOAT_VEC4                                   = 35666
	FRAGMENT_SHADER                              = 35632
	FRAMEBUFFER                                  = 36160
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = 36049
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = 36048
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 36051
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = 36050
	FRAMEBUFFER_BINDING                          = 36006
	FRAMEBUFFER_COMPLETE                         = 36053
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = 36054
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            = 36057
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = 36055
	FRAMEBUFFER_UNSUPPORTED                      = 36061
	FRONT                                        = 1028
	FRONT_AND_BACK                               = 1032
	FRONT_FACE                                   = 2886
	FUNC_ADD                                     = 32774
	FUNC_REVERSE_SUBTRACT                        = 32779
	FUNC_SUBTRACT                                = 32778
	GENERATE_MIPMAP_HINT                         = 33170
	GEQUAL                                       = 518
	GREATER                                      = 516
	GREEN_BITS                                   = 3411
	HIGH_FLOAT                                   = 36338
	HIGH_INT                                     = 36341
	INCR                                         = 7682
	INCR_WRAP                                    = 34055
	INFO_LOG_LENGTH                              = 35716
	INT                                          = 5124
	INT_VEC2                                     = 35667
	INT_VEC3                                     = 35668
	INT_VEC4                                     = 35669
	INVALID_ENUM                                 = 1280
	INVALID_FRAMEBUFFER_OPERATION                = 1286
	INVALID_OPERATION                            = 1282
	INVALID_VALUE                                = 1281
	INVERT                                       = 5386
	KEEP                                         = 7680
	LEQUAL                                       = 515
	LESS                                         = 513
	LINEAR                                       = 9729
	LINEAR_MIPMAP_LINEAR                         = 9987
	LINEAR_MIPMAP_NEAREST                        = 9985
	LINES                                        = 1
	LINE_LOOP                                    = 2
	LINE_STRIP                                   = 3
	LINE_WIDTH                                   = 2849
	LINK_STATUS                                  = 35714
	LOW_FLOAT                                    = 36336
	LOW_INT                                      = 36339
	LUMINANCE                                    = 6409
	LUMINANCE_ALPHA                              = 6410
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             = 35661
	MAX_CUBE_MAP_TEXTURE_SIZE                    = 34076
	MAX_FRAGMENT_UNIFORM_VECTORS                 = 36349
	MAX_RENDERBUFFER_SIZE                        = 34024
	MAX_TEXTURE_IMAGE_UNITS                      = 34930
	MAX_TEXTURE_SIZE                             = 3379
	MAX_VARYING_VECTORS                          = 36348
	MAX_VERTEX_ATTRIBS                           = 34921
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               = 35660
	MAX_VERTEX_UNIFORM_VECTORS                   = 36347
	MAX_VIEWPORT_DIMS                            = 3386
	MEDIUM_FLOAT                                 = 36337
	MEDIUM_INT                                   = 36340
	MIRRORED_REPEAT                              = 33648
	NEAREST                                      = 9728
	NEAREST_MIPMAP_LINEAR                        = 9986
	NEAREST_MIPMAP_NEAREST                       = 9984
	NEVER                                        = 512
	NICEST                                       = 4354
	NONE                                         = 0
	NOTEQUAL                                     = 517
	NO_ERROR                                     = 0
	NUM_COMPRESSED_TEXTURE_FORMATS               = 34466
	ONE                                          = 1
	ONE_MINUS_CONSTANT_ALPHA                     = 32772
	ONE_MINUS_CONSTANT_COLOR                     = 32770
	ONE_MINUS_DST_ALPHA                          = 773
	ONE_MINUS_DST_COLOR                          = 775
	ONE_MINUS_SRC_ALPHA                          = 771
	ONE_MINUS_SRC_COLOR                          = 769
	OUT_OF_MEMORY                                = 1285
	PACK_ALIGNMENT                               = 3333
	POINTS                                       = 0
	POLYGON_OFFSET_FACTOR                        = 32824
	POLYGON_OFFSET_FILL                          = 32823
	POLYGON_OFFSET_UNITS                         = 10752
	RED_BITS                                     = 3410
	RENDERBUFFER                                 = 36161
	RENDERBUFFER_ALPHA_SIZE                      = 36179
	RENDERBUFFER_BINDING                         = 36007
	RENDERBUFFER_BLUE_SIZE                       = 36178
	RENDERBUFFER_DEPTH_SIZE                      = 36180
	RENDERBUFFER_GREEN_SIZE                      = 36177
	RENDERBUFFER_HEIGHT                          = 36163
	RENDERBUFFER_INTERNAL_FORMAT                 = 36164
	RENDERBUFFER_RED_SIZE                        = 36176
	RENDERBUFFER_STENCIL_SIZE                    = 36181
	RENDERBUFFER_WIDTH                           = 36162
	RENDERER                                     = 7937
	REPEAT                                       = 10497
	REPLACE                                      = 7681
	RGB                                          = 6407
	RGB5_A1                                      = 32855
	RGB565                                       = 36194
	RGBA                                         = 6408
	RGBA4                                        = 32854
	SAMPLER_2D                                   = 35678
	SAMPLER_CUBE                                 = 35680
	SAMPLES                                      = 32937
	SAMPLE_ALPHA_TO_COVERAGE                     = 32926
	SAMPLE_BUFFERS                               = 32936
	SAMPLE_COVERAGE                              = 32928
	SAMPLE_COVERAGE_INVERT                       = 32939
	SAMPLE_COVERAGE_VALUE                        = 32938
	SCISSOR_BOX                                  = 3088
	SCISSOR_TEST                                 = 3089
	SHADER_COMPILER                              = 36346
	SHADER_SOURCE_LENGTH                         = 35720
	SHADER_TYPE                                  = 35663
	SHADING_LANGUAGE_VERSION                     = 35724
	SHORT                                        = 5122
	SRC_ALPHA                                    = 770
	SRC_ALPHA_SATURATE                           = 776
	SRC_COLOR                                    = 768
	STATIC_DRAW                                  = 35044
	STENCIL_ATTACHMENT                           = 36128
	STENCIL_BACK_FAIL                            = 34817
	STENCIL_BACK_FUNC                            = 34816
	STENCIL_BACK_PASS_DEPTH_FAIL                 = 34818
	STENCIL_BACK_PASS_DEPTH_PASS                 = 34819
	STENCIL_BACK_REF                             = 36003
	STENCIL_BACK_VALUE_MASK                      = 36004
	STENCIL_BACK_WRITEMASK                       = 36005
	STENCIL_BITS                                 = 3415
	STENCIL_BUFFER_BIT                           = 1024
	STENCIL_CLEAR_VALUE                          = 2961
	STENCIL_FAIL                                 = 2964
	STENCIL_FUNC                                 = 2962
	STENCIL_INDEX                                = 6401
	STENCIL_INDEX8                               = 36168
	STENCIL_PASS_DEPTH_FAIL                      = 2965
	STENCIL_PASS_DEPTH_PASS                      = 2966
	STENCIL_REF                                  = 2967
	STENCIL_TEST                                 = 2960
	STENCIL_VALUE_MASK                           = 2963
	STENCIL_WRITEMASK                            = 2968
	STREAM_DRAW                                  = 35040
	SUBPIXEL_BITS                                = 3408
	TEXTURE                                      = 5890
	TEXTURE0                                     = 33984
	TEXTURE1                                     = 33985
	TEXTURE2                                     = 33986
	TEXTURE3                                     = 33987
	TEXTURE4                                     = 33988
	TEXTURE5                                     = 33989
	TEXTURE6                                     = 33990
	TEXTURE7                                     = 33991
	TEXTURE8                                     = 33992
	TEXTURE9                                     = 33993
	TEXTURE10                                    = 33994
	TEXTURE11                                    = 33995
	TEXTURE12                                    = 33996
	TEXTURE13                                    = 33997
	TEXTURE14                                    = 33998
	TEXTURE15                                    = 33999
	TEXTURE16                                    = 34000
	TEXTURE17                                    = 34001
	TEXTURE18                                    = 34002
	TEXTURE19                                    = 34003
	TEXTURE20                                    = 34004
	TEXTURE21                                    = 34005
	TEXTURE22                                    = 34006
	TEXTURE23                                    = 34007
	TEXTURE24                                    = 34008
	TEXTURE25                                    = 34009
	TEXTURE26                                    = 34010
	TEXTURE27                                    = 34011
	TEXTURE28                                    = 34012
	TEXTURE29                                    = 34013
	TEXTURE30                                    = 34014
	TEXTURE31                                    = 34015
	TEXTURE_2D                                   = 3553
	TEXTURE_BINDING_2D                           = 32873
	TEXTURE_BINDING_CUBE_MAP                     = 34068
	TEXTURE_CUBE_MAP                             = 34067
	TEXTURE_CUBE_MAP_NEGATIVE_X                  = 34070
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  = 34072
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  = 34074
	TEXTURE_CUBE_MAP_POSITIVE_X                  = 34069
	TEXTURE_CUBE_MAP_POSITIVE_Y                  = 34071
	TEXTURE_CUBE_MAP_POSITIVE_Z                  = 34073
	TEXTURE_MAG_FILTER                           = 10240
	TEXTURE_MIN_FILTER                           = 10241
	TEXTURE_WRAP_S                               = 10242
	TEXTURE_WRAP_T                               = 10243
	TRIANGLES                                    = 4
	TRIANGLE_FAN                                 = 6
	TRIANGLE_STRIP                               = 5
	UNPACK_ALIGNMENT                             = 3317
	UNPACK_COLORSPACE_CONVERSION_WEBGL           = 37443
	UNPACK_FLIP_Y_WEBGL                          = 37440
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               = 37441
	UNSIGNED_BYTE                                = 5121
	UNSIGNED_INT                                 = 5125
	UNSIGNED_SHORT                               = 5123
	UNSIGNED_SHORT_4_4_4_4                       = 32819
	UNSIGNED_SHORT_5_5_5_1                       = 32820
	UNSIGNED_SHORT_5_6_5                         = 33635
	VALIDATE_STATUS                              = 35715
	VENDOR                                       = 7936
	VERSION                                      = 7938
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           = 34975
	VERTEX_ATTRIB_ARRAY_ENABLED                  = 34338
	VERTEX_ATTRIB_ARRAY_NORMALIZED               = 34922
	VERTEX_ATTRIB_ARRAY_POINTER                  = 34373
	VERTEX_ATTRIB_ARRAY_SIZE                     = 34339
	VERTEX_ATTRIB_ARRAY_STRIDE                   = 34340
	VERTEX_ATTRIB_ARRAY_TYPE                     = 34341
	VERTEX_SHADER                                = 35633
	VIEWPORT                                     = 2978
	ZERO                                         = 0
)

type ContextAttributes struct {
	// If Alpha is true, the drawing buffer has an alpha channel for
	// the purposes of performing OpenGL destination alpha operations
	// and compositing with the page.
	Alpha bool

	// If Depth is true, the drawing buffer has a depth buffer of at least 16 bits.
	Depth bool

	// If Stencil is true, the drawing buffer has a stencil buffer of at least 8 bits.
	Stencil bool

	// If Antialias is true and the implementation supports antialiasing
	// the drawing buffer will perform antialiasing using its choice of
	// technique (multisample/supersample) and quality.
	Antialias bool

	// If PremultipliedAlpha is true the page compositor will assume the
	// drawing buffer contains colors with premultiplied alpha.
	// This flag is ignored if the alpha flag is false.
	PremultipliedAlpha bool

	// If the value is true the buffers will not be cleared and will preserve
	// their values until cleared or overwritten by the author.
	PreserveDrawingBuffer bool
}

// Returns a copy of the default WebGL context attributes.
func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

type Context struct{ js.Object }

// NewContext takes an HTML5 canvas object and optional context attributes.
// If an error is returned it means you won't have access to WebGL
// functionality.
func NewContext(canvas js.Object, ca *ContextAttributes) (*Context, error) {
	if js.Global("window").Get("WebGLRenderingContext").IsUndefined() {
		return nil, errors.New("Your browser doesn't appear to support webgl.")
	}

	if ca == nil {
		ca = DefaultAttributes()
	}

	attrs := map[string]bool{
		"alpha":                 ca.Alpha,
		"depth":                 ca.Depth,
		"stencil":               ca.Stencil,
		"antialias":             ca.Antialias,
		"premultipliedAlpha":    ca.PremultipliedAlpha,
		"preserveDrawingBuffer": ca.PreserveDrawingBuffer,
	}
	gl := canvas.Call("getContext", "webgl", attrs)
	if gl.IsNull() {
		gl = canvas.Call("getContext", "experimental-webgl", attrs)
		if gl.IsNull() {
			return nil, errors.New("Creating a webgl context has failed.")
		}
	}
	return &Context{gl}, nil
}

// Returns the context attributes active on the context. These values might
// be different than what was requested on context creation if the
// browser's implementation doesn't support a feature.
func (c *Context) GetContextAttributes() ContextAttributes {
	ca := c.Call("getContextAttributes")
	return ContextAttributes{
		ca.Get("alpha").Bool(),
		ca.Get("depth").Bool(),
		ca.Get("stencil").Bool(),
		ca.Get("antialias").Bool(),
		ca.Get("premultipliedAlpha").Bool(),
		ca.Get("preservedDrawingBuffer").Bool(),
	}
}

// Specifies the active texture unit.
func (c *Context) ActiveTexture(texture int) {
	c.Call("activeTexture", texture)
}

// Attaches a WebGLShader object to a WebGLProgram object.
func (c *Context) AttachShader(program js.Object, shader js.Object) {
	c.Call("attachShader", program, shader)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program js.Object, index int, name string) {
	c.Call("bindAttribLocation", program, index, name)
}

// Associates a buffer with a buffer target.
func (c *Context) BindBuffer(target int, buffer js.Object) {
	c.Call("bindBuffer", target, buffer)
}

// Associates a WebGLFramebuffer object with the FRAMEBUFFER bind target.
func (c *Context) BindFramebuffer(target int, framebuffer js.Object) {
	c.Call("bindFramebuffer", target, framebuffer)
}

// Binds a WebGLRenderbuffer object to be used for rendering.
func (c *Context) BindRenderbuffer(target int, renderbuffer js.Object) {
	c.Call("bindRenderbuffer", target, renderbuffer)
}

// Binds a named texture object to a target.
func (c *Context) BindTexture(target int, texture js.Object) {
	c.Call("bindTexture", target, texture)
}

// The GL_BLEND_COLOR may be used to calculate the source and destination blending factors.
func (c *Context) BlendColor(r, g, b, a float64) {
	c.Call("blendColor", r, g, b, a)
}

// Sets the equation used to blend RGB and Alpha values of an incoming source
// fragment with a destination values as stored in the fragment's frame buffer.
func (c *Context) BlendEquation(mode int) {
	c.Call("blendEquation", mode)
}

// Controls the blending of an incoming source fragment's R, G, B, and A values
// with a destination R, G, B, and A values as stored in the fragment's WebGLFramebuffer.
func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

// Sets the blending factors used to combine source and destination pixels.
func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Call("blendFunc", sfactor, dfactor)
}

// Sets the weighting factors that are used by blendEquationSeparate.
func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// Creates a buffer in memory and initializes it with array data.
// If no array is provided, the contents of the buffer is initialized to 0.
func (c *Context) BufferData(target int, data js.Object, usage int) {
	c.Call("bufferData", target, data, usage)
}

// Used to modify or update some or all of a data store for a bound buffer object.
func (c *Context) BufferSubData(target int, offset int, data js.Object) {
	c.Call("bufferSubData", target, offset, data)
}

// Returns whether the currently bound WebGLFramebuffer is complete.
// If not complete, returns the reason why.
func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Call("checkFramebufferStatus", target).Int()
}

// Sets all pixels in a specific buffer to the same value.
func (c *Context) Clear(flags int) {
	c.Call("clear", flags)
}

// Specifies color values to use by the clear method to clear the color buffer.
func (c *Context) ClearColor(r, g, b, a float32) {
	c.Call("clearColor", r, g, b, a)
}

// Clears the depth buffer to a specific value.
func (c *Context) ClearDepth(depth float64) {
	c.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Call("clearStencil", s)
}

// Lets you set whether individual colors can be written when
// drawing or rendering to a framebuffer.
func (c *Context) ColorMask(r, g, b, a bool) {
	c.Call("colorMask", r, g, b, a)
}

// Compiles the GLSL shader source into binary data used by the WebGLProgram object.
func (c *Context) CompileShader(shader js.Object) {
	c.Call("compileShader", shader)
}

// Copies a rectangle of pixels from the current WebGLFramebuffer into a texture image.
func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.Call("copyTexImage2D", target, level, internal, x, y, w, h, border)
}

// Replaces a portion of an existing 2D texture image with data from the current framebuffer.
func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, w, h)
}

// Creates and initializes a WebGLBuffer.
func (c *Context) CreateBuffer() js.Object {
	return c.Call("createBuffer")
}

// Returns a WebGLFramebuffer object.
func (c *Context) CreateFramebuffer() js.Object {
	return c.Call("createFramebuffer")
}

// Creates an empty WebGLProgram object to which vector and fragment
// WebGLShader objects can be bound.
func (c *Context) CreateProgram() js.Object {
	return c.Call("createProgram")
}

// Creates and returns a WebGLRenderbuffer object.
func (c *Context) CreateRenderbuffer() js.Object {
	return c.Call("createRenderbuffer")
}

// Returns an empty vertex or fragment shader object based on the type specified.
func (c *Context) CreateShader(typ int) js.Object {
	return c.Call("createShader", typ)
}

// Used to generate a WebGLTexture object to which images can be bound.
func (c *Context) CreateTexture() js.Object {
	return c.Call("createTexture")
}

// Sets whether or not front, back, or both facing facets are able to be culled.
func (c *Context) CullFace(mode int) {
	c.Call("cullFace", mode)
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer js.Object) {
	c.Call("deleteBuffer", buffer)
}

// Deletes a specific WebGLFramebuffer object. If you delete the
// currently bound framebuffer, the default framebuffer will be bound.
// Deleting a framebuffer detaches all of its attachments.
func (c *Context) DeleteFramebuffer(framebuffer js.Object) {
	c.Call("deleteFramebuffer", framebuffer)
}

// Flags a specific WebGLProgram object for deletion if currently active.
// It will be deleted when it is no longer being used.
// Any shader objects associated with the program will be detached.
// They will be deleted if they were already flagged for deletion.
func (c *Context) DeleteProgram(program js.Object) {
	c.Call("deleteProgram", program)
}

// Deletes the specified renderbuffer object. If the renderbuffer is
// currently bound, it will become unbound. If the renderbuffer is
// attached to the currently bound framebuffer, it is detached.
func (c *Context) DeleteRenderbuffer(renderbuffer js.Object) {
	c.Call("deleteRenderbuffer", renderbuffer)
}

// Deletes a specific shader object.
func (c *Context) DeleteShader(shader js.Object) {
	c.Call("deleteShader", shader)
}

// Deletes a specific texture object.
func (c *Context) DeleteTexture(texture js.Object) {
	c.Call("deleteTexture", texture)
}

// Sets a function to use to compare incoming pixel depth to the
// current depth buffer value.
func (c *Context) DepthFunc(fun int) {
	c.Call("depthFunc", fun)
}

// Sets whether or not you can write to the depth buffer.
func (c *Context) DepthMask(flag bool) {
	c.Call("depthMask", flag)
}

// Sets the depth range for normalized coordinates to canvas or viewport depth coordinates.
func (c *Context) DepthRange(zNear, zFar float64) {
	c.Call("depthRange", zNear, zFar)
}

// Detach a shader object from a program object.
func (c *Context) DetachShader(program, shader js.Object) {
	c.Call("detachShader", program, shader)
}

// Turns off specific WebGL capabilities for this context.
func (c *Context) Disable(cap int) {
	c.Call("disable", cap)
}

// Turns off a vertex attribute array at a specific index position.
func (c *Context) DisableVertexAttribArray(index int) {
	c.Call("disableVertexAttribArray", index)
}

// Render geometric primitives from bound and enabled vertex data.
func (c *Context) DrawArrays(mode, first, count int) {
	c.Call("drawArrays", mode, first, count)
}

// Renders geometric primitives indexed by element array data.
func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.Call("drawElements", mode, count, typ, offset)
}

// Turns on specific WebGL capabilities for this context.
func (c *Context) Enable(cap int) {
	c.Call("enable", cap)
}

// Turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (c *Context) EnableVertexAttribArray(index int) {
	c.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Call("finish")
}

func (c *Context) Flush() {
	c.Call("flush")
}

// Attaches a WebGLRenderbuffer object as a logical buffer to the
// currently bound WebGLFramebuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer js.Object) {
	c.Call("framebufferRenderBuffer", target, attachment, renderbufferTarget, renderbuffer)
}

// Attaches a texture to a WebGLFramebuffer object.
func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture js.Object, level int) {
	c.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}

// Sets whether or not polygons are considered front-facing based
// on their winding direction.
func (c *Context) FrontFace(mode int) {
	c.Call("frontFace", mode)
}

// Creates a set of textures for a WebGLTexture object with image
// dimensions from the original size of the image down to a 1x1 image.
func (c *Context) GenerateMipmap(target int) {
	c.Call("generateMipmap", target)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a vertex attribute at a specific index position in a program object.
func (c *Context) GetActiveAttrib(program js.Object, index int) js.Object {
	return c.Call("getActiveAttrib", program, index)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a uniform attribute at a specific index position in a program object.
func (c *Context) GetActiveUniform(program js.Object, index int) js.Object {
	return c.Call("getActiveUniform", program, index)
}

// Returns a slice of WebGLShaders bound to a WebGLProgram.
func (c *Context) GetAttachedShaders(program js.Object) []js.Object {
	objs := c.Call("getAttachedShaders", program)
	shaders := make([]js.Object, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		shaders[i] = objs.Index(i)
	}
	return shaders
}

// Returns an index to the location in a program of a named attribute variable.
func (c *Context) GetAttribLocation(program js.Object, name string) int {
	return c.Call("getAttribLocation", program, name).Int()
}

// TODO: Create type specific variations.
// Returns the type of a parameter for a given buffer.
func (c *Context) GetBufferParameter(target, pname int) js.Object {
	return c.Call("getBufferParameter", target, pname)
}

// TODO: Create type specific variations.
// Returns the natural type value for a constant parameter.
func (c *Context) GetParameter(pname int) js.Object {
	return c.Call("getParameter", pname)
}

// Returns a value for the WebGL error flag and clears the flag.
func (c *Context) GetError() int {
	return c.Call("getError").Int()
}

// TODO: Create type specific variations.
// Enables a passed extension, otherwise returns null.
func (c *Context) GetExtension(name string) js.Object {
	return c.Call("getExtension", name)
}

// TODO: Create type specific variations.
// Gets a parameter value for a given target and attachment.
func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) js.Object {
	return c.Call("getFramebufferAttachmentParameter", target, attachment, pname)
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program js.Object, pname int) int {
	return c.Call("getProgramParameter", program, pname).Int()
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program js.Object, pname int) bool {
	return c.Call("getProgramParameter", program, pname).Bool()
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program js.Object) string {
	return c.Call("getProgramInfoLog", program).String()
}

// TODO: Create type specific variations.
// Returns a renderbuffer parameter from the currently bound WebGLRenderbuffer object.
func (c *Context) GetRenderbufferParameter(target, pname int) js.Object {
	return c.Call("getRenderbufferParameter", target, pname)
}

// TODO: Create type specific variations.
// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameter(shader js.Object, pname int) js.Object {
	return c.Call("getShaderParameter", shader, pname)
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameterb(shader js.Object, pname int) bool {
	return c.Call("getShaderParameter", shader, pname).Bool()
}

// Returns errors which occur when compiling a shader.
func (c *Context) GetShaderInfoLog(shader js.Object) string {
	return c.Call("getShaderInfoLog", shader).String()
}

// Returns source code string associated with a shader object.
func (c *Context) GetShaderSource(shader js.Object) string {
	return c.Call("getShaderSource", shader).String()
}

// Returns a slice of supported extension strings.
func (c *Context) GetSupportedExtensions() []string {
	ext := c.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

// TODO: Create type specific variations.
// Returns the value for a parameter on an active texture unit.
func (c *Context) GetTexParameter(target, pname int) js.Object {
	return c.Call("getTexParameter", target, pname)
}

// TODO: Create type specific variations.
// Gets the uniform value for a specific location in a program.
func (c *Context) GetUniform(program, location js.Object) js.Object {
	return c.Call("getUniform", program, location)
}

// Returns a WebGLUniformLocation object for the location
// of a uniform variable within a WebGLProgram object.
func (c *Context) GetUniformLocation(program js.Object, name string) js.Object {
	return c.Call("getUniformLocation", program, name)
}

// TODO: Create type specific variations.
// Returns data for a particular characteristic of a vertex
// attribute at an index in a vertex attribute array.
func (c *Context) GetVertexAttrib(index, pname int) js.Object {
	return c.Call("getVertexAttrib", index, pname)
}

// Returns the address of a specified vertex attribute.
func (c *Context) GetVertexAttribOffset(index, pname int) int {
	return c.Call("getVertexAttribOffset", index, pname).Int()
}

// public function hint(target:GLenum, mode:GLenum) : Void;

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsBuffer(buffer js.Object) bool {
	return c.Call("isBuffer", buffer).Bool()
}

// Returns whether the WebGL context has been lost.
func (c *Context) IsContextLost() bool {
	return c.Call("isContextLost").Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsFramebuffer(framebuffer js.Object) bool {
	return c.Call("isFramebuffer", framebuffer).Bool()
}

// Returns true if program object is valid, false otherwise.
func (c *Context) IsProgram(program js.Object) bool {
	return c.Call("isProgram", program).Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsRenderbuffer(renderbuffer js.Object) bool {
	return c.Call("isRenderbuffer", renderbuffer).Bool()
}

// Returns true if shader is valid, false otherwise.
func (c *Context) IsShader(shader js.Object) bool {
	return c.Call("isShader", shader).Bool()
}

// Returns true if texture is valid, false otherwise.
func (c *Context) IsTexture(texture js.Object) bool {
	return c.Call("isTexture", texture).Bool()
}

// Returns whether or not a WebGL capability is enabled for this context.
func (c *Context) IsEnabled(capability int) bool {
	return c.Call("isEnabled", capability).Bool()
}

// Sets the width of lines in WebGL.
func (c *Context) LineWidth(width float64) {
	c.Call("lineWidth", width)
}

// Links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (c *Context) LinkProgram(program js.Object) {
	c.Call("linkProgram", program)
}

// Sets pixel storage modes for readPixels and unpacking of textures
// with texImage2D and texSubImage2D.
func (c *Context) PixelStorei(pname, param int) {
	c.Call("pixelStorei", pname, param)
}

// Sets the implementation-specific units and scale factor
// used to calculate fragment depth values.
func (c *Context) PolygonOffset(factor, units float64) {
	c.Call("polygonOffset", factor, units)
}

// TODO: Figure out if pixels should be a slice.
// Reads pixel data into an ArrayBufferView object from a
// rectangular area in the color buffer of the active frame buffer.
func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels js.Object) {
	c.Call("readPixels", x, y, width, height, format, typ, pixels)
}

// Creates or replaces the data store for the currently bound WebGLRenderbuffer object.
func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.Call("renderbufferStorage", target, internalFormat, width, height)
}

//func (c *Context) SampleCoverage(value float64, invert bool) {
//	c.Call("sampleCoverage", value, invert)
//}

// Sets the dimensions of the scissor box.
func (c *Context) Scissor(x, y, width, height int) {
	c.Call("scissor", x, y, width, height)
}

// Sets and replaces shader source code in a shader object.
func (c *Context) ShaderSource(shader js.Object, source string) {
	c.Call("shaderSource", shader, source)
}

// public function stencilFunc(func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilFuncSeparate(face:GLenum, func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;
// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

// Loads the supplied pixel data into a texture.
func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, image js.Object) {
	c.Call("texImage2D", target, level, internalFormat, format, kind, image)
}

// Sets texture parameters for the current texture unit.
func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Call("texParameteri", target, pname, param)
}

// Replaces a portion of an existing 2D texture image with all of another image.
func (c *Context) TexSubImage2D(target, level, xoffset, yoffset, format, typ int, image js.Object) {
	c.Call("texSubImage2D", target, level, xoffset, yoffset, format, typ, image)
}

// Assigns a floating point value to a uniform variable for the current program object.
func (c *Context) Uniform1f(location js.Object, x float32) {
	c.Call("uniform1f", location, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location js.Object, x int) {
	c.Call("uniform1i", location, x)
}

// Assigns 2 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform2f(location js.Object, x, y float32) {
	c.Call("uniform2f", location, x, y)
}

// Assigns 2 integer values to a uniform variable for the current program object.
func (c *Context) Uniform2i(location js.Object, x, y int) {
	c.Call("uniform2i", location, x, y)
}

// Assigns 3 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform3f(location js.Object, x, y, z float32) {
	c.Call("uniform3f", location, x, y, z)
}

// Assigns 3 integer values to a uniform variable for the current program object.
func (c *Context) Uniform3i(location js.Object, x, y, z int) {
	c.Call("uniform3i", location, x, y, z)
}

// Assigns 4 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform4f(location js.Object, x, y, z, w float32) {
	c.Call("uniform4f", location, x, y, z, w)
}

// Assigns 4 integer values to a uniform variable for the current program object.
func (c *Context) Uniform4i(location js.Object, x, y, z, w int) {
	c.Call("uniform4i", location, x, y, z, w)
}

// public function uniform1fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform1iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform2fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform2iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform3fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform3iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform4fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform4iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;

// Sets values for a 2x2 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix2fv(location js.Object, transpose bool, value []float32) {
	c.Call("uniformMatrix2fv", location, transpose, value)
}

// Sets values for a 3x3 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix3fv(location js.Object, transpose bool, value []float32) {
	c.Call("uniformMatrix3fv", location, transpose, value)
}

// Sets values for a 4x4 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix4fv(location js.Object, transpose bool, value []float32) {
	c.Call("uniformMatrix4fv", location, transpose, value)
}

// Set the program object to use for rendering.
func (c *Context) UseProgram(program js.Object) {
	c.Call("useProgram", program)
}

// Returns whether a given program can run in the current WebGL state.
func (c *Context) ValidateProgram(program js.Object) {
	c.Call("validateProgram", program)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

// public function vertexAttrib1f(indx:GLuint, x:GLfloat) : Void;
// public function vertexAttrib2f(indx:GLuint, x:GLfloat, y:GLfloat) : Void;
// public function vertexAttrib3f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function vertexAttrib4f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;
// public function vertexAttrib1fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib2fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib3fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib4fv(indx:GLuint, values:ArrayAccess<Float>) : Void;

// Represents a rectangular viewable area that contains
// the rendering results of the drawing buffer.
func (c *Context) Viewport(x, y, width, height int) {
	c.Call("viewport", x, y, width, height)
}
