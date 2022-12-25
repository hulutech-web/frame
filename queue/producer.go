package queue

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"gitee.com/hulutech/frame/helpers/pb"
	"gitee.com/hulutech/frame/helpers/zone"
	message "gitee.com/hulutech/frame/queue/protocol_buffers"
)

type producer struct {
	topicName   string
	channelName string
	param       proto.Message
	retries     uint32
	delay       zone.Duration
}

func NewProducer(topicName string, channelName string, param proto.Message, retries uint32, delay zone.Duration) *producer {
	return &producer{
		topicName:   topicName,
		channelName: channelName,
		param:       param,
		retries:     retries,
		delay:       delay,
	}
}

func (p *producer) Push() error {
	// compress param
	paramPb, err := proto.Marshal(p.param)
	if err != nil {
		return err
	}

	// compress message
	return push(p.topicName, p.channelName, &message.Message{
		Hash:     "", // is empty when first push
		Param:    paramPb,
		Retries:  p.retries,
		PushedAt: ptypes.TimestampNow(),
		Delay:    ptypes.DurationProto(p.delay),
		Tried:    0,
	})
}

func push(topicName string, channelName string, msg *message.Message) error {
	// compress message
	messagePb, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return Queue().Push(topicName, channelName, *pb.DurationConvert(msg.Delay), messagePb)
}
