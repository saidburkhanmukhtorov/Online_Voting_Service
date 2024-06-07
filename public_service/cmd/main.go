package main

import (
	"log"
	"net"

	pub "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
	"github.com/myfirstgo/online_voting_service/public_service/service"


	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":7080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serPub := grpc.NewServer()
	log.Printf("server listening at %v", liss.Addr())
	if err := serPub.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	pub.RegisterPartyServiceServer(serPub, service.NewPartyService(db))
	pub.RegisterPublicServiceServer(serPub, service.NewPublicService(db))

	if err := serPub.Serve(liss); err!= nil {
        log.Fatalf("failed to serve: %v", err)
    }
	log.Printf("server listening at %v", liss.Addr())
	

}
