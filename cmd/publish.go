/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"amqp-publisher/utils"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mandatory bool
var immediate bool
var contentType string

var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "Send publishing to an AMQP exchange server",
	Example: `amqp-publisher publish --amqp-uri amqp://localhost:5672 --queue events ping`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("publish expects only 1 message")
		}
		return nil
	},
	PreRun: func(_ *cobra.Command, _ []string) {
		if amqpUri == "" {
			amqpUri = os.Getenv("AMQP_URI")
		}

		if amqpUri == "" {
			fmt.Println("Error: AMQP connection string must be provided, either via 'AMQP_URI' environment variable, or via the '--amqp-uri' flag")
			os.Exit(1)
		}

		if queue == "" {
			queue = os.Getenv("QUEUE")
		}
		if queue == "" {
			fmt.Println("Error: Queue name must be provided, either via 'QUEUE' environment variable, or via the '--queue' flag")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.ValidatePublishContentType(contentType)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		messageBytes := []byte(args[0])

		if contentType == "application/json" {
			if err = utils.ValidateJsonMessage(messageBytes); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if err = utils.SendAMQPMessage(amqpUri, queue, mandatory, immediate, contentType, messageBytes); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Message successfully published to queue '" + queue + "'.")
	},
}

func init() {
	publishCmd.Flags().StringVar(&amqpUri, "amqp-uri", "", "URI connection string to AMQP server")
	publishCmd.Flags().StringVar(&queue, "queue", "", "Queue name / routing key")
	publishCmd.Flags().BoolVar(&mandatory, "mandatory", false, "Undeliver publishing if no queue is bound that matches routing key")
	publishCmd.Flags().BoolVar(&immediate, "immediate", false, "Undeliver publishing if no consumer on the matched queue is ready to accept the delivery.")
	publishCmd.Flags().StringVar(&contentType, "content-type", "text/plain", "MIME content type of the message")

}
