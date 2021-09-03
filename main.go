package main

import (
	"context"
	"fmt"
	"github.com/INEFFABLE-games/Birds/config"
	"github.com/INEFFABLE-games/Birds/protocol"
	"github.com/INEFFABLE-games/Birds/repository"
	"github.com/INEFFABLE-games/Birds/server"
	"github.com/INEFFABLE-games/Birds/service"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
)

func getMongoConnection(connString string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(connString)

	conn, err := mongo.Connect(context.Background(), clientOptions)

	return conn, err
}

func main() {
	cfg := config.NewConfig()

	conn, err := getMongoConnection(cfg.MongoConnString)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "main",
			"action":  "ping mongo connection",
		}).Errorf("unable connect to mongo %v", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.Port))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "main",
			"action":  "create listener",
		}).Errorf("unable to create listener %v", err.Error())
	}

	birdRepo := repository.NewBirdRepository(conn)
	birdService := service.NewBirdService(birdRepo)
	birdHandler := server.NewBirdHandler(birdService)

	grpcServer := grpc.NewServer()

	protocol.RegisterBirdServiceServer(grpcServer, birdHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
