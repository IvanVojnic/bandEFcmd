package repository

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ClosePool is a func to close connection to db
func ClosePool(myPool *pgxpool.Pool) {
	if myPool != nil {
		myPool.Close()
	}
}

// NewPostgresDB func to init and connect to db
func NewClientConn(port string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}
	return conn
}
