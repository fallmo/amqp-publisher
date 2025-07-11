### AMQP-PUBLISHER

HELP

```
Usage:
amqp-publisher publish [--flags] [message]


Examples:
amqp-publisher publish --amqp-uri "amqp://localhost:5672" --queue events hello
amqp-publisher publish --amqp-uri "amqp://localhost:5672" --queue events --content-type application/json '{"msg": "hello"}'
AMQP_URI="amqp://localhost:5672" amqp-publisher publish  --queue events --immediate hello


Flags:
      --amqp-uri string       URI connection string to AMQP server
      --content-type string   MIME content type of the message (default "text/plain")
  -h, --help                  help for publish
      --immediate             Undeliver publishing if no consumer on the matched queue is ready to accept the delivery.
      --mandatory             Undeliver publishing if no queue is bound that matches routing key
      --queue string          Queue name / routing key

```
