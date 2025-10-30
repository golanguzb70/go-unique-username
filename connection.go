package gouniqueusername

import (
	"fmt"
	"gouniqueusername/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConfig struct {
	Host string
	Port string
}

func NewClient(cfg GRPCConfig) (pb.DbServiceClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewDbServiceClient(conn), nil
}
