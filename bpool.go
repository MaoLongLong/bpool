package bpool

type BytePoolCap struct {
	c    chan []byte
	w    int
	wcap int
}

func NewBytePoolCap(maxSize, width, capWidth int) *BytePoolCap {
	return &BytePoolCap{
		c:    make(chan []byte, maxSize),
		w:    width,
		wcap: capWidth,
	}
}

func (bp *BytePoolCap) Get() (b []byte) {
	select {
	case b = <-bp.c:
	default:
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}
	return
}

func (bp *BytePoolCap) Put(b []byte) {
	select {
	case bp.c <- b:
	default:
	}
}

func (bp *BytePoolCap) Width() int {
	return bp.w
}

func (bp *BytePoolCap) WidthCap() int {
	return bp.wcap
}
