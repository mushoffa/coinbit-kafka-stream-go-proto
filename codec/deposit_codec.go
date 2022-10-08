package codec

import (
	"fmt"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"google.golang.org/protobuf/proto"
)

type DepositMoneyCodec struct {}

func (c *DepositMoneyCodec) Encode(value interface{}) ([]byte, error) {
	if _, deposit := value.(*pb.DepositMoney); !deposit {
		return nil, fmt.Errorf("Expected *pb.DepositMoney type, got %T", value)
	}

	return proto.Marshal(value.(*pb.DepositMoney))
}

func (c *DepositMoneyCodec) Decode(data []byte) (interface{}, error) {
	var deposit pb.DepositMoney
	return &deposit, proto.Unmarshal(data, &deposit)
}

type DepositMoneyHistoryCodec struct {}

func (c *DepositMoneyHistoryCodec) Encode(value interface{}) ([]byte, error) {
	if _, deposits := value.(*pb.DepositMoneyHistory); !deposits {
		return nil, fmt.Errorf("Expected *pb.DepositMoneyHistory type, got %T", value)
	}

	return proto.Marshal(value.(*pb.DepositMoneyHistory))
}

func (c *DepositMoneyHistoryCodec) Decode(data []byte) (interface{}, error) {
	var deposit pb.DepositMoneyHistory
	return &deposit, proto.Unmarshal(data, &deposit)
}