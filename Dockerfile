FROM golang:1.24 AS build-stage

WORKDIR /amqp-publisher

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/ -a -installsuffix cgo .

FROM redhat/ubi9-minimal:9.6 AS serve-stage

COPY  --from=build-stage /dist/amqp-publisher /bin/

CMD ["amqp-publisher"]