package main

import (
	"context"
	"fmt"
	"grpc-client/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	//Create User
	createUserResp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
		Name: "Amal",
		Age:  21,
	})
	if err != nil {
		log.Fatal(err)
	}

	//GetAllUsers
	getAllUsersResp, err := client.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})
	if err != nil {
		log.Fatal(err,)
	}

	//GetUserByName
	getUserByNameResp, err := client.GetUserByName(context.Background(), &pb.GetUserByNameRequest{Name: "Ashiq"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createUserResp)
	fmt.Println(getAllUsersResp)
	fmt.Println(getUserByNameResp)

	if err != nil {
		log.Fatal(err)
	}

}
