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

	if err := event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	event.Publish(gpubsub, []byte("Test message!"))

	time.Sleep(time.Minute)
}
