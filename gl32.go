// https://www.opengl.org/sdk/docs/man3
package gl32

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
//
// #include <stdlib.h>
// #include <OpenGL/gl3.h>
import "C"
import "unsafe"

type (
	Enum     C.GLenum
	Boolean  C.GLboolean
	Bitfield C.GLbitfield
	Byte     C.GLbyte
	Short    C.GLshort
	Int      C.GLint
	Sizei    C.GLsizei
	Ubyte    C.GLubyte
	Ushort   C.GLushort
	Uint     C.GLuint
	Half     C.GLhalf
	Float    C.GLfloat
	Clampf   C.GLclampf
	Double   C.GLdouble
	Clampd   C.GLclampd
	Char     C.GLchar
	Pointer  unsafe.Pointer
	Sync     C.GLsync
	Int64    C.GLint64
	Uint64   C.GLuint64
	Intptr   C.GLintptr
	Sizeiptr C.GLsizeiptr
)


// VERSION_1_1
const (
	ALPHA = C.GL_ALPHA
	ALWAYS = C.GL_ALWAYS
	AND = C.GL_AND
	AND_INVERTED = C.GL_AND_INVERTED
	AND_REVERSE = C.GL_AND_REVERSE
	BACK = C.GL_BACK
	BACK_LEFT = C.GL_BACK_LEFT
	BACK_RIGHT = C.GL_BACK_RIGHT
	BLEND = C.GL_BLEND
	BLEND_DST = C.GL_BLEND_DST
	BLEND_SRC = C.GL_BLEND_SRC
	BLUE = C.GL_BLUE
	BYTE = C.GL_BYTE
	CCW = C.GL_CCW
	CLEAR = C.GL_CLEAR
	COLOR = C.GL_COLOR
	COLOR_BUFFER_BIT = C.GL_COLOR_BUFFER_BIT
	COLOR_CLEAR_VALUE = C.GL_COLOR_CLEAR_VALUE
	COLOR_LOGIC_OP = C.GL_COLOR_LOGIC_OP
	COLOR_WRITEMASK = C.GL_COLOR_WRITEMASK
	COPY = C.GL_COPY
	COPY_INVERTED = C.GL_COPY_INVERTED
	CULL_FACE = C.GL_CULL_FACE
	CULL_FACE_MODE = C.GL_CULL_FACE_MODE
	CW = C.GL_CW
	DECR = C.GL_DECR
	DEPTH = C.GL_DEPTH
	DEPTH_BUFFER_BIT = C.GL_DEPTH_BUFFER_BIT
	DEPTH_CLEAR_VALUE = C.GL_DEPTH_CLEAR_VALUE
	DEPTH_COMPONENT = C.GL_DEPTH_COMPONENT
	DEPTH_FUNC = C.GL_DEPTH_FUNC
	DEPTH_RANGE = C.GL_DEPTH_RANGE
	DEPTH_TEST = C.GL_DEPTH_TEST
	DEPTH_WRITEMASK = C.GL_DEPTH_WRITEMASK
	DITHER = C.GL_DITHER
	DONT_CARE = C.GL_DONT_CARE
	DOUBLE = C.GL_DOUBLE
	DOUBLEBUFFER = C.GL_DOUBLEBUFFER
	DRAW_BUFFER = C.GL_DRAW_BUFFER
	DST_ALPHA = C.GL_DST_ALPHA
	DST_COLOR = C.GL_DST_COLOR
	EQUAL = C.GL_EQUAL
	EQUIV = C.GL_EQUIV
	EXTENSIONS = C.GL_EXTENSIONS
	FALSE = C.GL_FALSE
	FASTEST = C.GL_FASTEST
	FILL = C.GL_FILL
	FLOAT = C.GL_FLOAT
	FRONT = C.GL_FRONT
	FRONT_AND_BACK = C.GL_FRONT_AND_BACK
	FRONT_FACE = C.GL_FRONT_FACE
	FRONT_LEFT = C.GL_FRONT_LEFT
	FRONT_RIGHT = C.GL_FRONT_RIGHT
	GEQUAL = C.GL_GEQUAL
	GREATER = C.GL_GREATER
	GREEN = C.GL_GREEN
	INCR = C.GL_INCR
	INT = C.GL_INT
	INVALID_ENUM = C.GL_INVALID_ENUM
	INVALID_OPERATION = C.GL_INVALID_OPERATION
	INVALID_VALUE = C.GL_INVALID_VALUE
	INVERT = C.GL_INVERT
	KEEP = C.GL_KEEP
	LEFT = C.GL_LEFT
	LEQUAL = C.GL_LEQUAL
	LESS = C.GL_LESS
	LINE = C.GL_LINE
	LINEAR = C.GL_LINEAR
	LINEAR_MIPMAP_LINEAR = C.GL_LINEAR_MIPMAP_LINEAR
	LINEAR_MIPMAP_NEAREST = C.GL_LINEAR_MIPMAP_NEAREST
	LINES = C.GL_LINES
	LINE_LOOP = C.GL_LINE_LOOP
	LINE_SMOOTH = C.GL_LINE_SMOOTH
	LINE_SMOOTH_HINT = C.GL_LINE_SMOOTH_HINT
	LINE_STRIP = C.GL_LINE_STRIP
	LINE_WIDTH = C.GL_LINE_WIDTH
	LINE_WIDTH_GRANULARITY = C.GL_LINE_WIDTH_GRANULARITY
	LINE_WIDTH_RANGE = C.GL_LINE_WIDTH_RANGE
	MAX_TEXTURE_SIZE = C.GL_MAX_TEXTURE_SIZE
	MAX_VIEWPORT_DIMS = C.GL_MAX_VIEWPORT_DIMS
	NAND = C.GL_NAND
	NEAREST = C.GL_NEAREST
	NEAREST_MIPMAP_LINEAR = C.GL_NEAREST_MIPMAP_LINEAR
	NEAREST_MIPMAP_NEAREST = C.GL_NEAREST_MIPMAP_NEAREST
	NEVER = C.GL_NEVER
	NICEST = C.GL_NICEST
	NONE = C.GL_NONE
	NOOP = C.GL_NOOP
	NOR = C.GL_NOR
	NOTEQUAL = C.GL_NOTEQUAL
	NO_ERROR = C.GL_NO_ERROR
	ONE = C.GL_ONE
	ONE_MINUS_DST_ALPHA = C.GL_ONE_MINUS_DST_ALPHA
	ONE_MINUS_DST_COLOR = C.GL_ONE_MINUS_DST_COLOR
	ONE_MINUS_SRC_ALPHA = C.GL_ONE_MINUS_SRC_ALPHA
	ONE_MINUS_SRC_COLOR = C.GL_ONE_MINUS_SRC_COLOR
	OR = C.GL_OR
	OR_INVERTED = C.GL_OR_INVERTED
	OR_REVERSE = C.GL_OR_REVERSE
	OUT_OF_MEMORY = C.GL_OUT_OF_MEMORY
	PACK_ALIGNMENT = C.GL_PACK_ALIGNMENT
	PACK_LSB_FIRST = C.GL_PACK_LSB_FIRST
	PACK_ROW_LENGTH = C.GL_PACK_ROW_LENGTH
	PACK_SKIP_PIXELS = C.GL_PACK_SKIP_PIXELS
	PACK_SKIP_ROWS = C.GL_PACK_SKIP_ROWS
	PACK_SWAP_BYTES = C.GL_PACK_SWAP_BYTES
	POINT = C.GL_POINT
	POINTS = C.GL_POINTS
	POINT_SIZE = C.GL_POINT_SIZE
	POINT_SIZE_GRANULARITY = C.GL_POINT_SIZE_GRANULARITY
	POINT_SIZE_RANGE = C.GL_POINT_SIZE_RANGE
	PROXY_TEXTURE_1D = C.GL_PROXY_TEXTURE_1D
	PROXY_TEXTURE_2D = C.GL_PROXY_TEXTURE_2D
	TRUE = C.GL_TRUE
	UNPACK_ALIGNMENT = C.GL_UNPACK_ALIGNMENT
	UNPACK_LSB_FIRST = C.GL_UNPACK_LSB_FIRST
	UNPACK_ROW_LENGTH = C.GL_UNPACK_ROW_LENGTH
	UNPACK_SKIP_PIXELS = C.GL_UNPACK_SKIP_PIXELS
	UNPACK_SKIP_ROWS = C.GL_UNPACK_SKIP_ROWS
	UNPACK_SWAP_BYTES = C.GL_UNPACK_SWAP_BYTES
	UNSIGNED_BYTE = C.GL_UNSIGNED_BYTE
	UNSIGNED_INT = C.GL_UNSIGNED_INT
	UNSIGNED_SHORT = C.GL_UNSIGNED_SHORT
	VENDOR = C.GL_VENDOR
	VERSION = C.GL_VERSION
	VIEWPORT = C.GL_VIEWPORT
	XOR = C.GL_XOR
	ZERO = C.GL_ZERO
)
// VERSION_1_2
const (
	ALIASED_LINE_WIDTH_RANGE = C.GL_ALIASED_LINE_WIDTH_RANGE
	BGR = C.GL_BGR
	BGRA = C.GL_BGRA
	MAX_3D_TEXTURE_SIZE = C.GL_MAX_3D_TEXTURE_SIZE
	MAX_ELEMENTS_INDICES = C.GL_MAX_ELEMENTS_INDICES
	MAX_ELEMENTS_VERTICES = C.GL_MAX_ELEMENTS_VERTICES
	PACK_IMAGE_HEIGHT = C.GL_PACK_IMAGE_HEIGHT
	PACK_SKIP_IMAGES = C.GL_PACK_SKIP_IMAGES
	PROXY_TEXTURE_3D = C.GL_PROXY_TEXTURE_3D
	UNPACK_IMAGE_HEIGHT = C.GL_UNPACK_IMAGE_HEIGHT
	UNPACK_SKIP_IMAGES = C.GL_UNPACK_SKIP_IMAGES
	UNSIGNED_BYTE_2_3_3_REV = C.GL_UNSIGNED_BYTE_2_3_3_REV
	UNSIGNED_BYTE_3_3_2 = C.GL_UNSIGNED_BYTE_3_3_2
	UNSIGNED_INT_10_10_10_2 = C.GL_UNSIGNED_INT_10_10_10_2
	UNSIGNED_INT_2_10_10_10_REV = C.GL_UNSIGNED_INT_2_10_10_10_REV
	UNSIGNED_INT_8_8_8_8 = C.GL_UNSIGNED_INT_8_8_8_8
	UNSIGNED_INT_8_8_8_8_REV = C.GL_UNSIGNED_INT_8_8_8_8_REV
	UNSIGNED_SHORT_1_5_5_5_REV = C.GL_UNSIGNED_SHORT_1_5_5_5_REV
	UNSIGNED_SHORT_4_4_4_4 = C.GL_UNSIGNED_SHORT_4_4_4_4
	UNSIGNED_SHORT_4_4_4_4_REV = C.GL_UNSIGNED_SHORT_4_4_4_4_REV
	UNSIGNED_SHORT_5_5_5_1 = C.GL_UNSIGNED_SHORT_5_5_5_1
	UNSIGNED_SHORT_5_6_5 = C.GL_UNSIGNED_SHORT_5_6_5
	UNSIGNED_SHORT_5_6_5_REV = C.GL_UNSIGNED_SHORT_5_6_5_REV
)
// VERSION_1_3
const (
	ACTIVE_TEXTURE = C.GL_ACTIVE_TEXTURE
	COMPRESSED_RGB = C.GL_COMPRESSED_RGB
	COMPRESSED_RGBA = C.GL_COMPRESSED_RGBA
	COMPRESSED_TEXTURE_FORMATS = C.GL_COMPRESSED_TEXTURE_FORMATS
	MAX_CUBE_MAP_TEXTURE_SIZE = C.GL_MAX_CUBE_MAP_TEXTURE_SIZE
	NUM_COMPRESSED_TEXTURE_FORMATS = C.GL_NUM_COMPRESSED_TEXTURE_FORMATS
	PROXY_TEXTURE_CUBE_MAP = C.GL_PROXY_TEXTURE_CUBE_MAP
)
// VERSION_1_4
const (
	BLEND_DST_ALPHA = C.GL_BLEND_DST_ALPHA
	BLEND_DST_RGB = C.GL_BLEND_DST_RGB
	BLEND_SRC_ALPHA = C.GL_BLEND_SRC_ALPHA
	BLEND_SRC_RGB = C.GL_BLEND_SRC_RGB
	DECR_WRAP = C.GL_DECR_WRAP
	DEPTH_COMPONENT16 = C.GL_DEPTH_COMPONENT16
	DEPTH_COMPONENT24 = C.GL_DEPTH_COMPONENT24
	DEPTH_COMPONENT32 = C.GL_DEPTH_COMPONENT32
	INCR_WRAP = C.GL_INCR_WRAP
	MAX_TEXTURE_LOD_BIAS = C.GL_MAX_TEXTURE_LOD_BIAS
	MIRRORED_REPEAT = C.GL_MIRRORED_REPEAT
	POINT_FADE_THRESHOLD_SIZE = C.GL_POINT_FADE_THRESHOLD_SIZE
)
// VERSION_1_5
const (
	ARRAY_BUFFER = C.GL_ARRAY_BUFFER
	ARRAY_BUFFER_BINDING = C.GL_ARRAY_BUFFER_BINDING
	BUFFER_ACCESS = C.GL_BUFFER_ACCESS
	BUFFER_MAPPED = C.GL_BUFFER_MAPPED
	BUFFER_MAP_POINTER = C.GL_BUFFER_MAP_POINTER
	BUFFER_SIZE = C.GL_BUFFER_SIZE
	BUFFER_USAGE = C.GL_BUFFER_USAGE
	CURRENT_QUERY = C.GL_CURRENT_QUERY
	DYNAMIC_COPY = C.GL_DYNAMIC_COPY
	DYNAMIC_DRAW = C.GL_DYNAMIC_DRAW
	DYNAMIC_READ = C.GL_DYNAMIC_READ
	ELEMENT_ARRAY_BUFFER = C.GL_ELEMENT_ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER_BINDING = C.GL_ELEMENT_ARRAY_BUFFER_BINDING
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = C.GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING
	WRITE_ONLY = C.GL_WRITE_ONLY
)
// VERSION_2_0
const (
	ACTIVE_ATTRIBUTES = C.GL_ACTIVE_ATTRIBUTES
	ACTIVE_ATTRIBUTE_MAX_LENGTH = C.GL_ACTIVE_ATTRIBUTE_MAX_LENGTH
	ACTIVE_UNIFORMS = C.GL_ACTIVE_UNIFORMS
	ACTIVE_UNIFORM_MAX_LENGTH = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	ATTACHED_SHADERS = C.GL_ATTACHED_SHADERS
	BLEND_EQUATION_ALPHA = C.GL_BLEND_EQUATION_ALPHA
	BLEND_EQUATION_RGB = C.GL_BLEND_EQUATION_RGB
	BOOL = C.GL_BOOL
	BOOL_VEC2 = C.GL_BOOL_VEC2
	BOOL_VEC3 = C.GL_BOOL_VEC3
	BOOL_VEC4 = C.GL_BOOL_VEC4
	CURRENT_PROGRAM = C.GL_CURRENT_PROGRAM
	CURRENT_VERTEX_ATTRIB = C.GL_CURRENT_VERTEX_ATTRIB
	DELETE_STATUS = C.GL_DELETE_STATUS
	DRAW_BUFFER0 = C.GL_DRAW_BUFFER0
	DRAW_BUFFER1 = C.GL_DRAW_BUFFER1
	DRAW_BUFFER10 = C.GL_DRAW_BUFFER10
	DRAW_BUFFER11 = C.GL_DRAW_BUFFER11
	DRAW_BUFFER12 = C.GL_DRAW_BUFFER12
	DRAW_BUFFER13 = C.GL_DRAW_BUFFER13
	DRAW_BUFFER14 = C.GL_DRAW_BUFFER14
	DRAW_BUFFER15 = C.GL_DRAW_BUFFER15
	DRAW_BUFFER2 = C.GL_DRAW_BUFFER2
	DRAW_BUFFER3 = C.GL_DRAW_BUFFER3
	DRAW_BUFFER4 = C.GL_DRAW_BUFFER4
	DRAW_BUFFER5 = C.GL_DRAW_BUFFER5
	DRAW_BUFFER6 = C.GL_DRAW_BUFFER6
	DRAW_BUFFER7 = C.GL_DRAW_BUFFER7
	DRAW_BUFFER8 = C.GL_DRAW_BUFFER8
	DRAW_BUFFER9 = C.GL_DRAW_BUFFER9
	FLOAT_MAT2 = C.GL_FLOAT_MAT2
	FLOAT_MAT3 = C.GL_FLOAT_MAT3
	FLOAT_MAT4 = C.GL_FLOAT_MAT4
	FLOAT_VEC2 = C.GL_FLOAT_VEC2
	FLOAT_VEC3 = C.GL_FLOAT_VEC3
	FLOAT_VEC4 = C.GL_FLOAT_VEC4
	FRAGMENT_SHADER = C.GL_FRAGMENT_SHADER
	FRAGMENT_SHADER_DERIVATIVE_HINT = C.GL_FRAGMENT_SHADER_DERIVATIVE_HINT
	INFO_LOG_LENGTH = C.GL_INFO_LOG_LENGTH
	INT_VEC2 = C.GL_INT_VEC2
	INT_VEC3 = C.GL_INT_VEC3
	INT_VEC4 = C.GL_INT_VEC4
	LINK_STATUS = C.GL_LINK_STATUS
	LOWER_LEFT = C.GL_LOWER_LEFT
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = C.GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS
	MAX_DRAW_BUFFERS = C.GL_MAX_DRAW_BUFFERS
	MAX_FRAGMENT_UNIFORM_COMPONENTS = C.GL_MAX_FRAGMENT_UNIFORM_COMPONENTS
	MAX_TEXTURE_IMAGE_UNITS = C.GL_MAX_TEXTURE_IMAGE_UNITS
	MAX_VARYING_FLOATS = C.GL_MAX_VARYING_FLOATS
	MAX_VERTEX_ATTRIBS = C.GL_MAX_VERTEX_ATTRIBS
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = C.GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS
	MAX_VERTEX_UNIFORM_COMPONENTS = C.GL_MAX_VERTEX_UNIFORM_COMPONENTS
	UPPER_LEFT = C.GL_UPPER_LEFT
	VALIDATE_STATUS = C.GL_VALIDATE_STATUS
	VERTEX_ATTRIB_ARRAY_ENABLED = C.GL_VERTEX_ATTRIB_ARRAY_ENABLED
	VERTEX_ATTRIB_ARRAY_NORMALIZED = C.GL_VERTEX_ATTRIB_ARRAY_NORMALIZED
	VERTEX_ATTRIB_ARRAY_POINTER = C.GL_VERTEX_ATTRIB_ARRAY_POINTER
	VERTEX_ATTRIB_ARRAY_SIZE = C.GL_VERTEX_ATTRIB_ARRAY_SIZE
	VERTEX_ATTRIB_ARRAY_STRIDE = C.GL_VERTEX_ATTRIB_ARRAY_STRIDE
	VERTEX_ATTRIB_ARRAY_TYPE = C.GL_VERTEX_ATTRIB_ARRAY_TYPE
	VERTEX_PROGRAM_POINT_SIZE = C.GL_VERTEX_PROGRAM_POINT_SIZE
	VERTEX_SHADER = C.GL_VERTEX_SHADER
)
// VERSION_2_1
const (
	COMPRESSED_SRGB = C.GL_COMPRESSED_SRGB
	COMPRESSED_SRGB_ALPHA = C.GL_COMPRESSED_SRGB_ALPHA
	FLOAT_MAT2x3 = C.GL_FLOAT_MAT2x3
	FLOAT_MAT2x4 = C.GL_FLOAT_MAT2x4
	FLOAT_MAT3x2 = C.GL_FLOAT_MAT3x2
	FLOAT_MAT3x4 = C.GL_FLOAT_MAT3x4
	FLOAT_MAT4x2 = C.GL_FLOAT_MAT4x2
	FLOAT_MAT4x3 = C.GL_FLOAT_MAT4x3
	PIXEL_PACK_BUFFER = C.GL_PIXEL_PACK_BUFFER
	PIXEL_PACK_BUFFER_BINDING = C.GL_PIXEL_PACK_BUFFER_BINDING
	PIXEL_UNPACK_BUFFER = C.GL_PIXEL_UNPACK_BUFFER
	PIXEL_UNPACK_BUFFER_BINDING = C.GL_PIXEL_UNPACK_BUFFER_BINDING
)
// VERSION_3_0
const (
	BGRA_INTEGER = C.GL_BGRA_INTEGER
	BGR_INTEGER = C.GL_BGR_INTEGER
	BLUE_INTEGER = C.GL_BLUE_INTEGER
	BUFFER_ACCESS_FLAGS = C.GL_BUFFER_ACCESS_FLAGS
	BUFFER_MAP_LENGTH = C.GL_BUFFER_MAP_LENGTH
	BUFFER_MAP_OFFSET = C.GL_BUFFER_MAP_OFFSET
	CLIP_DISTANCE0 = C.GL_CLIP_DISTANCE0
	CLIP_DISTANCE1 = C.GL_CLIP_DISTANCE1
	CLIP_DISTANCE2 = C.GL_CLIP_DISTANCE2
	CLIP_DISTANCE3 = C.GL_CLIP_DISTANCE3
	CLIP_DISTANCE4 = C.GL_CLIP_DISTANCE4
	CLIP_DISTANCE5 = C.GL_CLIP_DISTANCE5
	CLIP_DISTANCE6 = C.GL_CLIP_DISTANCE6
	CLIP_DISTANCE7 = C.GL_CLIP_DISTANCE7
	COLOR_ATTACHMENT0 = C.GL_COLOR_ATTACHMENT0
	COLOR_ATTACHMENT1 = C.GL_COLOR_ATTACHMENT1
	COLOR_ATTACHMENT10 = C.GL_COLOR_ATTACHMENT10
	COLOR_ATTACHMENT11 = C.GL_COLOR_ATTACHMENT11
	COLOR_ATTACHMENT12 = C.GL_COLOR_ATTACHMENT12
	COLOR_ATTACHMENT13 = C.GL_COLOR_ATTACHMENT13
	COLOR_ATTACHMENT14 = C.GL_COLOR_ATTACHMENT14
	COLOR_ATTACHMENT15 = C.GL_COLOR_ATTACHMENT15
	COLOR_ATTACHMENT2 = C.GL_COLOR_ATTACHMENT2
	COLOR_ATTACHMENT3 = C.GL_COLOR_ATTACHMENT3
	COLOR_ATTACHMENT4 = C.GL_COLOR_ATTACHMENT4
	COLOR_ATTACHMENT5 = C.GL_COLOR_ATTACHMENT5
	COLOR_ATTACHMENT6 = C.GL_COLOR_ATTACHMENT6
	COLOR_ATTACHMENT7 = C.GL_COLOR_ATTACHMENT7
	COLOR_ATTACHMENT8 = C.GL_COLOR_ATTACHMENT8
	COLOR_ATTACHMENT9 = C.GL_COLOR_ATTACHMENT9
	COMPARE_REF_TO_TEXTURE = C.GL_COMPARE_REF_TO_TEXTURE
	COMPRESSED_RED = C.GL_COMPRESSED_RED
	COMPRESSED_RED_RGTC1 = C.GL_COMPRESSED_RED_RGTC1
	COMPRESSED_RG = C.GL_COMPRESSED_RG
	COMPRESSED_RG_RGTC2 = C.GL_COMPRESSED_RG_RGTC2
	COMPRESSED_SIGNED_RED_RGTC1 = C.GL_COMPRESSED_SIGNED_RED_RGTC1
	COMPRESSED_SIGNED_RG_RGTC2 = C.GL_COMPRESSED_SIGNED_RG_RGTC2
	CONTEXT_FLAGS = C.GL_CONTEXT_FLAGS
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT = C.GL_CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT
	DEPTH24_STENCIL8 = C.GL_DEPTH24_STENCIL8
	DEPTH32F_STENCIL8 = C.GL_DEPTH32F_STENCIL8
	DEPTH_ATTACHMENT = C.GL_DEPTH_ATTACHMENT
	DEPTH_COMPONENT32F = C.GL_DEPTH_COMPONENT32F
	DEPTH_STENCIL = C.GL_DEPTH_STENCIL
	DEPTH_STENCIL_ATTACHMENT = C.GL_DEPTH_STENCIL_ATTACHMENT
	DRAW_FRAMEBUFFER = C.GL_DRAW_FRAMEBUFFER
	DRAW_FRAMEBUFFER_BINDING = C.GL_DRAW_FRAMEBUFFER_BINDING
	FIXED_ONLY = C.GL_FIXED_ONLY
	FLOAT_32_UNSIGNED_INT_24_8_REV = C.GL_FLOAT_32_UNSIGNED_INT_24_8_REV
	FRAMEBUFFER = C.GL_FRAMEBUFFER
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING = C.GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE = C.GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE
	FRAMEBUFFER_ATTACHMENT_RED_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE = C.GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL
	FRAMEBUFFER_BINDING = C.GL_FRAMEBUFFER_BINDING
	FRAMEBUFFER_COMPLETE = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_DEFAULT = C.GL_FRAMEBUFFER_DEFAULT
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER = C.GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE = C.GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER = C.GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	FRAMEBUFFER_SRGB = C.GL_FRAMEBUFFER_SRGB
	FRAMEBUFFER_UNDEFINED = C.GL_FRAMEBUFFER_UNDEFINED
	FRAMEBUFFER_UNSUPPORTED = C.GL_FRAMEBUFFER_UNSUPPORTED
	GREEN_INTEGER = C.GL_GREEN_INTEGER
	HALF_FLOAT = C.GL_HALF_FLOAT
	INTERLEAVED_ATTRIBS = C.GL_INTERLEAVED_ATTRIBS
	INT_SAMPLER_1D = C.GL_INT_SAMPLER_1D
	INT_SAMPLER_1D_ARRAY = C.GL_INT_SAMPLER_1D_ARRAY
	INT_SAMPLER_2D = C.GL_INT_SAMPLER_2D
	INT_SAMPLER_2D_ARRAY = C.GL_INT_SAMPLER_2D_ARRAY
	INT_SAMPLER_3D = C.GL_INT_SAMPLER_3D
	INT_SAMPLER_CUBE = C.GL_INT_SAMPLER_CUBE
	INVALID_FRAMEBUFFER_OPERATION = C.GL_INVALID_FRAMEBUFFER_OPERATION
	MAJOR_VERSION = C.GL_MAJOR_VERSION
	MAP_FLUSH_EXPLICIT_BIT = C.GL_MAP_FLUSH_EXPLICIT_BIT
	MAP_INVALIDATE_BUFFER_BIT = C.GL_MAP_INVALIDATE_BUFFER_BIT
	MAP_INVALIDATE_RANGE_BIT = C.GL_MAP_INVALIDATE_RANGE_BIT
	MAP_READ_BIT = C.GL_MAP_READ_BIT
	MAP_UNSYNCHRONIZED_BIT = C.GL_MAP_UNSYNCHRONIZED_BIT
	MAP_WRITE_BIT = C.GL_MAP_WRITE_BIT
	MAX_ARRAY_TEXTURE_LAYERS = C.GL_MAX_ARRAY_TEXTURE_LAYERS
	MAX_CLIP_DISTANCES = C.GL_MAX_CLIP_DISTANCES
	MAX_COLOR_ATTACHMENTS = C.GL_MAX_COLOR_ATTACHMENTS
	MAX_PROGRAM_TEXEL_OFFSET = C.GL_MAX_PROGRAM_TEXEL_OFFSET
	MAX_RENDERBUFFER_SIZE = C.GL_MAX_RENDERBUFFER_SIZE
	MAX_SAMPLES = C.GL_MAX_SAMPLES
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = C.GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS = C.GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS = C.GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS
	MAX_VARYING_COMPONENTS = C.GL_MAX_VARYING_COMPONENTS
	MINOR_VERSION = C.GL_MINOR_VERSION
	MIN_PROGRAM_TEXEL_OFFSET = C.GL_MIN_PROGRAM_TEXEL_OFFSET
	NUM_EXTENSIONS = C.GL_NUM_EXTENSIONS
	PRIMITIVES_GENERATED = C.GL_PRIMITIVES_GENERATED
	PROXY_TEXTURE_1D_ARRAY = C.GL_PROXY_TEXTURE_1D_ARRAY
	PROXY_TEXTURE_2D_ARRAY = C.GL_PROXY_TEXTURE_2D_ARRAY
	UNSIGNED_INT_10F_11F_11F_REV = C.GL_UNSIGNED_INT_10F_11F_11F_REV
	UNSIGNED_INT_24_8 = C.GL_UNSIGNED_INT_24_8
	UNSIGNED_INT_5_9_9_9_REV = C.GL_UNSIGNED_INT_5_9_9_9_REV
	UNSIGNED_INT_SAMPLER_1D = C.GL_UNSIGNED_INT_SAMPLER_1D
	UNSIGNED_INT_SAMPLER_1D_ARRAY = C.GL_UNSIGNED_INT_SAMPLER_1D_ARRAY
	UNSIGNED_INT_SAMPLER_2D = C.GL_UNSIGNED_INT_SAMPLER_2D
	UNSIGNED_INT_SAMPLER_2D_ARRAY = C.GL_UNSIGNED_INT_SAMPLER_2D_ARRAY
	UNSIGNED_INT_SAMPLER_3D = C.GL_UNSIGNED_INT_SAMPLER_3D
	UNSIGNED_INT_SAMPLER_CUBE = C.GL_UNSIGNED_INT_SAMPLER_CUBE
	UNSIGNED_INT_VEC2 = C.GL_UNSIGNED_INT_VEC2
	UNSIGNED_INT_VEC3 = C.GL_UNSIGNED_INT_VEC3
	UNSIGNED_INT_VEC4 = C.GL_UNSIGNED_INT_VEC4
	UNSIGNED_NORMALIZED = C.GL_UNSIGNED_NORMALIZED
	VERTEX_ATTRIB_ARRAY_INTEGER = C.GL_VERTEX_ATTRIB_ARRAY_INTEGER
)
// VERSION_3_1
const (
	ACTIVE_UNIFORM_BLOCKS = C.GL_ACTIVE_UNIFORM_BLOCKS
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH = C.GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH
	COPY_READ_BUFFER = C.GL_COPY_READ_BUFFER
	COPY_WRITE_BUFFER = C.GL_COPY_WRITE_BUFFER
	INT_SAMPLER_2D_RECT = C.GL_INT_SAMPLER_2D_RECT
	INT_SAMPLER_BUFFER = C.GL_INT_SAMPLER_BUFFER
	INVALID_INDEX = C.GL_INVALID_INDEX
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS = C.GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS
	MAX_COMBINED_UNIFORM_BLOCKS = C.GL_MAX_COMBINED_UNIFORM_BLOCKS
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS = C.GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS
	MAX_FRAGMENT_UNIFORM_BLOCKS = C.GL_MAX_FRAGMENT_UNIFORM_BLOCKS
	MAX_RECTANGLE_TEXTURE_SIZE = C.GL_MAX_RECTANGLE_TEXTURE_SIZE
	MAX_TEXTURE_BUFFER_SIZE = C.GL_MAX_TEXTURE_BUFFER_SIZE
	MAX_UNIFORM_BLOCK_SIZE = C.GL_MAX_UNIFORM_BLOCK_SIZE
	MAX_UNIFORM_BUFFER_BINDINGS = C.GL_MAX_UNIFORM_BUFFER_BINDINGS
	MAX_VERTEX_UNIFORM_BLOCKS = C.GL_MAX_VERTEX_UNIFORM_BLOCKS
	PRIMITIVE_RESTART = C.GL_PRIMITIVE_RESTART
	PRIMITIVE_RESTART_INDEX = C.GL_PRIMITIVE_RESTART_INDEX
	PROXY_TEXTURE_RECTANGLE = C.GL_PROXY_TEXTURE_RECTANGLE
	UNIFORM_ARRAY_STRIDE = C.GL_UNIFORM_ARRAY_STRIDE
	UNIFORM_BLOCK_ACTIVE_UNIFORMS = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES
	UNIFORM_BLOCK_BINDING = C.GL_UNIFORM_BLOCK_BINDING
	UNIFORM_BLOCK_DATA_SIZE = C.GL_UNIFORM_BLOCK_DATA_SIZE
	UNIFORM_BLOCK_INDEX = C.GL_UNIFORM_BLOCK_INDEX
	UNIFORM_BLOCK_NAME_LENGTH = C.GL_UNIFORM_BLOCK_NAME_LENGTH
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = C.GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER = C.GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER
	UNIFORM_BUFFER = C.GL_UNIFORM_BUFFER
	UNIFORM_BUFFER_BINDING = C.GL_UNIFORM_BUFFER_BINDING
	UNIFORM_BUFFER_OFFSET_ALIGNMENT = C.GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT
	UNIFORM_BUFFER_SIZE = C.GL_UNIFORM_BUFFER_SIZE
	UNIFORM_BUFFER_START = C.GL_UNIFORM_BUFFER_START
	UNIFORM_IS_ROW_MAJOR = C.GL_UNIFORM_IS_ROW_MAJOR
	UNIFORM_MATRIX_STRIDE = C.GL_UNIFORM_MATRIX_STRIDE
	UNIFORM_NAME_LENGTH = C.GL_UNIFORM_NAME_LENGTH
	UNIFORM_OFFSET = C.GL_UNIFORM_OFFSET
	UNIFORM_SIZE = C.GL_UNIFORM_SIZE
	UNIFORM_TYPE = C.GL_UNIFORM_TYPE
	UNSIGNED_INT_SAMPLER_2D_RECT = C.GL_UNSIGNED_INT_SAMPLER_2D_RECT
	UNSIGNED_INT_SAMPLER_BUFFER = C.GL_UNSIGNED_INT_SAMPLER_BUFFER
)
// VERSION_3_2
const (
	ALREADY_SIGNALED = C.GL_ALREADY_SIGNALED
	CONDITION_SATISFIED = C.GL_CONDITION_SATISFIED
	CONTEXT_COMPATIBILITY_PROFILE_BIT = C.GL_CONTEXT_COMPATIBILITY_PROFILE_BIT
	CONTEXT_CORE_PROFILE_BIT = C.GL_CONTEXT_CORE_PROFILE_BIT
	CONTEXT_PROFILE_MASK = C.GL_CONTEXT_PROFILE_MASK
	DEPTH_CLAMP = C.GL_DEPTH_CLAMP
	FIRST_VERTEX_CONVENTION = C.GL_FIRST_VERTEX_CONVENTION
	FRAMEBUFFER_ATTACHMENT_LAYERED = C.GL_FRAMEBUFFER_ATTACHMENT_LAYERED
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS = C.GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS
	GEOMETRY_INPUT_TYPE = C.GL_GEOMETRY_INPUT_TYPE
	GEOMETRY_OUTPUT_TYPE = C.GL_GEOMETRY_OUTPUT_TYPE
	GEOMETRY_SHADER = C.GL_GEOMETRY_SHADER
	GEOMETRY_VERTICES_OUT = C.GL_GEOMETRY_VERTICES_OUT
	INT_SAMPLER_2D_MULTISAMPLE = C.GL_INT_SAMPLER_2D_MULTISAMPLE
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY = C.GL_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	LAST_VERTEX_CONVENTION = C.GL_LAST_VERTEX_CONVENTION
	LINES_ADJACENCY = C.GL_LINES_ADJACENCY
	LINE_STRIP_ADJACENCY = C.GL_LINE_STRIP_ADJACENCY
	MAX_COLOR_TEXTURE_SAMPLES = C.GL_MAX_COLOR_TEXTURE_SAMPLES
	MAX_DEPTH_TEXTURE_SAMPLES = C.GL_MAX_DEPTH_TEXTURE_SAMPLES
	MAX_FRAGMENT_INPUT_COMPONENTS = C.GL_MAX_FRAGMENT_INPUT_COMPONENTS
	MAX_GEOMETRY_INPUT_COMPONENTS = C.GL_MAX_GEOMETRY_INPUT_COMPONENTS
	MAX_GEOMETRY_OUTPUT_COMPONENTS = C.GL_MAX_GEOMETRY_OUTPUT_COMPONENTS
	MAX_GEOMETRY_OUTPUT_VERTICES = C.GL_MAX_GEOMETRY_OUTPUT_VERTICES
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS = C.GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS = C.GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS
	MAX_GEOMETRY_UNIFORM_COMPONENTS = C.GL_MAX_GEOMETRY_UNIFORM_COMPONENTS
	MAX_INTEGER_SAMPLES = C.GL_MAX_INTEGER_SAMPLES
	MAX_SAMPLE_MASK_WORDS = C.GL_MAX_SAMPLE_MASK_WORDS
	MAX_SERVER_WAIT_TIMEOUT = C.GL_MAX_SERVER_WAIT_TIMEOUT
	MAX_VERTEX_OUTPUT_COMPONENTS = C.GL_MAX_VERTEX_OUTPUT_COMPONENTS
	OBJECT_TYPE = C.GL_OBJECT_TYPE
	PROGRAM_POINT_SIZE = C.GL_PROGRAM_POINT_SIZE
	PROVOKING_VERTEX = C.GL_PROVOKING_VERTEX
	PROXY_TEXTURE_2D_MULTISAMPLE = C.GL_PROXY_TEXTURE_2D_MULTISAMPLE
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY = C.GL_PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY
	UNSIGNALED = C.GL_UNSIGNALED
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	WAIT_FAILED = C.GL_WAIT_FAILED
)

// VERSION_1_0

// https://www.opengl.org/sdk/docs/man3/xhtml/glCullFace.xml
func CullFace(mode Enum)  {
	C.glCullFace((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFrontFace.xml
func FrontFace(mode Enum)  {
	C.glFrontFace((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glHint.xml
func Hint(target Enum, mode Enum)  {
	C.glHint((C.GLenum)(target), (C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glLineWidth.xml
func LineWidth(width Float)  {
	C.glLineWidth((C.GLfloat)(width))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPointSize.xml
func PointSize(size Float)  {
	C.glPointSize((C.GLfloat)(size))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum)  {
	C.glPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.glScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.glTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.glTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.glTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.glTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum)  {
	C.glDrawBuffer((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
func Clear(mask Bitfield)  {
	C.glClear((C.GLbitfield)(mask))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float)  {
	C.glClearColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.glClearStencil((C.GLint)(s))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearDepth.xml
func ClearDepth(depth Double)  {
	C.glClearDepth((C.GLclampd)(depth))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilMask.xml
func StencilMask(mask Uint)  {
	C.glStencilMask((C.GLuint)(mask))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.glColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDepthMask.xml
func DepthMask(flag Boolean)  {
	C.glDepthMask((C.GLboolean)(flag))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDisable.xml
func Disable(cap Enum)  {
	C.glDisable((C.GLenum)(cap))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEnable.xml
func Enable(cap Enum)  {
	C.glEnable((C.GLenum)(cap))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFinish.xml
func Finish()  {
	C.glFinish()
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFlush.xml
func Flush()  {
	C.glFlush()
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.glBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glLogicOp.xml
func LogicOp(opcode Enum)  {
	C.glLogicOp((C.GLenum)(opcode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.glStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.glStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum)  {
	C.glDepthFunc((C.GLenum)(func_))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float)  {
	C.glPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int)  {
	C.glPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum)  {
	C.glReadBuffer((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.glReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean)  {
	C.glGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double)  {
	C.glGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.glGetError())
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float)  {
	C.glGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int)  {
	C.glGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.glGetString((C.GLenum)(name)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.glGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.glGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.glGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.glGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.glGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.glIsEnabled((C.GLenum)(cap)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double)  {
	C.glDepthRange((C.GLclampd)(near_), (C.GLclampd)(far_))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.glViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_1

// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei)  {
	C.glDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Pointer)  {
	C.glDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float)  {
	C.glPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)  {
	C.glCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)  {
	C.glCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)  {
	C.glCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.glCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint)  {
	C.glBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint)  {
	C.glDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint)  {
	C.glGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.glIsTexture((C.GLuint)(texture)))
}
// VERSION_1_2

// https://www.opengl.org/sdk/docs/man3/xhtml/glBlendColor.xml
func BlendColor(red Float, green Float, blue Float, alpha Float)  {
	C.glBlendColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum)  {
	C.glBlendEquation((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.glDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.glTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.glCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_3

// https://www.opengl.org/sdk/docs/man3/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum)  {
	C.glActiveTexture((C.GLenum)(texture))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glSampleCoverage.xml
func SampleCoverage(value Float, invert Boolean)  {
	C.glSampleCoverage((C.GLclampf)(value), (C.GLboolean)(invert))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.glCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.glCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.glCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.glCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.glCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.glCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Pointer)  {
	C.glGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}
// VERSION_1_4

// https://www.opengl.org/sdk/docs/man3/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.glBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei)  {
	C.glMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(drawcount))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, drawcount Sizei)  {
	C.glMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float)  {
	C.glPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float)  {
	C.glPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int)  {
	C.glPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int)  {
	C.glPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_1_5

// https://www.opengl.org/sdk/docs/man3/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint)  {
	C.glGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint)  {
	C.glDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.glIsQuery((C.GLuint)(id)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint)  {
	C.glBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEndQuery.xml
func EndQuery(target Enum)  {
	C.glEndQuery((C.GLenum)(target))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int)  {
	C.glGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int)  {
	C.glGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)  {
	C.glGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint)  {
	C.glBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint)  {
	C.glDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint)  {
	C.glGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.glIsBuffer((C.GLuint)(buffer)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Pointer, usage Enum)  {
	C.glBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.glBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.glGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) Pointer {
	return (Pointer)(C.glMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.glUnmapBuffer((C.GLenum)(target)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.glGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Pointer)  {
	C.glGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// VERSION_2_0

// https://www.opengl.org/sdk/docs/man3/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)  {
	C.glBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum)  {
	C.glDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.glStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)  {
	C.glStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint)  {
	C.glStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint)  {
	C.glAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char)  {
	C.glBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCompileShader.xml
func CompileShader(shader Uint)  {
	C.glCompileShader((C.GLuint)(shader))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.glCreateProgram())
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.glCreateShader((C.GLenum)(type_)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint)  {
	C.glDeleteProgram((C.GLuint)(program))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint)  {
	C.glDeleteShader((C.GLuint)(shader))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint)  {
	C.glDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint)  {
	C.glDisableVertexAttribArray((C.GLuint)(index))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint)  {
	C.glEnableVertexAttribArray((C.GLuint)(index))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.glGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.glGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)  {
	C.glGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.glGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int)  {
	C.glGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.glGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int)  {
	C.glGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.glGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)  {
	C.glGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.glGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float)  {
	C.glGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int)  {
	C.glGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double)  {
	C.glGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float)  {
	C.glGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int)  {
	C.glGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Pointer)  {
	C.glGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.glIsProgram((C.GLuint)(program)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.glIsShader((C.GLuint)(shader)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
func LinkProgram(program Uint)  {
	C.glLinkProgram((C.GLuint)(program))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)  {
	C.glShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func UseProgram(program Uint)  {
	C.glUseProgram((C.GLuint)(program))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float)  {
	C.glUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float)  {
	C.glUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)  {
	C.glUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.glUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int)  {
	C.glUniform1i((C.GLint)(location), (C.GLint)(v0))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int)  {
	C.glUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)  {
	C.glUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.glUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float)  {
	C.glUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float)  {
	C.glUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float)  {
	C.glUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float)  {
	C.glUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int)  {
	C.glUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int)  {
	C.glUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int)  {
	C.glUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int)  {
	C.glUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint)  {
	C.glValidateProgram((C.GLuint)(program))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.glVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_2_1

// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.glUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// VERSION_3_0

// https://www.opengl.org/sdk/docs/man3/xhtml/glColorMaski.xml
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean)  {
	C.glColorMaski((C.GLuint)(index), (C.GLboolean)(r), (C.GLboolean)(g), (C.GLboolean)(b), (C.GLboolean)(a))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBooleani_v.xml
func GetBooleani_v(target Enum, index Uint, data *Boolean)  {
	C.glGetBooleani_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetIntegeri_v.xml
func GetIntegeri_v(target Enum, index Uint, data *Int)  {
	C.glGetIntegeri_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEnablei.xml
func Enablei(target Enum, index Uint)  {
	C.glEnablei((C.GLenum)(target), (C.GLuint)(index))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDisablei.xml
func Disablei(target Enum, index Uint)  {
	C.glDisablei((C.GLenum)(target), (C.GLuint)(index))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsEnabledi.xml
func IsEnabledi(target Enum, index Uint) Boolean {
	return (Boolean)(C.glIsEnabledi((C.GLenum)(target), (C.GLuint)(index)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBeginTransformFeedback.xml
func BeginTransformFeedback(primitiveMode Enum)  {
	C.glBeginTransformFeedback((C.GLenum)(primitiveMode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEndTransformFeedback.xml
func EndTransformFeedback()  {
	C.glEndTransformFeedback()
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindBufferRange.xml
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr)  {
	C.glBindBufferRange((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindBufferBase.xml
func BindBufferBase(target Enum, index Uint, buffer Uint)  {
	C.glBindBufferBase((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTransformFeedbackVaryings.xml
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum)  {
	C.glTransformFeedbackVaryings((C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTransformFeedbackVarying.xml
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char)  {
	C.glGetTransformFeedbackVarying((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClampColor.xml
func ClampColor(target Enum, clamp Enum)  {
	C.glClampColor((C.GLenum)(target), (C.GLenum)(clamp))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBeginConditionalRender.xml
func BeginConditionalRender(id Uint, mode Enum)  {
	C.glBeginConditionalRender((C.GLuint)(id), (C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glEndConditionalRender.xml
func EndConditionalRender()  {
	C.glEndConditionalRender()
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribIPointer.xml
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.glVertexAttribIPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribIiv.xml
func GetVertexAttribIiv(index Uint, pname Enum, params *Int)  {
	C.glGetVertexAttribIiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribIuiv.xml
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint)  {
	C.glGetVertexAttribIuiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1i.xml
func VertexAttribI1i(index Uint, x Int)  {
	C.glVertexAttribI1i((C.GLuint)(index), (C.GLint)(x))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2i.xml
func VertexAttribI2i(index Uint, x Int, y Int)  {
	C.glVertexAttribI2i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3i.xml
func VertexAttribI3i(index Uint, x Int, y Int, z Int)  {
	C.glVertexAttribI3i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4i.xml
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int)  {
	C.glVertexAttribI4i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1ui.xml
func VertexAttribI1ui(index Uint, x Uint)  {
	C.glVertexAttribI1ui((C.GLuint)(index), (C.GLuint)(x))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2ui.xml
func VertexAttribI2ui(index Uint, x Uint, y Uint)  {
	C.glVertexAttribI2ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3ui.xml
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint)  {
	C.glVertexAttribI3ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4ui.xml
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.glVertexAttribI4ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1iv.xml
func VertexAttribI1iv(index Uint, v *Int)  {
	C.glVertexAttribI1iv((C.GLuint)(index), (*C.GLint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2iv.xml
func VertexAttribI2iv(index Uint, v *Int)  {
	C.glVertexAttribI2iv((C.GLuint)(index), (*C.GLint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3iv.xml
func VertexAttribI3iv(index Uint, v *Int)  {
	C.glVertexAttribI3iv((C.GLuint)(index), (*C.GLint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4iv.xml
func VertexAttribI4iv(index Uint, v *Int)  {
	C.glVertexAttribI4iv((C.GLuint)(index), (*C.GLint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1uiv.xml
func VertexAttribI1uiv(index Uint, v *Uint)  {
	C.glVertexAttribI1uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2uiv.xml
func VertexAttribI2uiv(index Uint, v *Uint)  {
	C.glVertexAttribI2uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3uiv.xml
func VertexAttribI3uiv(index Uint, v *Uint)  {
	C.glVertexAttribI3uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4uiv.xml
func VertexAttribI4uiv(index Uint, v *Uint)  {
	C.glVertexAttribI4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4bv.xml
func VertexAttribI4bv(index Uint, v *Byte)  {
	C.glVertexAttribI4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4sv.xml
func VertexAttribI4sv(index Uint, v *Short)  {
	C.glVertexAttribI4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4ubv.xml
func VertexAttribI4ubv(index Uint, v *Ubyte)  {
	C.glVertexAttribI4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4usv.xml
func VertexAttribI4usv(index Uint, v *Ushort)  {
	C.glVertexAttribI4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformuiv.xml
func GetUniformuiv(program Uint, location Int, params *Uint)  {
	C.glGetUniformuiv((C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocation.xml
func BindFragDataLocation(program Uint, color Uint, name *Char)  {
	C.glBindFragDataLocation((C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(name))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetFragDataLocation.xml
func GetFragDataLocation(program Uint, name *Char) Int {
	return (Int)(C.glGetFragDataLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1ui.xml
func Uniform1ui(location Int, v0 Uint)  {
	C.glUniform1ui((C.GLint)(location), (C.GLuint)(v0))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2ui.xml
func Uniform2ui(location Int, v0 Uint, v1 Uint)  {
	C.glUniform2ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3ui.xml
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint)  {
	C.glUniform3ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4ui.xml
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)  {
	C.glUniform4ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform1uiv.xml
func Uniform1uiv(location Int, count Sizei, value *Uint)  {
	C.glUniform1uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform2uiv.xml
func Uniform2uiv(location Int, count Sizei, value *Uint)  {
	C.glUniform2uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform3uiv.xml
func Uniform3uiv(location Int, count Sizei, value *Uint)  {
	C.glUniform3uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniform4uiv.xml
func Uniform4uiv(location Int, count Sizei, value *Uint)  {
	C.glUniform4uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIiv.xml
func TexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.glTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIuiv.xml
func TexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.glTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIiv.xml
func GetTexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.glGetTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIuiv.xml
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.glGetTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferiv.xml
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int)  {
	C.glClearBufferiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferuiv.xml
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint)  {
	C.glClearBufferuiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferfv.xml
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float)  {
	C.glClearBufferfv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(value))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferfi.xml
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int)  {
	C.glClearBufferfi((C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetStringi.xml
func GetStringi(name Enum, index Uint) *Ubyte {
	return (*Ubyte)(C.glGetStringi((C.GLenum)(name), (C.GLuint)(index)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsRenderbuffer.xml
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return (Boolean)(C.glIsRenderbuffer((C.GLuint)(renderbuffer)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindRenderbuffer.xml
func BindRenderbuffer(target Enum, renderbuffer Uint)  {
	C.glBindRenderbuffer((C.GLenum)(target), (C.GLuint)(renderbuffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteRenderbuffers.xml
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.glDeleteRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.glGenRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glRenderbufferStorage.xml
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei)  {
	C.glRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetRenderbufferParameteriv.xml
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.glGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsFramebuffer.xml
func IsFramebuffer(framebuffer Uint) Boolean {
	return (Boolean)(C.glIsFramebuffer((C.GLuint)(framebuffer)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func BindFramebuffer(target Enum, framebuffer Uint)  {
	C.glBindFramebuffer((C.GLenum)(target), (C.GLuint)(framebuffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteFramebuffers.xml
func DeleteFramebuffers(n Sizei, framebuffers *Uint)  {
	C.glDeleteFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n Sizei, framebuffers *Uint)  {
	C.glGenFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCheckFramebufferStatus.xml
func CheckFramebufferStatus(target Enum) Enum {
	return (Enum)(C.glCheckFramebufferStatus((C.GLenum)(target)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture1D.xml
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.glFramebufferTexture1D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture2D.xml
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.glFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture3D.xml
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int)  {
	C.glFramebufferTexture3D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferRenderbuffer.xml
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint)  {
	C.glFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetFramebufferAttachmentParameteriv.xml
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int)  {
	C.glGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenerateMipmap.xml
func GenerateMipmap(target Enum)  {
	C.glGenerateMipmap((C.GLenum)(target))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBlitFramebuffer.xml
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)  {
	C.glBlitFramebuffer((C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glRenderbufferStorageMultisample.xml
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.glRenderbufferStorageMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTextureLayer.xml
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int)  {
	C.glFramebufferTextureLayer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glMapBufferRange.xml
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) Pointer {
	return (Pointer)(C.glMapBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFlushMappedBufferRange.xml
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr)  {
	C.glFlushMappedBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glBindVertexArray.xml
func BindVertexArray(array Uint)  {
	C.glBindVertexArray((C.GLuint)(array))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteVertexArrays.xml
func DeleteVertexArrays(n Sizei, arrays *Uint)  {
	C.glDeleteVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n Sizei, arrays *Uint)  {
	C.glGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsVertexArray.xml
func IsVertexArray(array Uint) Boolean {
	return (Boolean)(C.glIsVertexArray((C.GLuint)(array)))
}
// VERSION_3_1

// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawArraysInstanced.xml
func DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei)  {
	C.glDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Pointer, instancecount Sizei)  {
	C.glDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexBuffer.xml
func TexBuffer(target Enum, internalformat Enum, buffer Uint)  {
	C.glTexBuffer((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glPrimitiveRestartIndex.xml
func PrimitiveRestartIndex(index Uint)  {
	C.glPrimitiveRestartIndex((C.GLuint)(index))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glCopyBufferSubData.xml
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr)  {
	C.glCopyBufferSubData((C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformIndices.xml
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint)  {
	C.glGetUniformIndices((C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(uniformIndices))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformsiv.xml
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int)  {
	C.glGetActiveUniformsiv((C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(uniformIndices), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformName.xml
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char)  {
	C.glGetActiveUniformName((C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformName))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformBlockIndex.xml
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint {
	return (Uint)(C.glGetUniformBlockIndex((C.GLuint)(program), (*C.GLchar)(uniformBlockName)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformBlockiv.xml
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int)  {
	C.glGetActiveUniformBlockiv((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformBlockName.xml
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char)  {
	C.glGetActiveUniformBlockName((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformBlockName))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glUniformBlockBinding.xml
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint)  {
	C.glUniformBlockBinding((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
// VERSION_3_2

// https://www.opengl.org/sdk/docs/man3/xhtml/glGetInteger64i_v.xml
func GetInteger64i_v(target Enum, index Uint, data *Int64)  {
	C.glGetInteger64i_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(data))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferParameteri64v.xml
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64)  {
	C.glGetBufferParameteri64v((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture.xml
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int)  {
	C.glFramebufferTexture((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.glDrawElementsBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawRangeElementsBaseVertex.xml
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.glDrawRangeElementsBaseVertex((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsInstancedBaseVertex.xml
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, instancecount Sizei, basevertex Int)  {
	C.glDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Pointer, drawcount Sizei, basevertex *Int)  {
	C.glMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount), (*C.GLint)(basevertex))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glProvokingVertex.xml
func ProvokingVertex(mode Enum)  {
	C.glProvokingVertex((C.GLenum)(mode))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glFenceSync.xml
func FenceSync(condition Enum, flags Bitfield) Sync {
	return (Sync)(C.glFenceSync((C.GLenum)(condition), (C.GLbitfield)(flags)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glIsSync.xml
func IsSync(sync Sync) Boolean {
	return (Boolean)(C.glIsSync((C.GLsync)(sync)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteSync.xml
func DeleteSync(sync Sync)  {
	C.glDeleteSync((C.GLsync)(sync))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glClientWaitSync.xml
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum {
	return (Enum)(C.glClientWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout)))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glWaitSync.xml
func WaitSync(sync Sync, flags Bitfield, timeout Uint64)  {
	C.glWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetInteger64v.xml
func GetInteger64v(pname Enum, params *Int64)  {
	C.glGetInteger64v((C.GLenum)(pname), (*C.GLint64)(params))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetSynciv.xml
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int)  {
	C.glGetSynciv((C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2DMultisample.xml
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean)  {
	C.glTexImage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage3DMultisample.xml
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean)  {
	C.glTexImage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glGetMultisamplefv.xml
func GetMultisamplefv(pname Enum, index Uint, val *Float)  {
	C.glGetMultisamplefv((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}
// https://www.opengl.org/sdk/docs/man3/xhtml/glSampleMaski.xml
func SampleMaski(index Uint, mask Bitfield)  {
	C.glSampleMaski((C.GLuint)(index), (C.GLbitfield)(mask))
}

//Go bool to GL boolean.
func GLBool(b bool) Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

// GL boolean to Go bool.
func GoBool(b Boolean) bool {
	return b == TRUE
}

// Go string to GL string.
func GLString(str string) *Char {
	return (*Char)(C.CString(str))
}

// Allocates a GL string.
func GLStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

// Frees GL string.
func GLStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

// GL string (GLchar*) to Go string.
func GoString(str *Char) string {
	return C.GoString((*C.char)(str))
}

// GL string (GLubyte*) to Go string.
func GoStringUb(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

// GL string (GLchar*) with length to Go string.
func GoStringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

// Converts a list of Go strings to a slice of GL strings.
// Usefull for ShaderSource().
func GLStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

// Free GL string slice allocated by GLStringArray().
func GLStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

// Add offset to a pointer. Usefull for VertexAttribPointer, TexCoordPointer, NormalPointer, ... 
func Offset(p Pointer, o uintptr) Pointer {
	return Pointer(uintptr(p) + o)
}

// EOF
