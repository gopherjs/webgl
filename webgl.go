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
	Alpha, Depth, Stencil, Antialias, PremultipliedAlpha, PreserveDrawingBuffer bool
}

func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

type Context struct{ js.Object }

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

func (c *Context) IsContextLost() bool {
	return c.Call("isContextLost").Bool()
}

func (c *Context) GetSupportedExtensions() []string {
	ext := c.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

func (c *Context) GetExtension(name string) js.Object {
	return c.Call("getExtension", name)
}

func (c *Context) ActiveTexture(texture int) {
	c.Call("activeTexture", texture)
}

func (c *Context) AttachShader(program js.Object, shader js.Object) {
	c.Call("attachShader", program, shader)
}

func (c *Context) BindAttribLocation(program js.Object, index int, name string) {
	c.Call("bindAttribLocation", program, index, name)
}

func (c *Context) BindBuffer(target int, buffer js.Object) {
	c.Call("bindBuffer", target, buffer)
}

func (c *Context) BindFramebuffer(target int, framebuffer js.Object) {
	c.Call("bindFramebuffer", target, framebuffer)
}

func (c *Context) BindRenderbuffer(target int, renderbuffer js.Object) {
	c.Call("bindRenderbuffer", target, renderbuffer)
}

func (c *Context) BindTexture(target int, texture js.Object) {
	c.Call("bindTexture", target, texture)
}

func (c *Context) BlendColor(r, g, b, a float64) {
	c.Call("blendColor", r, g, b, a)
}

func (c *Context) BlendEquation(mode int) {
	c.Call("blendEquation", mode)
}

func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Call("blendFunc", sfactor, dfactor)
}

func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (c *Context) BufferData(target int, data js.Object, usage int) {
	c.Call("bufferData", target, data, usage)
}

func (c *Context) BufferSubData(target int, offset int, data js.Object) {
	c.Call("bufferSubData", target, offset, data)
}

func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Call("checkFramebufferStatus", target).Int()
}

func (c *Context) Clear(flags int) {
	c.Call("clear", flags)
}

func (c *Context) ClearColor(r, g, b, a float64) {
	c.Call("clearColor", r, g, b, a)
}

func (c *Context) ClearDepth(depth float64) {
	c.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Call("clearStencil", s)
}

func (c *Context) ColorMask(r, g, b, a bool) {
	c.Call("colorMask", r, g, b, a)
}

func (c *Context) CompileShader(shader js.Object) {
	c.Call("compileShader", shader)
}

// public function copyTexImage2D(target:GLenum, level:GLint, internalformat:GLenum, x:GLint, y:GLint,  width:GLsizei, height:GLsizei, border:GLint) : Void;
// public function copyTexSubImage2D(target:GLenum, level:GLint, xoffset:GLint, yoffset:GLint, x:GLint, y:GLint, width:GLsizei, height:GLsizei) : Void;

func (c *Context) CreateBuffer() js.Object {
	return c.Call("createBuffer")
}

func (c *Context) CreateFramebuffer() js.Object {
	return c.Call("createFramebuffer")
}

func (c *Context) CreateProgram() js.Object {
	return c.Call("createProgram")
}

func (c *Context) CreateRenderbuffer() js.Object {
	return c.Call("createRenderbuffer")
}

func (c *Context) CreateShader(typ int) js.Object {
	return c.Call("createShader", typ)
}

func (c *Context) CreateTexture() js.Object {
	return c.Call("createTexture")
}

// public function cullFace(mode:GLenum) : Void;
// public function deleteBuffer(buffer:WebGLBuffer) : Void;
// public function deleteFramebuffer(framebuffer:WebGLFramebuffer) : Void;
// public function deleteProgram(program:WebGLProgram) : Void;
// public function deleteRenderbuffer(renderbuffer:WebGLRenderbuffer) : Void;

func (c *Context) DeleteShader(shader js.Object) {
	c.Call("deleteShader", shader)
}

// public function deleteTexture(texture:WebGLTexture) : Void;
// public function depthFunc(func:GLenum) : Void;
// public function depthMask(flag:GLboolean) : Void;
// public function depthRange(zNear:GLclampf, zFar:GLclampf) : Void;
// public function detachShader(program:WebGLProgram, shader:WebGLShader) : Void;

func (c *Context) Disable(cap int) {
	c.Call("disable", cap)
}

// public function disableVertexAttribArray(index:GLuint) : Void;

func (c *Context) DrawArrays(mode, first, count int) {
	c.Call("drawArrays", mode, first, count)
}

// public function drawElements(mode:GLenum, count:GLsizei, type:GLenum, offset:GLintptr) : Void;

func (c *Context) Enable(cap int) {
	c.Call("enable", cap)
}

func (c *Context) EnableVertexAttribArray(index int) {
	c.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Call("finish")
}

func (c *Context) Flush() {
	c.Call("flush")
}

// public function framebufferRenderbuffer(target:GLenum, attachment:GLenum, renderbuffertarget:GLenum, renderbuffer:WebGLRenderbuffer) : Void;
// public function framebufferTexture2D(target:GLenum, attachment:GLenum, textarget:GLenum, texture:WebGLTexture, level:GLint) : Void;

func (c *Context) FrontFace(mode int) {
	c.Call("frontFace", mode)
}

func (c *Context) GenerateMipmap(target int) {
	c.Call("generateMipmap", target)
}

// public function getActiveAttrib(program:WebGLProgram, index:GLuint) : WebGLActiveInfo;
// public function getActiveUniform(program:WebGLProgram, index:GLuint) : WebGLActiveInfo;
// public function getAttachedShaders(program:WebGLProgram) : Array<WebGLShader>;

func (c *Context) GetAttribLocation(program js.Object, name string) int {
	return c.Call("getAttribLocation", program, name).Int()
}

// public function getParameter(pname:GLenum) : Dynamic;
// public function getBufferParameter(target:GLenum, pname:GLenum) : Dynamic;

func (c *Context) GetError() int {
	return c.Call("getError").Int()
}

// public function getFramebufferAttachmentParameter(target:GLenum, attachment:GLenum, pname:GLenum) : Dynamic;

func (c *Context) GetProgramParameteri(program js.Object, pname int) int {
	return c.Call("getProgramParameter", program, pname).Int()
}

func (c *Context) GetProgramParameterb(program js.Object, pname int) bool {
	return c.Call("getProgramParameter", program, pname).Bool()
}

func (c *Context) GetProgramInfoLog(program js.Object) string {
	return c.Call("getProgramInfoLog", program).String()
}

// public function getRenderbufferParameter(target:GLenum, pname:GLenum) : Dynamic;
// public function getShaderParameter(shader:WebGLShader, pname:GLenum) : Dynamic;

func (c *Context) GetShaderParameterb(shader js.Object, pname int) bool {
	return c.Call("getShaderParameter", shader, pname).Bool()
}

func (c *Context) GetShaderInfoLog(shader js.Object) string {
	return c.Call("getShaderInfoLog", shader).String()
}

func (c *Context) GetShaderSource(shader js.Object) string {
	return c.Call("getShaderSource", shader).String()
}

// public function getTexParameter(target:GLenum, pname:GLenum) : Dynamic;
// public function getUniform(program:WebGLProgram, location:WebGLUniformLocation) : Dynamic;

func (c *Context) GetUniformLocation(program js.Object, name string) js.Object {
	return c.Call("getUniformLocation", program, name)
}

// public function  getVertexAttrib(index:GLuint, pname:GLenum) : Dynamic;
// public function getVertexAttribOffset(index:GLuint, pname:GLenum) : GLsizeiptr;
// public function hint(target:GLenum, mode:GLenum) : Void;
// public function isBuffer(buffer:WebGLBuffer) : GLboolean;
// public function isEnabled(cap:GLenum) : GLboolean;
// public function isFramebuffer(framebuffer:WebGLFramebuffer) : GLboolean;
// public function isProgram(program:WebGLProgram) : GLboolean;
// public function  isRenderbuffer(renderbuffer:WebGLRenderbuffer) : GLboolean;
// public function isShader(shader:WebGLShader) : GLboolean;
// public function isTexture(texture:WebGLTexture) : GLboolean;

func (c *Context) LineWidth(width float64) {
	c.Call("lineWidth", width)
}

func (c *Context) LinkProgram(program js.Object) {
	c.Call("linkProgram", program)
}

func (c *Context) PixelStorei(pname, param int) {
	c.Call("pixelStorei", pname, param)
}

// public function polygonOffset(factor:GLfloat, units:GLfloat) : Void;
// public function readPixels(x:GLint, y:GLint, width:GLsizei, height:GLsizei, format:GLenum, type:GLenum, pixels:ArrayBufferView) : Void;
// public function renderbufferStorage(target:GLenum, internalformat:GLenum, width:GLsizei, height:GLsizei) : Void;
// public function sampleCoverage(value:GLclampf, invert:GLboolean) : Void;
// public function scissor(x:GLint, y:GLint, width:GLsizei, height:GLsizei) : Void;

func (c *Context) ShaderSource(shader js.Object, source string) {
	c.Call("shaderSource", shader, source)
}

// public function stencilFunc(func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilFuncSeparate(face:GLenum, func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;
// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, image js.Object) {
	c.Call("texImage2D", target, level, internalFormat, format, kind, image)
}

func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Call("texParameteri", target, pname, param)
}

// public function texSubImage2D(target:GLenum, level:GLint, xoffset:GLint , yoffset:GLint, format:GLenum, type:GLenum, video:HTMLVideoElement) : Void;
// public function uniform1f(location:WebGLUniformLocation, x:GLfloat) : Void;
// public function uniform1i(location:WebGLUniformLocation, x:GLint) : Void;
// public function uniform2f(location:WebGLUniformLocation, x:GLfloat, y:GLfloat) : Void;
// public function uniform2i(location:WebGLUniformLocation, x:GLint, y:GLint) : Void;
// public function uniform3f(location:WebGLUniformLocation, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function uniform3i(location:WebGLUniformLocation, x:GLint, y:GLint, z:GLint) : Void;
// public function uniform4f(location:WebGLUniformLocation, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;
// public function uniform4i(location:WebGLUniformLocation, x:GLint, y:GLint, z:GLint, w:GLint) : Void;
// public function uniform1fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform1iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform2fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform2iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform3fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform3iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform4fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform4iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniformMatrix2fv(location:WebGLUniformLocation, transpose:GLboolean, value:ArrayAccess<Float>) : Void;
// public function  uniformMatrix3fv(location:WebGLUniformLocation, transpose:GLboolean, value:ArrayAccess<Float>) : Void;

func (c *Context) UniformMatrix4fv(location js.Object, transpose bool, value []float32) {
	c.Call("uniformMatrix4fv", location, transpose, value)
}

func (c *Context) UseProgram(program js.Object) {
	c.Call("useProgram", program)
}

func (c *Context) ValidateProgram(program js.Object) {
	c.Call("validateProgram", program)
}

// public function vertexAttrib1f(indx:GLuint, x:GLfloat) : Void;
// public function vertexAttrib2f(indx:GLuint, x:GLfloat, y:GLfloat) : Void;
// public function vertexAttrib3f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function vertexAttrib4f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

// public function vertexAttrib1fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib2fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib3fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib4fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function viewport(x:GLint, y:GLint, width:GLsizei, height:GLsizei) : Void;
