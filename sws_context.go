package astiav

//#cgo pkg-config: libswscale
//#include <libswscale/swscale.h>
import "C"

// https://github.com/FFmpeg/FFmpeg/blob/n4.2.7/libswscale/swscale_internal.h#L280
type SWSContext struct {
	c *C.struct_SwsContext
}

func AllocSWSContext() *SWSContext {
	return &SWSContext{
		c: C.sws_alloc_context(),
	}
}

func (sc *SWSContext) Init(srcFilter, dstFilter *SWSFilter) int {
	return int(C.sws_init_context(sc.c, srcFilter.c, dstFilter.c))
}

func (sc *SWSContext) Free() {
	C.sws_freeContext(sc.c)
}

func (sc *SWSContext) SrcW() int {
	return int(sc.c.srcW)
}

func (sc *SWSContext) SetSrcW(w int) {
	sc.c.srcW = C.int(w)
}

func (sc *SWSContext) SrcH() int {
	return int(sc.c.srcH)
}

func (sc *SWSContext) SetSrcH(h int) {
	sc.c.srcH = C.int(h)
}

func (sc *SWSContext) SrcFormat() PixelFormat {
	return PixelFormat(sc.c.srcFormat)
}

func (sc *SWSContext) SetSrcFormat(srcFormat PixelFormat) {
	sc.c.srcFormat = C.enum_AVPixelFormat(srcFormat)
}

func (sc *SWSContext) DstW() int {
	return int(sc.c.dstW)
}

func (sc *SWSContext) SetDstW(w int) {
	sc.c.dstW = C.int(w)
}

func (sc *SWSContext) DstH() int {
	return int(sc.c.dstH)
}

func (sc *SWSContext) SetDstH(h int) {
	sc.c.dstH = C.int(h)
}

func (sc *SWSContext) DstFormat() PixelFormat {
	return PixelFormat(sc.c.dstFormat)
}

func (sc *SWSContext) SetDstFormat(dstFormat PixelFormat) {
	sc.c.dstFormat = C.enum_AVPixelFormat(dstFormat)
}

func (sc *SWSContext) Param() [2]int {
	return [2]int{int(sc.c.param[0]), int(sc.c.param[1])}
}

func (sc *SWSContext) SetParam(a, b int) {
	sc.c.param[0] = C.int(a)
	sc.c.param[1] = C.int(b)
}
