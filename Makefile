run:
	go run ./main.go

build:
	go build  -o dist/ .
	
dev:
	~/go/bin/air --build.cmd "make build" --build.bin "./dist/amqp-publisher" --build.exclude_dir "dist,manifests"

docker-build:
	docker build --platform linux/amd64 -t quay.io/mohamedf0/amqp-publisher .

docker-push:
	docker push quay.io/mohamedf0/amqp-publisher