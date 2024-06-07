package main

import (
	"log"
	"net"

	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/service"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	public.RegisterPartyServiceServer(s, service.NewPartyService(db))
	public.RegisterPublicServiceServer(s, service.NewPublicService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
