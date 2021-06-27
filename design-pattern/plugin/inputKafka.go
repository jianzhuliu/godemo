package plugin

import (
	"fmt"
	"reflect"
	"design-pattern/msg"
	"design-pattern/kafka"
)

type KafkaInput struct {
	Base
	consumer kafka.Consumer
}

func (i *KafkaInput) Receive() *msg.Message {
	if i.Status() != Started {
		fmt.Println("Kafka input plugin is not running, input nothing.")
		return nil
	}
	
	records := i.consumer.Poll()

	return msg.Builder().
		WithHeaderItem("Content-Type", "application/json").
		WithBodyItems(records.Items).
		Builder()
}

func (i *KafkaInput) Start() {
	i.status = Started
	fmt.Println("Kafka input plugin started")
}

func (i *KafkaInput) Stop() {
	i.status = Stopped
	fmt.Println("Kafka input plugin stopped")
}

func (i *KafkaInput) Init() {
	i.consumer = &kafka.MockConsumer{}
}

func init() {
	RegisterInput("kafka", reflect.TypeOf(KafkaInput{}))
}
