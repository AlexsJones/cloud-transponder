package main

import (
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/cloud-transponder/events/pubsub"
)

func main() {

	//Create our GCP pubsub
	gpubsub := gcloud.NewPubSub()

	//Create the GCP Pubsub configuration
	gconfig := gcloud.NewPubSubConfiguration()

	gconfig.Topic = "cadium"
	gconfig.ConnectionString = "beamery-trials"
	gconfig.SubscriptionString = "cadium:sub"

	event.Connect(gpubsub, gconfig)

	event.Subscribe(gpubsub, func(arg2 event.IMessage) {

		log.Println("Received message..")
		arg2.Ack()
	})

}
