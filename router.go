package stream

type Router interface {
	StartStream(s Stream) error
	AddStreamBlock(b StreamBlock) error
	SubscribeStream(conf StreamConfig) error
}
