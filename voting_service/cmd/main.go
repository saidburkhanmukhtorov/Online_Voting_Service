package main

import (
	"log"
	"net"
	vote "vote/genproto"
	"vote/service"
	"vote/storage/postgres"

	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serVoting := grpc.NewServer()
	vote.RegisterElectionServiceServer(serVoting, service.NewElectionService(db))
	vote.RegisterCandidateServiceServer(serVoting, service.NewCandidateService(db))
	vote.RegisterPublicVoteServiceServer(serVoting, service.NewPublicVoteService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := serVoting.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}