package stream

import (
    "context"
)

type Router interface {
	StartStream(context.Context, *Stream) error
	AddStreamBlock(context.Context, *StreamBlock) error
	SubscribeStream(context.Context, *StreamConfig) error
}
