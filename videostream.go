package stream

import (
ipld "github.com/ipfs/go-ipld-format"
)

/*
 * VideoStream is one implementation of stream for videos.
 */
type VideoStream struct {
	ID          string  // hash value of the stream
	//Rate        float32 // in Kbps
	Nsubstreams int     // number of substreams
	blocklist   *blocklist	//Used to store blocks
}

/*
 * Add a video chunk into video stream.
 */
func (s *VideoStream) AddFile(nd ipld.Node) error {
	return nil
}

func NewVideoStream(c *StreamConfig) *VideoStream {
	return &VideoStream{
		ID: c.StreamID,
		Nsubstreams: c.Nsubstreams,
	}
}
