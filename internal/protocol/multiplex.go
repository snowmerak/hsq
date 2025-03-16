package protocol

import (
	"bytes"
	"io"

	"github.com/lemon-mint/hsq/internal/protocol/itrie"
)

type MultiplexBuffer struct {
	Buffer       *bytes.Buffer
	CommonHeader *CommonHeader
}

type Multiplexer struct {
	slot *itrie.ITrie[MultiplexBuffer]
}

func NewMultiplexer() *Multiplexer {
	return &Multiplexer{slot: itrie.New[MultiplexBuffer]()}
}

func (m *Multiplexer) Append(commonHeader *CommonHeader, data []byte) (io.Reader, bool, error) {
	key1, key2 := commonHeader.Key()
	key := (uint64(key1) << 32) | uint64(key2)
	buffer := m.slot.Search(key)
	if buffer == nil {
		buffer = &MultiplexBuffer{
			Buffer:       bytes.NewBuffer(nil),
			CommonHeader: commonHeader,
		}
		buffer = m.slot.InsertIfNotExists(key, buffer)
	}
	buffer.Buffer.Write(data)

	switch commonHeader.MessageSequence {
	case commonHeader.MessageMaxSequence:
		m.slot.Delete(key)
		return buffer.Buffer, true, nil
	default:
		return nil, false, nil
	}
}
