package main
import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)
const (
	bootstrapServers = "kafka01:9093"
)

func main() {

	topic := "go-test-topic"
	createTopic(topic)
	// Produce a new record to the topic...
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
                "bootstrap.servers": bootstrapServers,
                "security.protocol": "SSL",
				"ssl.ca.location": "/root/ssl/caRoot.pem",
                "ssl.key.location": "/root/ssl/certificate.key",
                "ssl.certificate.location": "/root/ssl/certificate.pem",
				// "debug": "security,broker",
				"enable.ssl.certificate.verification" : false,
                // "ssl.key.password": "radisys",
		})

	if err != nil {
		panic(fmt.Sprintf("Failed to create producer: %s", err))
	}

	value := "golang test value"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic,
			Partition: kafka.PartitionAny},
		Value: []byte(value)}, nil)

	// Wait for delivery report
	e := <-producer.Events()

	message := e.(*kafka.Message)
	if message.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n",
			message.TopicPartition)
	} else {
		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
			*message.TopicPartition.Topic,
			message.TopicPartition.Partition,
			message.TopicPartition.Offset)
	}

	producer.Close()

	// Now consumes the record and print its value...
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
                "bootstrap.servers": bootstrapServers,
                "security.protocol": "SSL",
				"ssl.ca.location": "/root/ssl/caRoot.pem",
                "ssl.key.location": "/root/ssl/certificate.key",
                "ssl.certificate.location": "/root/ssl/certificate.pem",
				// "debug": "security,broker",
				"enable.ssl.certificate.verification" : false,
                // "ssl.key.password": "radisys",
		"session.timeout.ms":      6000,
		"group.id":                "my-group",
		"auto.offset.reset":       "earliest"})

	if err != nil {
		panic(fmt.Sprintf("Failed to create consumer: %s", err))
	}

	topics := []string{topic}
	consumer.SubscribeTopics(topics, nil)

	for {
		message, err := consumer.ReadMessage(100 * time.Millisecond)
		if err == nil {
			fmt.Printf("consumed from topic %s [%d] at offset %v: "+
				string(message.Value), *message.TopicPartition.Topic,
				message.TopicPartition.Partition, message.TopicPartition.Offset)
		}
	}

	consumer.Close()

}

func createTopic(topic string) {
	fmt.Println("11111111")
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
                "bootstrap.servers": bootstrapServers,
                "security.protocol": "SSL",
				"ssl.ca.location": "/root/ssl/caRoot.pem",
                "ssl.key.location": "/root/ssl/certificate.key",
                "ssl.certificate.location": "/root/ssl/certificate.pem",
				"enable.ssl.certificate.verification" : false,
		"broker.version.fallback": "0.10.0.0",
		"api.version.fallback.ms": 0})

	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("22222")

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDuration, err := time.ParseDuration("60s")
	if err != nil {
		panic("time.ParseDuration(60s)")
	}

	results, err := adminClient.CreateTopics(ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 3}},
		kafka.SetAdminOperationTimeout(maxDuration))

	if err != nil {
		fmt.Printf("Problem during the topic creation: %v\n", err)
		os.Exit(1)
	}

	// Check for specific topic errors
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError &&
			result.Error.Code() != kafka.ErrTopicAlreadyExists {
			fmt.Printf("Topic creation failed for %s: %v",
				result.Topic, result.Error.String())
			os.Exit(1)
		}
	}

	adminClient.Close()

}


