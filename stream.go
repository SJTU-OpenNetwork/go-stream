package stream

import (
	"github.com/ipfs/go-cid"
)

type Stream struct {
	ID          uint64  // hash value of the stream
	Rate        float32 // in Kbps
	Nsubstreams int     // number of substreams
}

type StreamBlock struct {
	StreamID uint64
	BlockID  cid.Cid
}

type StreamConfig struct {
	ID         uint64
	StreamMap  uint64 // 0010 means need only the second sub-stream
	StartIndex int    // only download blocks after StartIndex
}
