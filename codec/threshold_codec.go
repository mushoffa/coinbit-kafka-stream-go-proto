package codec

import (
	"fmt"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"google.golang.org/protobuf/proto"
)

type ThresholdCodec struct {}

func (c *ThresholdCodec) Encode(value interface{}) ([]byte, error) {
	if _, threshold := value.(*pb.Threshold); !threshold {
		return nil, fmt.Errorf("Expected *pb.Threshold type, got %T", value)
	}

	return proto.Marshal(value.(*pb.Threshold))
}

func (c *ThresholdCodec) Decode(data []byte) (interface{}, error) {
	var threshold pb.Threshold
	return &threshold, proto.Unmarshal(data, &threshold)
}