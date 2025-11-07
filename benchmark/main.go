package main

import (
	"context"
	"fmt"
	"gouniqueusername"
	"gouniqueusername/benchmark/config"
	"gouniqueusername/pb"
	"math/rand"
	"time"
)

var (
	BenchBatchInsertLenth = 5000
	BenchInsertLenth      = 5000
	BenchExistsLenth      = 50000
	BenchDeleteLenth      = 5000
)

func main() {
	config.Load()
	client, err := gouniqueusername.NewClient(gouniqueusername.GRPCConfig{
		Host: "localhost",
		Port: "9000",
	})
	if err != nil {
		panic(err)
	}

	insert := Generate(BenchInsertLenth)
	batchInsert := Generate(BenchBatchInsertLenth)
	exists := Generate(BenchExistsLenth)
	delete := Generate(BenchDeleteLenth)

	start := time.Now()

	BenchInsert(client, insert)
	BenchBatchInsert(client, batchInsert, 1000)
	BenchExists(client, exists)
	BenchDelete(client, delete)

	totalOperations := BenchBatchInsertLenth + BenchInsertLenth + BenchExistsLenth + BenchDeleteLenth
	fmt.Printf("Total time taken: %s, Total operations: %d\n", time.Since(start), totalOperations)
}

func Generate(length int) [][]byte {
	response := make([][]byte, length)
	for i := range len(response) {
		response[i] = RandomUsername(4, 15)
	}

	return response
}

func RandomUsername(minLen, maxLen int) []byte {
	if minLen > maxLen {
		minLen, maxLen = maxLen, minLen
	}

	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(maxLen-minLen+1) + minLen
	result := make([]byte, length)

	for i := range result {
		result[i] = config.GlobalConfig.CharList[rand.Intn(len(config.GlobalConfig.CharList))]
	}
	return result
}

func BenchInsert(c pb.DbServiceClient, values [][]byte) {
	for _, val := range values {
		c.Insert(context.Background(), &pb.SingleInsertRequest{
			Value: string(val),
		})
	}
}

func BenchExists(c pb.DbServiceClient, values [][]byte) {
	for _, val := range values {
		c.CheckIfExists(context.Background(), &pb.CheckIfExistsRequest{
			Value: string(val),
		})
	}
}

func BenchBatchInsert(c pb.DbServiceClient, values [][]byte, batchSize int) {
	valuesStr := []string{}
	for i := 0; i < len(values); i += batchSize {
		end := i + batchSize
		if end > len(values) {
			end = len(values)
		}
		for i, _ := range values[i:end] {
			valuesStr = append(valuesStr, string(values[i]))
		}
		c.BatchInsert(context.Background(), &pb.BatchInsertRequest{
			Values: valuesStr,
		})
	}
}

func BenchDelete(c pb.DbServiceClient, values [][]byte) {
	for _, val := range values {
		c.Delete(context.Background(), &pb.SingleDeleteRequest{Value: string(val)})
	}
}
