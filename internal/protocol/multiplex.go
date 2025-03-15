package protocol

import (
	"github.com/lemon-mint/hsq/internal/protocol/itrie"
)

type MultiplexBuffer struct {
}

type Multiplexer struct {
	slot *itrie.ITrie[MultiplexBuffer]
}
