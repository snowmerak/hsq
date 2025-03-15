package protocol

import "math"

type MultiplexBuffer struct {
}

type Multiplexer struct {
	slot [math.MaxUint64]*MultiplexBuffer
}
