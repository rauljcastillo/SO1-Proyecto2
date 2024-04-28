package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type data struct {
	Date  string `json:"date"`
	Album string `json:"album"`
	Name  string `json:"name"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

func main() {
	con, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092",
		"group.id":          "mygroup",
	})

	if err != nil {
		log.Fatal("No se pudo crear el consumidor")
	}

	topic := "mi-topic"
	if er := con.SubscribeTopics([]string{topic}, nil); er != nil {
		log.Fatal("No se pudo suscribir al topic")
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	//Conexion a mongo
	credential := options.Credential{Username: os.Getenv("USER"), Password: os.Getenv("PASS")}
	fmt.Println(os.Getenv("SERVICE_MONGO"))
	uri := fmt.Sprintf("mongodb://%s.so1.svc:27017", os.Getenv("SERVICE_MONGO"))
	db, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAuth(credential))
	collections := db.Database("mydb").Collection("artista")
	//Conexion a redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s.so1.svc:6379", os.Getenv("SERVICE_REDIS")),
		Password: os.Getenv("PASSWORD"),
		DB:       0,
	})

	if err != nil {
		log.Fatal("Error al conectar a mongo")
	}

	run := true
	for run {
		select {
		case <-sigchan:
			run = false
		default:
			ev, err := con.ReadMessage(50 * time.Millisecond)
			if err != nil {
				continue
			}
			var data data
			json.Unmarshal(ev.Value, &data)
			fmt.Println(data.Name)
			if err := json.Unmarshal(ev.Value, &data); err != nil {
				fmt.Println("Error")
				return
			}
			fmt.Println(data.Name, data.Album)
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()
				data.Date = time.Now().Format("2/01/2006 15:04:05")
				v, er := collections.InsertOne(ctx, data)
				if er != nil {
					fmt.Println(er)
					return
				}
				fmt.Println(v.InsertedID)
			}()
			go func() {
				if er := client.HIncrBy(context.Background(), "Votos", data.Name+"-"+data.Album+"-"+data.Year, 1).Err(); er != nil {
					log.Fatal("Error al realizar conteo de votos")
				}
				if er := client.IncrBy(context.Background(), "Total", 1).Err(); er != nil {
					log.Fatal("Error al realizar conteo total de votos")
				}
			}()
		}

	}

	con.Close()
}
