package message

import (
	"encoding/json"
	"errors"
	//"time"
	
	"github.com/asiainfoLDP/datahub_commons/mq"
)

// todo: maybe it is better to package the message push and save time in the header
// the following time is the event happen time

type Message struct {
	Type     string      `json:"type,omitempty"`
	Receiver string      `json:"receiver,omitempty"`
	Sender   string      `json:"sender,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	//Time     time.Time   `json:"time,omitempty"` 
		// this is the time to insert message into table.
		// if there are other times related to this messsage,
		// pls put them in Data field.
}

func PushMessageToQueue(queue mq.MessageQueue, topic string, key []byte, message *Message) error {
	json_bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	
	return queue.SendAsyncMessage(topic, key, json_bytes)
}

func ParseJsonMessage(msgData []byte) (*Message, error) {
	if msgData == nil {
		return nil, errors.New("message data can't be nil")
	}

	msg := &Message{}
	err := json.Unmarshal(msgData, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil	
}

//=====================================
// 
//=====================================

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	IsHTML  bool   `json:"ishtml"`
}

func ParseJsonEmail(msgData []byte) (*Email, error) {
	if msgData == nil {
		return nil, errors.New("message data can't be nil")
	}

	msg := &Email{}
	err := json.Unmarshal(msgData, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil	
}
