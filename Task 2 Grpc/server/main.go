package main

import (
	"grpc-server/pb"
	"grpc-server/repository"
	"grpc-server/service"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	err := repository.ConnectToDatabase()

	if err != nil {
		log.Fatal("Error while connecting db", err)
	}

	SVC_PORT := os.Getenv("PORT")

	server, err := net.Listen("tcp", SVC_PORT)

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, service.GetService())

	grpcServer.Serve(server)

}
