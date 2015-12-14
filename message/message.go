package message

import (
	"encoding/json"
	"errors"
	
	"github.com/asiainfoLDP/datahub_commons/mq"
)

type Message struct {
	Type     string      `json:"type,omitempty"`
	Receiver string      `json:"receiver,omitempty"`
	Sender   string      `json:"sender,omitempty"`
	Time     string      `json:"time,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

// this one is to get a better performance
type Message2 struct {
	Type     string      `json:"type,omitempty"`
	Receiver string      `json:"receiver,omitempty"`
	Sender   string      `json:"sender,omitempty"`
	Time     string      `json:"time,omitempty"`
	Data     []byte      `json:"data,omitempty"`
}

func PushMessageToQueue(queue mq.MessageQueue, topic string, key []byte, message *Message) error {
	json_bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	
	return queue.SendAsyncMessage(topic, key, json_bytes)
}

func ParseMessage2(data []byte) (*Message2, error) {
	if data == nil {
		return nil, errors.New("data can't be nil")
	}

	message := &Message2{}
	err := json.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}

	return message, nil	
}