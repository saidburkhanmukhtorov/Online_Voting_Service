package main

import (
	"fmt"
	"log"

	// "github.com/google/uuid"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"

	"context"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	partyDb := postgres.NewParty(db.Db)



	r, err := partyDb.GetById(context.TODO(), &public.PartyById{
		Id:         "173121f5-18d0-4b13-aca3-05d0426df97c",
	})

	fmt.Println(r, err)

	// id := uuid.New().String()
	// pb, err := partyDb.Create(context.TODO(), &public.PartyCreate{

	// 	Id:          id,
	// 	Name:        "sdgfdg",
	// 	Slogan:      "gdf",
	// 	OpenedDate:  "2022-01-01",
	// 	Description: "test",
	// })

	// fmt.Println(pb, err)
	// liss, err := net.Listen("tcp", ":8081")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// vote.RegisterElectionServiceServer(s, service.NewElectionService(db))
	// log.Printf("server listening at %v", liss.Addr())
	// if err := s.Serve(liss); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
