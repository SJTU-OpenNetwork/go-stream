package stream

import (
    "context"
	"github.com/ipfs/go-blocks/blockservice"
)

type StreamTree struct {
	bsrv blockservice.BlockService
}

func (st *StreamTree) StartStream(ctx context.Context, s *Stream) error {
    return nil
}
func (st *StreamTree) AddStreamBlock(ctx context.Context, b *StreamBlock) error {
    return nil
}
func (st *StreamTree) SubscribeStream(ctx context.Context, conf *StreamConfig) error {
    return nil
}
