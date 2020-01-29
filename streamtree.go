package stream

import (
	"github.com/ipfs/go-blocks/blockservice"
)

type StreamTree struct {
	bsrv blockservice.BlockService
}

func (st *StreamTree) StartStream(s *Stream) error {
    return nil
}
func (st *StreamTree) AddStreamBlock(b *StreamBlock) error {
    return nil
}
func (st *StreamTree) SubscribeStream(conf *StreamConfig) error {
    return nil
}
