package codec

import "github.com/vmihailenco/msgpack"

type SerializeType byte

const (
	MessagePack SerializeType = iota
)

type MessagePackCodec struct {
}

func (c MessagePackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c MessagePackCodec) Decode()
