package main

import (
	"github.com/AJackTi/rabbitmq-go-example/consumer/consts"
	"github.com/AJackTi/rabbitmq-go-example/consumer/handlers"
	"github.com/AJackTi/rabbitmq-go-example/consumer/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")

	exampleQueue := utils.RMQConsumer{
		consts.EXAMPLE_QUEUE,
		connectionString,
		handlers.HandleExample,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()

	// Multiple listeners can be specified here

	<-forever
}
