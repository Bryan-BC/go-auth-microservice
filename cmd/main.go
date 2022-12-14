package main

import (
	"log"
	"net"

	"github.com/Bryan-BC/go-auth-microservice/pkg/config"
	"github.com/Bryan-BC/go-auth-microservice/pkg/db"
	"github.com/Bryan-BC/go-auth-microservice/pkg/pb"
	"github.com/Bryan-BC/go-auth-microservice/pkg/services"
	"github.com/Bryan-BC/go-auth-microservice/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Panicf("Error loading config, %s \n", err)
	}

	db := db.Init(c.DBURL)

	jwt := utils.JWTWrapper{
		Secret:          c.Secret,
		Issuer:          "go-auth-microservice",
		ExpirationHours: 24 * 7,
	}

	listener, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Panicf("Error listening, %s \n", err)
	}

	log.Printf("Auth microservice listening on port %s \n", c.Port)

	s := services.Server{
		JWT:       &jwt,
		DBPointer: &db,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Panicf("Error serving auth microservice, %s", err)
	}
}
