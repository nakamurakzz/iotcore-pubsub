test:
	go test -v ./...

build:
	go build -o bin/go-pubsub-function pubsub/main.go

deploy-pubsub:
	cd pubsub; gcloud functions deploy go-pubsub-function --gen2 --runtime=go121 --region=asia-northeast1 --source=./ --entry-point=CreateSensorData --trigger-topic=sensors ; cd ..