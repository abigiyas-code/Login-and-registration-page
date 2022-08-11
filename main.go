package main

import (
	"github.com/abigiya/internal/Internal/adaptor/framework/https"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root@:26257/users?sslmode=disable"
	serverAddress = "0.0.0.0:8181"
)

func main() {
	server := https.NewDbAdaptor(dbDriver, dbSource)
	server.RequestSendStart(serverAddress)
}
