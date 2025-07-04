package utils

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendAMQPMessage(amqpUri string, queue string, mandatory bool, immediate bool, contentType string, body []byte) error {
	conn, err := amqp.Dial(amqpUri)

	if err != nil {
		return errors.New(err.Error() + "\nFailed to connect to AMQP")
	}

	defer conn.Close()

	channel, err := conn.Channel()

	if err != nil {
		return errors.New(err.Error() + "\nFailed to get AMQP channel")
	}
	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx,
		"",        // exchange
		queue,     // routing key
		mandatory, // mandatory
		immediate, // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})

	if err != nil {
		return errors.New(err.Error() + "\nFailed to publish amqp message")
	}

	return nil
}

func ValidatePublishContentType(contentType string) error {
	validContentTypes := []string{"application/json", "text/plain"}
	if !slices.Contains(validContentTypes, contentType) {
		return errors.New("Error: content-type must be one of the following " + strings.Join(validContentTypes, ", ") + ".")
	}

	return nil
}

func ValidateJsonMessage(bytes []byte) error {
	parsedJson := map[string]interface{}{}

	if err := json.Unmarshal(bytes, &parsedJson); err != nil {
		return errors.New(err.Error() + "\nFailed to parse message in json format")
	}

	return nil
}
