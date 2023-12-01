package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("CreateSensorData", createSensorData)
}

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type SensorData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

// createSensorData is triggered by a Pub/Sub message containing sensor data
func createSensorData(ctx context.Context, e event.Event) error {
	log.Printf("Event: %v", e)
	var m MessagePublishedData
	if err := e.DataAs(&m); err != nil {
		msg := fmt.Sprintf("error while decoding data: %s", err.Error())
		log.Print(msg)
		return fmt.Errorf(msg)
	}

	var sensorData SensorData
	if err := json.Unmarshal(m.Message.Data, &sensorData); err != nil {
		msg := fmt.Sprintf("error while unmarshalling data: %s", err.Error())
		log.Print(msg)
		return fmt.Errorf(msg)
	}

	log.Printf("Temperature: %v, Humidity: %v", sensorData.Temperature, sensorData.Humidity)
	return nil
}
