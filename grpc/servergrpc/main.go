package main

import (
	"context"
	"fmt"
	"log"
	protosf "main/protof"
	"net"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
)

type server struct {
	protosf.UnimplementedInfoServer
}

func (srv *server) GetInfo(c context.Context, data *protosf.Req) (*protosf.Res, error) {
	topic := "mi-topic"
	proc, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092",
	})

	if err != nil {
		log.Fatal("Error al crear el proc")
	}
	er := proc.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("data"),
		Value: []byte(fmt.Sprintf(`{"name": "%s","album": "%s","year": "%s","rank": "%s"}`,
			data.Name,
			data.Album,
			data.Year,
			data.Rank)),
	}, nil)
	if er != nil {
		return &protosf.Res{}, nil
	}
	proc.Flush(1000)
	proc.Close()
	return &protosf.Res{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Error al escuchar")
	}
	s := grpc.NewServer()
	protosf.RegisterInfoServer(s, &server{})
	log.Printf("server listening on port %d", 8081)
	if er := s.Serve(lis); er != nil {
		log.Fatal("Error en el server")
	}
}
