package stream

import (
    "context"
    peer "github.com/libp2p/go-libp2p-core/peer"
)

type Router interface {
	StartStream(context.Context, *Stream) error
	AddStreamBlock(context.Context, *StreamBlock) error
	SubscribeStream(context.Context, *StreamConfig) error
	StartWorker(context.Context, *StreamConfig, peer.ID) error
}
