package main

import (
	"log"
	"time"

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
	gconfig.SubscriptionString = "cadium-sub"

	if err := event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	if err := event.Subscribe(gpubsub, func(arg2 event.IMessage) {

		log.Printf("Received: %s", string(arg2.GetRaw()))
		arg2.Ack()
	}); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Minute)
}
