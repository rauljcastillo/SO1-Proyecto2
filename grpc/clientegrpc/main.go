package main

import (
	//confp "Tarea4/protosf"
	"context"
	"fmt"
	"log"
	protosf "main/protof"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insert(c *gin.Context) {
	var data data
	if er := c.BindJSON(&data); er != nil {
		log.Fatal("Error")
	}
	fmt.Println(data.Name, data.Album, data.Rank, data.Year)
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar %v", err)
	}
	fun := protosf.NewInfoClient(conn)
	_, er := fun.GetInfo(context.Background(), &protosf.Req{
		Name:  data.Name,
		Album: data.Album,
		Year:  data.Year,
		Rank:  data.Rank,
	})
	if er != nil {
		log.Fatal("Ocurrio algo")
	}

}
func main() {
	r := gin.Default()
	r.POST("/grpc/data", insert)
	r.Run()
}
