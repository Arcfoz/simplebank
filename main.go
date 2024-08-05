package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/arcfoz/simplebank/api"
	db "github.com/arcfoz/simplebank/db/sqlc"
	"github.com/arcfoz/simplebank/gapi"
	"github.com/arcfoz/simplebank/pb"
	"github.com/arcfoz/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGrcpServer(config, store)

}

func runGrcpServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grcpServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grcpServer, server)
	reflection.Register(grcpServer)

	listener, err := net.Listen("tcp", config.GPRCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grcpServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
