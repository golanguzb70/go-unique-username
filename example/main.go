package main

import (
	"context"
	"fmt"
	"gouniqueusername"
	"gouniqueusername/pb"
)

func main() {
	client, err := gouniqueusername.NewClient(gouniqueusername.GRPCConfig{
		Host: "localhost",
		Port: "9000",
	})
	if err != nil {
		panic(err)
	}

	client.Insert(context.Background(), &pb.SingleInsertRequest{
		Value: "azizbek",
	})

	client.BatchInsert(context.Background(), &pb.BatchInsertRequest{
		Values: []string{"aziz", "asad", "baxrom", "baxriddin"},
	})

	exists, err := client.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{Value: "azizbek"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exists - True:  %v\n", exists.Exists)

	exists, err = client.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{Value: "baxrom"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exists - True:  %v\n", exists.Exists)

	exists, err = client.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{Value: "aziz"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exists - True: %v\n", exists.Exists)

	exists, err = client.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{Value: "azim"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exists - False: %v\n", exists.Exists)

	_, err = client.Delete(context.Background(), &pb.SingleDeleteRequest{Value: "aziz"})
	if err != nil {
		panic(err)
	}

	exists, err = client.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{Value: "aziz"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Exists - False: %v\n", exists.Exists)
}
