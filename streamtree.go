package stream

import (
    "context"
	"github.com/ipfs/go-blocks/blockservice"
    peer "github.com/libp2p/go-libp2p-core/peer"
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
func (st *StreamTree) StartWorker(ctx context.Context, conf *StreamConfig, pid peer.ID) error {
    //stream := getStream(conf.ID)
    //call stream.StartWorker
    return nil
}
