# cloud-transponder

[![Maintainability](https://api.codeclimate.com/v1/badges/08da69d625b9649a6b5e/maintainability)](https://codeclimate.com/github/AlexsJones/cloud-transponder/maintainability)


A simple abstraction for messaging systems

Currently supported:
- GCP pubsub
- (Coming soon) Kafka

## Publish example

```go
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
	gconfig.ConnectionString = "some-gcloud-project"

	if err := event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	event.Publish(gpubsub, []byte("Test message!"))

	time.Sleep(time.Minute)
}

```

## Subscribe 
```go
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
	gconfig.ConnectionString = "some-gcloud-project"
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

```
