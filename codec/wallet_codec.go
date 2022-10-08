package codec

import (
	"fmt"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"google.golang.org/protobuf/proto"
)

type WalletCodec struct {}

func (c *WalletCodec) Encode(value interface{}) ([]byte, error) {
	if _, wallet := value.(*pb.Wallet); !wallet {
		return nil, fmt.Errorf("Expected *pb.Wallet type, got %T", value)
	}

	return proto.Marshal(value.(*pb.Wallet))
}

func (c *WalletCodec) Decode(data []byte) (interface{}, error) {
	var wallet pb.Wallet
	return &wallet, proto.Unmarshal(data, &wallet)
}