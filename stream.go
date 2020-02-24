package stream

import (
	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
)

type Stream interface {
	AddJob(ipld.Node) error
	//GetAll(func(*StreamBlock)) error
	//Subscribe(callback func(*ipld.Node))
	//Unsubscribe()
	StartWorker() error
}

type StreamBlock struct {
	StreamID string
	BlockID  cid.Cid
}


/*
 * Used to store streamblocks within a stream.
 */
type blocklist struct {

}

type StreamConfig struct {
	StreamID uint64
	Nsubstreams int
}

type SubStreamConfig struct {
	ID         uint64
	StreamMap  uint64 // 0010 means need only the second sub-stream
	StartIndex int    // only download blocks after StartIndex
}
