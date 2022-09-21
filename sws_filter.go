package astiav

//#cgo pkg-config: libswscale
//#include <libswscale/swscale.h>
import "C"

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libswscale/swscale.h#L123
type SWSFilter struct {
	c *C.struct_SwsFilter
}
